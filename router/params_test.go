// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

package router

import (
	"errors"
	"net/http"
	"strconv"
	"testing"
	"time"
)

type month time.Month

func (m *month) Scan(i interface{}) error {
	v, ok := i.(string)

	if !ok {
		return errors.New("unsupported type")
	}

	c, err := strconv.Atoi(v)

	if err != nil {
		return err
	} else if c < 0 || c > 12 {
		return errors.New("month out of bounds")
	}

	*m = month(c)

	return nil
}

func TestParams(t *testing.T) {
	req, err := http.NewRequest("GET", shortParam, nil)

	if err != nil {
		t.Fatal(err)
	}

	_, p, _ := routes.Get(req)

	if v, err := p.GetInt("year"); err != nil {
		t.Fatal(err)
	} else if v != 2015 {
		t.Fatal("unexpected value")
	}

	var b []byte

	if err := p.Scan("day", &b); err != nil {
		t.Fatal(err)
	} else if string(b) != "12" {
		t.Fatal("unexpected value")
	}

	var m month

	if err := p.Scan("month", &m); err != nil {
		t.Fatal(err)
	} else if time.Month(m).String() != "February" {
		t.Fatal("unexpected value")
	}
}