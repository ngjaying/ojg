// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestLtInt(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [lt 2 "$.src[2]" 4]]
           [set $.asm.b [lt 2 1]]
         ]`,
		"{src: [1 2 3]}",
	)
	tt.Equal(t, "{a:true b:false}", sen.String(root["asm"], &sopt))
}

func TestLtFloat(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a ["<" 1.1 "$.src[1]" 3.3]]
           [set $.asm.b ["<" 2.0 1.0]]
         ]`,
		"{src: [1.1 2.2]}",
	)
	tt.Equal(t, "{a:true b:false}", sen.String(root["asm"], &sopt))
}

func TestLtString(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [lt abc "$.src[1]" zz]]
           [set $.asm.b [lt def abc]]
         ]`,
		"{src: [abc xyz]}",
	)
	tt.Equal(t, "{a:true b:false}", sen.String(root["asm"], &sopt))
}

func TestLtWrongType(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"set", "$.asm.i", []interface{}{"lt", true, false}},
	})
	root := map[string]interface{}{}
	err := p.Execute(root)
	tt.NotNil(t, err)
}

func TestLtWrongType2(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"set", "$.asm.i", []interface{}{"lt", 1, false}},
	})
	root := map[string]interface{}{}
	err := p.Execute(root)
	tt.NotNil(t, err)
}
