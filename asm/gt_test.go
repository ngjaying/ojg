// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestGtInt(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [gt 4 "$.src[2]" 1]]
           [set $.asm.b [gt 1 2]]
         ]`,
		"{src: [1 2 3]}",
	)
	tt.Equal(t, "{a:true b:false}", sen.String(root["asm"], &sopt))
}

func TestGtFloat(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [">" 4.1 "$.src[1]" 1]]
           [set $.asm.b [">" 1.0 2.0]]
         ]`,
		"{src: [1.1 2.2]}",
	)
	tt.Equal(t, "{a:true b:false}", sen.String(root["asm"], &sopt))
}

func TestGtString(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [gt xyz "$.src[0]" aa]]
           [set $.asm.b [gt abc def]]
         ]`,
		"{src: [abc xyz]}",
	)
	tt.Equal(t, "{a:true b:false}", sen.String(root["asm"], &sopt))
}

func TestGtIntOthers(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"set", "$.asm.i", []interface{}{"gt", 9, int8(8), int16(7), int32(6), int64(5)}},
		[]interface{}{"set", "$.asm.u", []interface{}{"gt", uint(9), uint8(8), uint16(7), uint32(6), uint64(5)}},
	})
	root := map[string]interface{}{
		"src": []interface{}{},
	}
	err := p.Execute(root)
	tt.Nil(t, err)

	tt.Equal(t, "{i:true u:true}", sen.String(root["asm"], &sopt))
}

func TestGtWrongType(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"set", "$.asm.i", []interface{}{"gt", true, false}},
	})
	root := map[string]interface{}{}
	err := p.Execute(root)
	tt.NotNil(t, err)
}

func TestGtWrongType2(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"set", "$.asm.i", []interface{}{"gt", 1, false}},
	})
	root := map[string]interface{}{}
	err := p.Execute(root)
	tt.NotNil(t, err)
}
