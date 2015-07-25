// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

package router

import (
	"net/http"
	"strings"
)

// HandlerFunc defines the interface for
// routify functions, identical to http.HandlerFunc
// except that it includes the parsed URL parameters.
type HandlerFunc func(http.ResponseWriter, *http.Request, Params)

// Routes represents the defined routes.
type Routes map[string]*Route

// Params contains the parsed URL parameters.
type Params map[string]string

// Get returns the parameter value for the given
// key, returning false if the key is non-existent.
func (p Params) Get(k string) (string, bool) {
	v, ok := p[k]

	return v, ok
}

// Route represents an individual route/end-point.
type Route struct {
	Check       func(string) bool // Function to check if section is valid
	HandlerFunc HandlerFunc       // Handler function use to serve
	Table       Routes            // Child routes, simple, map-based
	Funcs       Routes            // Child routes requiring validation by the Check function
	Param       bool              // Determines if the route should be added to the returned parameters
}

// Get attempts to get a route for the given request.
func (m Routes) Get(r *http.Request) (HandlerFunc, Params, error) {
	c := m[r.Method]

	if c == nil {
		return nil, nil, ErrInvalidMethod
	}

	u := r.URL.Path

	if u == "" || u == "/" {
		if v, exists := m[r.Method].Table["/"]; exists {
			return v.HandlerFunc, nil, nil
		}

		return nil, nil, ErrRouteNotFound
	}

	if u[0] == '/' {
		u = u[1:]
	}

	var (
		route *Route
		p     = Params{} // This is often a wasted allocation, solve...
	)

	for {
		i := strings.IndexByte(u, '/')

		if i == -1 {
			if v, exists := c.Table[u]; exists {
				route = v
			} else if v = checkTable(u, c.Table, p); v != nil {
				route = v
			} else if v = checkFuncs(u, c.Funcs, p); v != nil {
				route = v
			}

			break
		}

		// Table lookups are faster and take priority
		if s, exists := c.Table[u[:i]]; exists {
			c, u = s, u[i+1:]

			continue
		}

		if m := checkTable(u[:i], c.Table, p); m != nil {
			c = m
		} else if m = checkFuncs(u[:i], c.Funcs, p); m != nil {
			c = m
		} else {
			break
		}

		u = u[i+1:]
	}

	if route != nil {
		return route.HandlerFunc, p, nil
	}

	return nil, nil, ErrRouteNotFound
}

func checkTable(u string, r Routes, p Params) *Route {
	for k, v := range r {
		if k[0] == '$' {
			p[k[1:]] = u

			return v
		}
	}

	return nil
}

func checkFuncs(u string, r Routes, p Params) *Route {
	for k, v := range r {
		if v.Check != nil && v.Check(u) {
			p[k] = u

			return v
		}
	}

	return nil
}
