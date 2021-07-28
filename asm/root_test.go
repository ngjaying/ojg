// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestRoot(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm [get [root src x]]]
         ]`,
		"{src: {x:3}}",
	)
	tt.Equal(t, `3`, sen.String(root["asm"], &sopt))
}

func TestRootArgNotString(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"root", 1},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestRootArgParseError(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"root", "[[["},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
