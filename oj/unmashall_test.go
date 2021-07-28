// Copyright (c) 2021, Peter Ohler, All rights reserved.

package oj_test

import (
	"testing"

	"github.com/ngjaying/ojg/alt"
	"github.com/ngjaying/ojg/oj"
	"github.com/ngjaying/ojg/tt"
)

func TestUnmarshal(t *testing.T) {
	var obj map[string]interface{}
	src := `{"x":3}`
	err := oj.Unmarshal([]byte(src), &obj)
	tt.Nil(t, err)
	tt.Equal(t, src, oj.JSON(obj))
	tt.Equal(t, 3.0, obj["x"])

	obj = nil
	p := oj.Parser{}
	err = p.Unmarshal([]byte(src), &obj)
	tt.Nil(t, err)
	tt.Equal(t, src, oj.JSON(obj))

	obj = nil
	err = oj.Unmarshal([]byte(src), &obj, &alt.Recomposer{})
	tt.Nil(t, err)
	tt.Equal(t, src, oj.JSON(obj))
}
