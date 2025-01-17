// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestMod(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [mod 7 3]]
           [set $.asm.b [mod 6 3]]
         ]`,
		"{src: []}",
	)
	opt := sopt
	opt.Indent = 2
	tt.Equal(t,
		`{
  a: 1
  b: 0
}`, sen.String(root["asm"], &opt))
}

func TestModArgCount(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"mod", 1, 2, 3},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestModArgType(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"mod", 1, true},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestModArgType2(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"mod", true, 1},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
