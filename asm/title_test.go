// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestTitle(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [title low]]
           [set $.asm.b [title UP]]
         ]`,
		"{src: []}",
	)
	opt := sopt
	opt.Indent = 2
	tt.Equal(t,
		`{
  a: Low
  b: UP
}`, sen.String(root["asm"], &opt))
}

func TestTitleArgCount(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"title", "x", "y"},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestTitleArgType(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"title", 1},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
