// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp_test

import (
	"fmt"
	"testing"

	"github.com/ngjaying/ojg/gen"
	"github.com/ngjaying/ojg/jp"
	"github.com/ngjaying/ojg/oj"
	"github.com/ngjaying/ojg/tt"
)

type delData struct {
	path     string
	data     string // JSON
	expect   string // JSON
	err      string
	noNode   bool
	noSimple bool
}

var (
	delTestData = []*delData{
		{path: "@.a", data: `{}`, expect: `{}`},
		{path: "@.a", data: `{"a":3}`, expect: `{}`},
		{path: "[1]", data: `[1,2,3]`, expect: `[1,null,3]`},
		{path: "a.*", data: `{"a":{"x":1,"y":2}}`, expect: `{"a":{}}`},
		{path: "[*]", data: `[1,2,3]`, expect: `[null,null,null]`},

		{path: "", data: `{}`, err: "can not delete with an empty expression"},
		{path: "$", data: `{}`, err: "can not delete the root"},
		{path: "@", data: `{}`, err: "can not delete an empty expression"},
		{path: "a[1,2]", data: `{}`, err: "can not delete with an expression ending with a Union"},
		{path: "a.b", data: `{"a":4}`, err: "/can not follow a .+ at 'a'/"},
		{path: "a[0]", data: `{}`, err: "can not deduce what element to add at 'a'"},
		{path: "[0].1", data: `[1]`, err: "/can not follow a .+ at '\\[0\\]'/"},
		{path: "[1]", data: `[1]`, err: "can not follow out of bounds array index at '[1]'"},
	}
	delOneTestData = []*delData{
		{path: "@.a", data: `{}`, expect: `{}`},
		{path: "@.a", data: `{"a":3}`, expect: `{}`},
		{path: "[1]", data: `[1,2,3]`, expect: `[1,null,3]`},
		{path: "a.*", data: `{"a":{"x":1}}`, expect: `{"a":{}}`},
		{path: "[*]", data: `[1,2,3]`, expect: `[null,2,3]`},
		{path: "..a", data: `{"x":{"a":1,"b":2}}`, expect: `{"x":{"b":2}}`},

		{path: "", data: `{}`, err: "can not delete with an empty expression"},
		{path: "$", data: `{}`, err: "can not delete the root"},
		{path: "@", data: `{}`, err: "can not delete an empty expression"},
		{path: "a[1,2]", data: `{}`, err: "can not delete with an expression ending with a Union"},
		{path: "a.b", data: `{"a":4}`, err: "/can not follow a .+ at 'a'/"},
		{path: "a[0]", data: `{}`, err: "can not deduce what element to add at 'a'"},
		{path: "[0].1", data: `[1]`, err: "/can not follow a .+ at '\\[0\\]'/"},
		{path: "[1]", data: `[1]`, err: "can not follow out of bounds array index at '[1]'"},
	}
)

func TestExprDel(t *testing.T) {
	for i, d := range delTestData {
		if testing.Verbose() {
			fmt.Printf("... %d: %s\n", i, d.path)
		}
		x, err := jp.ParseString(d.path)
		tt.Nil(t, err, i, " : ", x)

		var data interface{}
		if !d.noSimple {
			data, err = oj.ParseString(d.data)
			tt.Nil(t, err, i, " : ", x)
			err = x.Del(data)
			if 0 < len(d.err) {
				tt.NotNil(t, err, i, " : ", x)
				tt.Equal(t, d.err, err.Error(), i, " : ", x)
			} else {
				result := oj.JSON(data, &oj.Options{Sort: true})
				tt.Equal(t, d.expect, result, i, " : ", x)
			}
		}
		if !d.noNode {
			var p gen.Parser
			data, err = p.Parse([]byte(d.data))
			tt.Nil(t, err, i, " : ", x)
			err = x.Del(data)
			if 0 < len(d.err) {
				tt.NotNil(t, err, i, " : ", x)
				tt.Equal(t, d.err, err.Error(), i, " : ", x)
			} else {
				result := oj.JSON(data, &oj.Options{Sort: true})
				tt.Equal(t, d.expect, result, i, " : ", x)
			}
		}
	}
}

func TestExprDelOne(t *testing.T) {
	for i, d := range delOneTestData {
		if testing.Verbose() {
			fmt.Printf("... %d: %s\n", i, d.path)
		}
		x, err := jp.ParseString(d.path)
		tt.Nil(t, err, i, " : ", x)

		var data interface{}
		if !d.noSimple {
			data, err = oj.ParseString(d.data)
			tt.Nil(t, err, i, " : ", x)
			err = x.DelOne(data)
			if 0 < len(d.err) {
				tt.NotNil(t, err, i, " : ", x)
				tt.Equal(t, d.err, err.Error(), i, " : ", x)
			} else {
				result := oj.JSON(data, &oj.Options{Sort: true})
				tt.Equal(t, d.expect, result, i, " : ", x)
			}
		}
		if !d.noNode {
			var p gen.Parser
			data, err = p.Parse([]byte(d.data))
			tt.Nil(t, err, i, " : ", x)
			err = x.DelOne(data)
			if 0 < len(d.err) {
				tt.NotNil(t, err, i, " : ", x)
				tt.Equal(t, d.err, err.Error(), i, " : ", x)
			} else {
				result := oj.JSON(data, &oj.Options{Sort: true})
				tt.Equal(t, d.expect, result, i, " : ", x)
			}
		}
	}
}
