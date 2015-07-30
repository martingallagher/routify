// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

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
