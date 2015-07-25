// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

package main

import (
	"bytes"
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
	table         routemap
	funcs         routemap
	check, handle string
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

	buf := &bytes.Buffer{}

	fmt.Fprintf(buf, `package %s 

import "github.com/martingallagher/routify/router"

var %s = router.Routes{`, *packageName, *varName)

	for k, v := range r.routes {
		if len(v.table) == 0 && len(v.funcs) == 0 {
			continue
		}

		writeRule(buf, k, v)
	}

	buf.WriteString("\n}\n")

	if _, err = buf.WriteTo(f); err != nil {
		log.Fatal(err)
	}
}

func (r *routes) add(method, path, handle string) error {
	if _, exists := r.routes[method]; !exists {
		r.routes[method] = &route{table: routemap{}, funcs: routemap{}}
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

	for i, l := 0, len(p)-1; i <= l; i++ {
		if p[i] == "" {
			return errInvalidPath
		}

		isParam := p[i][0] == '$'
		m := c.table

		if isParam {
			if _, exists := r.params[p[i][1:]]; exists {
				p[i] = p[i][1:]
				m = c.funcs
			}
		}

		if _, exists := m[p[i]]; !exists {
			m[p[i]] = &route{table: routemap{}, funcs: routemap{}}
		}

		c = m[p[i]]

		if isParam {
			c.check = r.params[p[i]]
		}

		if i < l {
			continue
		}

		c.handle = handle
	}

	return nil
}

func writeRule(buf *bytes.Buffer, p string, r *route) {
	fmt.Fprintf(buf, "\n\"%s\": &router.Route{", p)

	if r.handle != "" {
		fmt.Fprintf(buf, "HandlerFunc: %s,", r.handle)
	}

	if r.check != "" {
		fmt.Fprintf(buf, "Check: %s,", r.check)
	}

	if len(r.table) > 0 {
		buf.WriteString("\nTable: router.Routes{")

		for k, v := range r.table {
			writeRule(buf, k, v)
		}

		buf.WriteString("},")
	}

	if len(r.funcs) > 0 {
		buf.WriteString("\nFuncs: router.Routes{")

		for k, v := range r.funcs {
			writeRule(buf, k, v)
		}

		buf.WriteString("},")
	}

	buf.WriteString("},")
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

	var m map[string]interface{}

	if err = yaml.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	var (
		l [][]string
		r = &routes{map[string]string{}, routemap{}}
	)

	for k, v := range m {
		switch u := strings.ToUpper(k); u {
		case "PARAMS", "GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HELP":
			p, ok := v.(map[interface{}]interface{})

			if !ok {
				continue
			}

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
			p, ok := v.(map[interface{}]interface{})

			if !ok {
				continue
			}

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
