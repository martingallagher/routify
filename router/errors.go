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

import "net/http"

var (
	// ErrInvalidMethod represents an invalid HTTP method.
	ErrInvalidMethod = NewError(http.StatusMethodNotAllowed, "invalid HTTP method")
	// ErrRouteNotFound represents HTTP 404.
	ErrRouteNotFound = NewError(http.StatusNotFound, "route not found")
	// ErrBadRequest represents HTTP 400.
	ErrBadRequest = NewError(http.StatusBadRequest, "bad request")
)

// Error represents a routing error.
type Error struct {
	code int
	err  string
}

// StatusCode returns the HTTP status code
// associated with the error.
func (e *Error) StatusCode() int {
	return e.code
}

// Error returns the error string.
func (e *Error) Error() string {
	return e.err
}

// NewError returns a new error with
// the given HTTP status code and error message.
func NewError(c int, s string) *Error {
	return &Error{c, s}
}
