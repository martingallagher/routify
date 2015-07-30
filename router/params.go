// Copyright 2015 Martin Gallagher. All rights reserved.
// Use of this source code is governed by the Apache License,
// Version 2.0 that can be found in the LICENSE file.

package router

import (
	"errors"
	"reflect"
	"strconv"
)

var (
	// ErrUnsupportedType - unsupported variable destination types.
	ErrUnsupportedType = errors.New("unsupported destination type")
	// ErrValueDestination - destination isn't a pointer.
	ErrValueDestination = errors.New("destination is a value, not a pointer")
)

type param struct{ k, v string }

// Params contains the parsed URL parameters.
type Params []param

// Get returns the parameter value for the given key.
func (p Params) Get(k string) string {
	for _, c := range p {
		if c.k == k {
			return c.v
		}
	}

	return ""
}

// GetInt attempts to get the given key as int64.
func (p Params) GetInt(k string) (int64, error) {
	v := p.Get(k)

	if v == "" {
		return 0, nil
	}

	return strconv.ParseInt(v, 10, 64)
}

// GetUint attempts to get the given key as uint64.
func (p Params) GetUint(k string) (uint64, error) {
	v := p.Get(k)

	if v == "" {
		return 0, nil
	}

	return strconv.ParseUint(v, 10, 64)
}

// Scan implements the Scanner interface.
func (p Params) Scan(k string, dst interface{}) error {
	switch v := dst.(type) {
	case *string:
		*v = p.Get(k)

		return nil

	case *[]byte:
		*v = []byte(p.Get(k))

		return nil
	}

	dpv := reflect.ValueOf(dst)

	if dpv.Kind() != reflect.Ptr {
		return ErrValueDestination
	}

	dv := reflect.Indirect(dpv)

	switch dv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		c, err := p.GetInt(k)

		if err != nil {
			return err
		}

		dv.SetInt(c)

		return nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		c, err := p.GetUint(k)

		if err != nil {
			return err
		}

		dv.SetUint(c)

		return nil

	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(p.Get(k), dv.Type().Bits())

		if err != nil {
			return err
		}

		dv.SetFloat(f)

		return nil
	}

	// Utilize Scanner interface
	if s, ok := dst.(Scanner); ok {
		return s.Scan(p.Get(k))
	}

	return ErrUnsupportedType
}

// Scanner is an interface used by Scan.
type Scanner interface {
	// An error should be returned if the value can not be stored
	// without loss of information.
	Scan(src interface{}) error
}
