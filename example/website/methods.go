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
