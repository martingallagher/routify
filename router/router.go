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

// Route represents an individual route/end-point.
type Route struct {
	Param       string            // Parameter name
	Check       func(string) bool // Function to check if section is valid
	HandlerFunc HandlerFunc       // Handler function use to serve
	Child       *Route            // Child route (parameter capture)
	Children    Routes            // Children (static paths)
}

// Get attempts to get a route for the given request.
func (m Routes) Get(r *http.Request) (HandlerFunc, Params, error) {
	route, exists := m[r.Method]

	if !exists {
		return nil, nil, ErrInvalidMethod
	}

	u := r.URL.Path

	if u == "" || u == "/" {
		if route, exists = m[r.Method].Children["/"]; exists && route.HandlerFunc != nil {
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

// VOID handler for testing.
func exampleHandler(w http.ResponseWriter, r *http.Request, p Params) {
}
