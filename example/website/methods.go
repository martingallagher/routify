// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

package website

import (
	"fmt"
	"net/http"

	"github.com/martingallagher/routify/router"
)

func index(w http.ResponseWriter, r *http.Request, p router.Params) {
	w.Write([]byte("Welcome! Please try /hello/$name or /printnum/123"))
}

func hello(w http.ResponseWriter, r *http.Request, p router.Params) {
	fmt.Fprintf(w, "Hello %s, have a good day!", p.Get("str"))
}

func printnum(w http.ResponseWriter, r *http.Request, p router.Params) {
	fmt.Fprintf(w, "printnum() = %s", p.Get("num"))
}

func validateNumber(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}
