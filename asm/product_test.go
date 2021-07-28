// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestProduct(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [product 1 2 3]]
           [set $.asm.b [* 2.2 2]]
           [set $.asm.c [* 2.2 2.5]]
           [set $.asm.d [product]]
           [set $.asm.e [product 1 2.3]]
         ]`,
		"{src: []}",
	)
	opt := sopt
	opt.Indent = 2
	tt.Equal(t,
		`{
  a: 6
  b: 4.4
  c: 5.5
  d: 0
  e: 2.3
}`, sen.String(root["asm"], &opt))
}

func TestProductArgType(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"product", 1, true},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
