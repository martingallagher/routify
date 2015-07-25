// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

package router

import (
	"net/http"
	"strconv"
)

// VOID handler for testing.
func exampleHandler(w http.ResponseWriter, r *http.Request, p Params) {
}

// IsYear tests if the string is a valid year (YYYY).
func IsYear(s string) bool {
	if len(s) != 4 {
		return false
	}

	if _, err := strconv.ParseInt(s, 10, 64); err != nil {
		return false
	}

	return true
}

// IsMonth tests if the string is a valid month (MM).
func IsMonth(s string) bool {
	if len(s) != 2 || s == "00" {
		return false
	}

	if i, err := strconv.ParseInt(s, 10, 64); err != nil || i < 1 || i > 12 {
		return false
	}

	return true
}

// IsDay tests if the string is a valid day (DD).
func IsDay(s string) bool {
	if len(s) != 2 || s == "00" {
		return false
	}

	if i, err := strconv.ParseInt(s, 10, 64); err != nil || i < 1 || i > 31 {
		return false
	}

	return true
}
