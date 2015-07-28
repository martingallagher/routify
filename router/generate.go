// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

//go:generate routify -i routes.yaml -p router -v routes
// Strip import for in-package testing
//go:generate sed -i "s/^.*import.*//g" routes.go
// Strip package identifiers
//go:generate sed -i "s/router\\.//g" routes.go
// Format
//go:generate gofmt -w -s routes.go

package router
