// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestSizeString(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm [size a_string]]
         ]`,
		"{src: []}",
	)
	tt.Equal(t, "8", sen.String(root["asm"]))
}

func TestSizeArray(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm [size [1 2 3]]]
         ]`,
		"{src: []}",
	)
	tt.Equal(t, "3", sen.String(root["asm"]))
}

func TestSizeMap(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm [size {a:1 b:2 c:3}]]
         ]`,
		"{src: []}",
	)
	tt.Equal(t, "3", sen.String(root["asm"]))
}

func TestSizeOther(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm [size true]]
         ]`,
		"{src: []}",
	)
	tt.Equal(t, "0", sen.String(root["asm"]))
}

func TestSizeArgCount(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"size", 1, 2},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
