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
		if v[0] == ':' || v[0] == '$' {
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
