// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestReverse(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [reverse [a b c]]]
           [set $.asm.b [reverse [1 b 3]]]
         ]`,
		"{src: []}",
	)
	tt.Equal(t, `{a:[c b a] b:[3 b 1]}`, sen.String(root["asm"], &sopt))
}

func TestReverseArgCount(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"reverse", []interface{}{}, 1},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestReverseArgType(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"reverse", 1},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
