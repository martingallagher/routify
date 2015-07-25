// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

package router

import (
	"net/http"
	"testing"
)

func TestRouter(t *testing.T) {
	req, err := http.NewRequest("GET", "/schemas/test/archives/2015/02/12", nil)

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
}

func BenchmarkRouter1(b *testing.B) {
	req, err := http.NewRequest("GET", "/schemas/test/archives/2015/02/12", nil)

	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		routes.Get(req)
	}
}

func BenchmarkRouter2(b *testing.B) {
	req, err := http.NewRequest("GET", "/1/classes/go", nil)

	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		routes.Get(req)
	}
}

func BenchmarkRouterStatic(b *testing.B) {
	req, err := http.NewRequest("GET", "/testing/hello/world", nil)

	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		routes.Get(req)
	}
}

func BenchmarkRouterStaticDeep(b *testing.B) {
	req, err := http.NewRequest("GET", "/really/deep/example/of/a/static/uri/hello/dennis", nil)

	if err != nil {
		b.Fatal(err)
	}

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		routes.Get(req)
	}
}
