// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestGet(t *testing.T) {
	root := testPlan(t,
		`[
           {one:1 two:2}
           [set $.at [get @.one]]
           [set $.root [get $.src.b]]
           [set $.arg [get @.x {x:1 y:2}]]
         ]`,
		"{src: {a:1 b:2 c:3}}",
	)
	tt.Equal(t, "1", sen.String(root["at"]))
	tt.Equal(t, "2", sen.String(root["root"]))
	tt.Equal(t, "1", sen.String(root["arg"]))
}

func TestGetArgCount(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"get"},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestGetArgNotExpr(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"get", 1},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestGetArgType(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"get", []interface{}{"sum"}},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
