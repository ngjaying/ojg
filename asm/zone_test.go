// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm_test

import (
	"testing"
	"time"

	"github.com/ngjaying/ojg/asm"
	"github.com/ngjaying/ojg/sen"
	"github.com/ngjaying/ojg/tt"
)

func TestZone(t *testing.T) {
	root := testPlan(t,
		`[
           [set $.asm.a [zone [time "2021-02-09T01:02:03Z"] "EST"]]
           [set $.asm.b [zone [time "2021-02-09T01:02:03-05:00"] "UTC"]]
           [set $.asm.c [zone [time "2021-02-09T01:02:03Z"] -18000]]
           [set $.asm.d [zone [time "2021-02-09T01:02:03Z"] "America/Toronto"]]
           [set $.asm.e [zone [time "2021-02-09T01:02:03Z"] "Unknown"]]
         ]`,
		"{src: []}",
	)
	opt := sopt
	opt.Indent = 2
	// Note the golang float64 does not have enough precision to represent a
	// time with nonoseconds.
	tt.Equal(t,
		`{
  a: "2021-02-08T20:02:03-05:00"
  b: "2021-02-09T06:02:03Z"
  c: "2021-02-08T20:02:03-05:00"
  d: "2021-02-08T20:02:03-05:00"
  e: "2021-02-09T01:02:03Z"
}`, sen.String(root["asm"], &opt))
}

func TestZoneArgCount(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"zone", 1, 2, 3},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestZoneNotTime(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"zone", 1, 2},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}

func TestZoneNotLoc(t *testing.T) {
	p := asm.NewPlan([]interface{}{
		[]interface{}{"zone", time.Now(), true},
	})
	err := p.Execute(map[string]interface{}{})
	tt.NotNil(t, err)
}
