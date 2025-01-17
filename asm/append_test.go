// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestAppend(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [append [] a]]
           [set $.asm.b [append [a] 1]]
         ]`,
		"{src: []}",
	)
	tt.Equal(t,
		`{a:[a] b:[a 1]}`, sen.String(root["asm"], &sopt))
}

func TestAppendArgCount(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"append", []interface{}{}, 1, 2},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestAppendArgType(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"append", 1, "x"},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
