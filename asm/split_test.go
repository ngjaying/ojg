// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestSplit(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [split "a b c" " "]]
           [set $.asm.b [split file-path-name "-"]]
         ]`,
		"{src: []}",
	)
	tt.Equal(t,
		`{a:[a b c] b:[file path name]}`, sen.String(root["asm"], &sopt))
}

func TestSplitArgCount(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"split", "x", "y", 1},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestSplitArgType(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"split", 1, "x"},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestSplitArgType2(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"split", "x", 1},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
