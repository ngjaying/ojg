// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt_test

import (
	"testing"
	"time"

	"github.com/ngjaying/ojg/alt"
	"github.com/ngjaying/ojg/tt"
)

func TestMatchInt(t *testing.T) {
	tt.Equal(t, true, alt.Match(map[string]interface{}{"x": 1}, map[string]interface{}{"x": 1, "y": 2}))
	tt.Equal(t, true, alt.Match(map[string]interface{}{"x": nil}, map[string]interface{}{"x": nil, "y": 2}))
	tt.Equal(t, true, alt.Match(map[string]interface{}{"x": nil}, map[string]interface{}{"y": 2}))
	tt.Equal(t, false, alt.Match(map[string]interface{}{"x": nil}, map[string]interface{}{"x": 1, "y": 2}))
	tt.Equal(t, false, alt.Match(map[string]interface{}{"x": 1, "z": 3}, map[string]interface{}{"x": 1, "y": 2}))
}

func TestMatchBool(t *testing.T) {
	tt.Equal(t, true, alt.Match(map[string]interface{}{"x": true}, map[string]interface{}{"x": true, "y": 2}))
	tt.Equal(t, false, alt.Match(map[string]interface{}{"x": true}, map[string]interface{}{"x": false, "y": 2}))
}

func TestMatchFloat(t *testing.T) {
	tt.Equal(t, true, alt.Match(map[string]interface{}{"x": 1.5}, map[string]interface{}{"x": 1.5, "y": 2}))
	tt.Equal(t, false, alt.Match(map[string]interface{}{"x": 1.5}, map[string]interface{}{"x": 2.5, "y": 2}))
}

func TestMatchString(t *testing.T) {
	tt.Equal(t, true, alt.Match(map[string]interface{}{"x": "a"}, map[string]interface{}{"x": "a"}))
	tt.Equal(t, false, alt.Match(map[string]interface{}{"x": "a"}, map[string]interface{}{"x": "b"}))
}

func TestMatchTime(t *testing.T) {
	tm := time.Date(2020, time.April, 12, 16, 34, 04, 123456789, time.UTC)
	tt.Equal(t, true, alt.Match(map[string]interface{}{"x": tm}, map[string]interface{}{"x": tm}))
	tt.Equal(t, false, alt.Match(map[string]interface{}{"x": tm}, map[string]interface{}{"x": "b"}))
}

func TestMatchSlice(t *testing.T) {
	tt.Equal(t, true,
		alt.Match(map[string]interface{}{"x": []interface{}{1}}, map[string]interface{}{"x": []interface{}{1}}))
	tt.Equal(t, false,
		alt.Match(map[string]interface{}{"x": []interface{}{1}}, map[string]interface{}{"x": []interface{}{2}}))
	tt.Equal(t, false,
		alt.Match(map[string]interface{}{"x": []interface{}{1}}, map[string]interface{}{"x": []interface{}{1, 2}}))
}

func TestMatchMap(t *testing.T) {
	tt.Equal(t, false, alt.Match(map[string]interface{}{"x": []interface{}{1}}, 7))
}

func TestMatchStruct(t *testing.T) {
	type Sample struct {
		Int int
	}
	type Dample struct {
		Int int
	}
	tt.Equal(t, true, alt.Match(&Sample{Int: 3}, &Sample{Int: 3}))
	tt.Equal(t, false, alt.Match(&Sample{Int: 3}, &Dample{Int: 3}))
}

func TestMatchSimplify(t *testing.T) {
	tt.Equal(t, true, alt.Match(&silly{val: 3}, &silly{val: 3}))
}
