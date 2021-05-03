// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"

	"github.com/ohler55/ojg"
)

type Options = ojg.Options
type Converter = ojg.Converter

var (
	DefaultOptions = ojg.DefaultOptions
	BrightOptions  = ojg.BrightOptions
	GoOptions      = ojg.GoOptions
	HTMLOptions    = ojg.HTMLOptions

	TimeRFC3339Converter = ojg.TimeRFC3339Converter
	TimeNanoConverter    = ojg.TimeNanoConverter
	MongoConverter       = ojg.MongoConverter
)

func init() {
	// Use different defaults for decompose except the Go defaults. Set
	// OmitNil and provide a CreateKey for all.
	DefaultOptions.OmitNil = true
	DefaultOptions.CreateKey = "type"
	BrightOptions.OmitNil = true
	BrightOptions.CreateKey = "type"
	HTMLOptions.OmitNil = true
	HTMLOptions.CreateKey = "type"
}

// Dup is an alias for Decompose.
func Dup(v interface{}, options ...*ojg.Options) interface{} {
	return Decompose(v, options...)
}

// Decompose creates a simple type converting non simple to simple types using
// either the Simplify() interface or reflection. Unlike Alter() a deep copy
// is returned leaving the original data unchanged.
func Decompose(v interface{}, options ...*ojg.Options) interface{} {
	opt := &DefaultOptions
	if 0 < len(options) {
		opt = options[0]
	}
	if opt.Converter != nil {
		v = opt.Converter.Convert(v)
	}
	return decompose(v, opt)
}

// Alter the data into all simple types converting non simple to simple types
// using either the Simplify() interface or reflection. Unlike Decompose() map and
// slices members are modified if necessary to assure all elements are simple
// types.
func Alter(v interface{}, options ...*ojg.Options) interface{} {
	opt := &DefaultOptions
	if 0 < len(options) {
		opt = options[0]
	}
	if opt.Converter != nil {
		v = opt.Converter.Convert(v)
	}
	return alter(v, opt)
}

// Recompose simple data into more complex go types.
func Recompose(v interface{}, tv ...interface{}) (out interface{}, err error) {
	return DefaultRecomposer.Recompose(v, tv...)
}

// NewRecomposer creates a new instance. The composers are a map of objects
// expected and functions to recompose them. If no function is provided then
// reflection is used instead.
func NewRecomposer(createKey string, composers map[interface{}]RecomposeFunc) (*Recomposer, error) {
	r := Recomposer{
		CreateKey: createKey,
		composers: map[string]*composer{},
	}
	for v, fun := range composers {
		rt := reflect.TypeOf(v)
		if _, err := r.registerComposer(rt, fun); err != nil {
			return nil, err
		}
	}
	return &r, nil
}
