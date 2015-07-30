// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

package router

import (
	"errors"
	"net/http"
	"strings"
)

// ErrInvalidPath - URL parse error; invalid route path.
var ErrInvalidPath = errors.New("invalid route path")

// HandlerFunc defines the interface for
// routify functions, identical to http.HandlerFunc
// except that it includes the parsed URL parameters.
type HandlerFunc func(http.ResponseWriter, *http.Request, Params)

// Router represents the defined routes and parameter validators.
type Router struct {
	Routes     Routes
	Validators map[string]func(string) bool
}

// Routes holds static route mappings.
type Routes map[string]*Route

// Validators holds parameter validating functions.
type Validators map[string]func(string) bool

// Route represents an individual route/end-point.
type Route struct {
	Param       string            // Parameter name
	Check       func(string) bool // Function to check if section is valid
	HandlerFunc HandlerFunc       // Handler function use to serve
	Child       *Route            // Child route (parameter capture)
	Children    Routes            // Child map (static paths)
}

// Get attempts to get a route for the given request.
func (r *Router) Get(req *http.Request) (HandlerFunc, Params, error) {
	route, exists := r.Routes[req.Method]

	if !exists {
		return nil, nil, ErrInvalidMethod
	}

	u := req.URL.Path

	if u == "" || u == "/" {
		if route, exists = route.Children["/"]; exists && route.HandlerFunc != nil {
			return route.HandlerFunc, nil, nil
		}

		return nil, nil, ErrRouteNotFound
	} else if u[0] == '/' {
		u = u[1:]
	}

	var p Params

	for {
		if v, exists := route.Children[u]; exists {
			route = v

			break
		}

		s := u
		i := strings.IndexByte(u, '/')

		if i != -1 {
			s = u[:i]
			u = u[i+1:]
		}

		if route.Child != nil {
			// Capture parameter
			route = route.Child

			if route.Check != nil && !route.Check(s) {
				return nil, nil, ErrRouteNotFound
			}

			p = append(p, param{route.Param, s})

			if i == -1 {
				break
			}
		} else if route, exists = route.Children[s]; !exists {
			// Static
			return nil, nil, ErrRouteNotFound
		}
	}

	if route == nil || route.HandlerFunc == nil {
		return nil, nil, ErrRouteNotFound
	}

	return route.HandlerFunc, p, nil
}

// Add adds a route for the given method to the routes map.
func (r *Router) Add(m, u string, h HandlerFunc) error {
	m = strings.ToUpper(m)

	var c *Route

	if r.Routes == nil {
		r.Routes = Routes{m: &Route{}}
		c = r.Routes[m]
	} else if _, exists := r.Routes[m]; !exists {
		r.Routes[m] = &Route{}
		c = r.Routes[m]
	}

	var p []string

	if u != "/" {
		p = strings.Split(u, "/")
	} else {
		p = []string{"/"}
	}

	for i, l := 0, len(p); i < l; i++ {
		if p[i] == "" {
			return ErrInvalidPath
		}

		// Parameter
		if p[i][0] == '$' {
			if c.Child == nil || c.Child.Param != p[i][1:] {
				c.Child = &Route{
					Param: p[i][1:],
					Check: r.Validators[p[i]],
				}
			}

			c = c.Child

			continue
		}

		v, n := staticPath(p[i:])
		i += n

		if c.Children == nil {
			c.Children = Routes{v: &Route{}}
		} else if _, exists := c.Children[v]; !exists {
			c.Children[v] = &Route{}
		}

		c = c.Children[v]
	}

	c.HandlerFunc = h

	return nil
}

// AddValidator adds a validating function to
// the validators map.
func (r *Router) AddValidator(n string, f func(string) bool) {
	if r.Validators == nil {
		r.Validators = Validators{n: f}
	} else {
		r.Validators[n] = f
	}
}

// ServeHTTP implements the Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h, p, err := r.Get(req)

	if err != nil {
		if e, ok := err.(*Error); ok {
			w.WriteHeader(e.code)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	h(w, req, p)
}

// VOID handler for testing.
func exampleHandler(w http.ResponseWriter, r *http.Request, p Params) {
}
