// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestArray(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [array? []]]
           [set $.asm.b [array? [a 2]]]
           [set $.asm.c [array? 3]]
         ]`,
		"{src: []}",
	)
	tt.Equal(t, "{a:true b:true c:false}", sen.String(root["asm"], &sopt))
}

func TestArrayArgCount(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"array?", 1, 2},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
