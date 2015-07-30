// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

package router

import (
	"net/http"
	"testing"
)

const (
	longStatic = "/static/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u"
	shortParam = "/schemas/test/archives/2015/02/12"
	longParam  = "/nofunc/a/b/c/d/e/f/g/h/i/j/k/l/m/n/o/p/q/r/s/t/u"
)

func TestRuntimeRouter(t *testing.T) {
	r := &Router{}
	r.AddValidator(":year", IsYear)
	r.AddValidator(":month", IsMonth)
	r.AddValidator(":day", IsDay)

	if err := r.Add("GET", "/schemas/:schema/archives/:year/:month/:day", exampleHandler); err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", shortParam, nil)

	if err != nil {
		t.Fatal(err)
	} else if _, p, err := r.Get(req); err != nil {
		t.Fatal(err)
	} else if p.Get("year") != "2015" {
		t.Fatal("unexpected value")
	}
}

func TestRouter(t *testing.T) {
	req, err := http.NewRequest("GET", shortParam, nil)

	if err != nil {
		t.Fatal(err)
	}

	h, p, err := routes.Get(req)

	if err != nil {
		t.Fatal(err)
	} else if len(p) == 0 {
		t.Fatal("empty params")
	} else if h == nil {
		t.Fatal("nil handler")
	}

	t.Logf("%#v", p)
}

func TestRouterLongSimple(t *testing.T) {
	req, err := http.NewRequest("GET", longParam, nil)

	if err != nil {
		t.Fatal(err)
	}

	h, p, err := routes.Get(req)

	if err != nil {
		t.Fatal(err)
	} else if len(p) == 0 || p.Get("g") == "" {
		t.Fatal("empty params")
	} else if h == nil {
		t.Fatal("nil handler")
	}

	t.Logf("%#v", p)
}

func TestRouterLongStatic(t *testing.T) {
	req, err := http.NewRequest("GET", longStatic, nil)

	if err != nil {
		t.Fatal(err)
	} else if h, _, err := routes.Get(req); err != nil {
		t.Fatal(err)
	} else if h == nil {
		t.Fatal("nil handler")
	}
}

func BenchmarkRouterLongSimple(b *testing.B) {
	req, err := http.NewRequest("GET", longParam, nil)

	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		routes.Get(req)
	}
}

func BenchmarkRouterLongStatic(b *testing.B) {
	req, err := http.NewRequest("GET", longStatic, nil)

	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		routes.Get(req)
	}
}

func BenchmarkRouterMultiParam(b *testing.B) {
	req, err := http.NewRequest("GET", shortParam, nil)

	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		routes.Get(req)
	}
}
