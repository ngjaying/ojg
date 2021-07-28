// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestNull(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [null? null]]
           [set $.asm.b [null? a_string]]
         ]`,
		"{src: []}",
	)
	tt.Equal(t, "{a:true b:false}", sen.String(root["asm"], &sopt))
}

func TestNullArgCount(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"null?", 1, 2},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
