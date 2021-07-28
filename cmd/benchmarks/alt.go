// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"testing"
	"time"

	"github.com/ngjaying/ojg/alt"
)

func altGenerify(b *testing.B) {
	tm := time.Date(2020, time.April, 12, 16, 34, 04, 123456789, time.UTC)
	for n := 0; n < b.N; n++ {
		native := benchmarkData(tm)
		_ = alt.Generify(native)
	}
}

func altGenAlter(b *testing.B) {
	tm := time.Date(2020, time.April, 12, 16, 34, 04, 123456789, time.UTC)
	for n := 0; n < b.N; n++ {
		native := benchmarkData(tm)
		_ = alt.GenAlter(native)
	}
}
