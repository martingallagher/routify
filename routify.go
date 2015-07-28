// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	inputFile       = flag.String("i", "routes.yaml", "Routes input file")
	outputFile      = flag.String("o", "routes.go", "Routes output file")
	packageName     = flag.String("p", "", "Package name")
	varName         = flag.String("v", "Routes", "Variable name")
	errInvalidInput = errors.New("missing routes input file")
	errInvalidPath  = errors.New("invalid route path")
)

type routemap map[string]*route

type routes struct {
	params map[string]string
	routes routemap
}

type route struct {
	child                *route
	children             routemap
	param, check, handle string
}

func main() {
	flag.Parse()
	log.SetFlags(log.Lmicroseconds)

	if *inputFile == "" {
		log.Fatal("input filename is required (use -i flag)")
	} else if *packageName == "" {
		log.Fatal("package name is required (use -p flag)")
	}

	f, err := os.OpenFile(*outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r, err := loadRoutes()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(f, `package %s 

import "github.com/martingallagher/routify/router"

var %s =  router.Routes{
`, *packageName, *varName)

	for k, v := range r.routes {
		if len(v.children) == 0 {
			continue
		}

		r.writeRule(f, k, v)
	}

	f.WriteString("\n}")

	if err = f.Sync(); err != nil {
		log.Fatal(err)
	}
}

func (r *routes) add(method, path, handle string) error {
	if _, exists := r.routes[method]; !exists {
		r.routes[method] = &route{children: routemap{}}
	}

	var (
		p []string
		c = r.routes[method]
	)

	if path != "/" {
		p = strings.Split(path, "/")
	} else {
		p = []string{"/"}
	}

	for _, v := range p {
		if v == "" {
			return errInvalidPath
		}

		// Parameter
		if v[0] == '$' {
			c.child = &route{
				param:    v[1:],
				check:    r.params[v],
				children: routemap{},
			}

			c = c.child

			continue
		}

		// Allocate map for static routes
		if _, exists := c.children[v]; !exists {
			c.children[v] = &route{children: routemap{}}
		}

		c = c.children[v]
	}

	c.handle = handle

	return nil
}

func (r *routes) writeChild(f *os.File, c *route) {
	fmt.Fprintf(f, "Child: &router.Route{\nParam: \"%s\",\n", c.param)

	if c.check != "" {
		fmt.Fprintf(f, "Check: %s,\n", c.check)
	}

	if c.handle != "" {
		fmt.Fprintf(f, "HandlerFunc: %s,\n", c.handle)
	}

	if len(c.children) > 0 {
		r.writeChildren(f, c)
	} else if c.child != nil {
		r.writeChild(f, c.child)
	}

	f.WriteString("},\n")
}

func (r *routes) writeChildren(f *os.File, c *route) {
	f.WriteString("Children: router.Routes{\n")

	for k, v := range c.children {
		r.writeRule(f, k, v)
	}

	f.WriteString("},\n")
}

func (r *routes) writeRule(f *os.File, p string, c *route) {
	fmt.Fprintf(f, "\"%s\": &router.Route{\n", p)

	if c.handle != "" {
		fmt.Fprintf(f, "HandlerFunc: %s,\n", c.handle)
	}

	if len(c.children) > 0 {
		r.writeChildren(f, c)
	} else if c.child != nil {
		r.writeChild(f, c.child)
	}

	f.WriteString("},\n")
}

func loadRoutes() (*routes, error) {
	f, err := os.Open(*inputFile)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)

	if err != nil {
		return nil, err
	}

	f.Close()

	var m map[string]interface{}

	if err = yaml.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	var (
		l [][]string
		r = &routes{map[string]string{}, routemap{}}
	)

	for k, v := range m {
		p, ok := v.(map[interface{}]interface{})

		if !ok {
			continue
		}

		switch u := strings.ToUpper(k); u {
		case "PARAMS", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HELP":
			for a, b := range p {
				t, ok := a.(string)

				if !ok {
					break
				}

				f, ok := b.(string)

				if !ok {
					break
				}

				if u == "PARAMS" {
					r.params[t] = f

					continue
				}

				l = append(l, []string{u, t, f})
			}

		default:
			for a, b := range p {
				t, ok := a.(string)

				if !ok {
					break
				} else if t != "GET" && t != "POST" && t != "PUT" && t != "PATCH" && t != "DELETE" && t != "OPTIONS" && t != "HELP" {
					continue
				}

				f, ok := b.(string)

				if !ok {
					continue
				}

				l = append(l, []string{t, k, f})
			}
		}
	}

	for _, c := range l {
		if err = r.add(c[0], c[1], c[2]); err != nil {
			return nil, err
		}
	}

	return r, nil
}
