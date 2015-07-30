// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

package router

// IsYear tests if the string is a valid year (YYYY).
func IsYear(s string) bool {
	return len(s) == 4 &&
		s[0] > 47 || s[0] < 58 &&
		s[1] > 47 || s[1] < 58 &&
		s[2] > 47 || s[2] < 58 &&
		s[3] > 47 || s[3] < 58
}

// IsMonth tests if the string is a valid month (MM).
func IsMonth(s string) bool {
	if len(s) != 2 || s[0] > 49 || s[1] > 57 || s[0] == 48 && s[1] == 48 {
		return false
	}

	return s[0] == 48 || s[1] < 51
}

// IsDay tests if the string is a valid day (DD).
func IsDay(s string) bool {
	if len(s) != 2 || s[0] > 51 || s[1] > 57 || s[0] == 48 && s[1] == 48 {
		return false
	}

	return s[0] < 51 || s[1] < 50
}

func staticPath(p []string) (string, int) {
	s := ""
	c := -1

	for i, v := range p {
		if v[0] == '$' {
			break
		}

		if i > 0 {
			s += "/"
		}

		s += v
		c++
	}

	return s, c
}
