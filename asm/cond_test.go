// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestCond(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [cond [true abc]]]
           [set $.asm.b [cond [false abc][true def]]]
           [set $.asm.c [cond [1 abc][true def]]]
           [set $.asm.d [cond [1 abc][false def]]]
           [set $.asm.e [cond]]
         ]`,
		"{src: []}",
	)
	opt := sopt
	opt.Indent = 2
	tt.Equal(t,
		`{
  a: abc
  b: def
  c: def
  d: null
  e: null
}`, sen.String(root["asm"], &opt))
}

func TestCondArgType(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"cond", 1, "x"},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestCondArgElementCount(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"cond", []interface{}{true, 1, 2}},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
