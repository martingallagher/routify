// Copyright Praegressus Limited. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

var testData = []struct{ method, route, url, param, value string }{
	{"GET", "/schemas/:schema/archives/:year/:month/:day", shortParam, "year", "2015"},
	{"POST", "/authorizations/", "/authorizations/", "", ""},
	{"POST", "/authorizations/:id", "/authorizations/123", "id", "123"},
	{"GET", "/repos/$owner/$repo/events", "/repos/1/2/events", "owner", "1"},
	{"GET", "/users/:user/received_events/public", "/users/1/received_events/public", "user", "1"},
}

func TestRuntimeRouter(t *testing.T) {
	r := &Router{}
	r.AddValidator(":year", IsYear)
	r.AddValidator(":month", IsMonth)
	r.AddValidator(":day", IsDay)

	for _, c := range testData {
		if err := r.Add(c.method, c.route, exampleHandler); err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(c.method, c.url, nil)

		if err != nil {
			t.Fatal(err)
		} else if h, p, err := r.Get(req); err != nil {
			t.Fatal(err)
		} else if h == nil {
			t.Fatal("nil handler")
		} else if c.param != "" && p.Get(c.param) != c.value {
			t.Fatal("unexpected value")
		}
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
