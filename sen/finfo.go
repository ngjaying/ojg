// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

import (
	"encoding"
	"encoding/json"
	"reflect"
	"unsafe"

	"github.com/ngjaying/ojg"
	"github.com/ngjaying/ojg/alt"
)

const (
	strMask   = byte(0x01)
	omitMask  = byte(0x02)
	embedMask = byte(0x04)

	aJustKey appendStatus = iota
	aWrote
	aSkip
	aChanged
)

type appendStatus byte

type appendFunc func(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, appendStatus)

// Field hold information about a struct field.
type finfo struct {
	rt      reflect.Type
	key     string
	kind    reflect.Kind
	elem    *sinfo
	Append  appendFunc
	iAppend appendFunc
	jkey    []byte
	index   []int
	offset  uintptr
}

func (f *finfo) keyLen() int {
	return len(f.jkey)
}

func appendJustKey(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, appendStatus) {
	v := rv.FieldByIndex(fi.index).Interface()
	buf = append(buf, fi.jkey...)
	return buf, v, aJustKey
}

func appendPtrNotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, appendStatus) {
	v := rv.FieldByIndex(fi.index).Interface()
	if (*[2]uintptr)(unsafe.Pointer(&v))[1] == 0 { // real nil check
		return buf, nil, aSkip
	}
	buf = append(buf, fi.jkey...)
	return buf, v, aJustKey
}

func appendSliceNotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, appendStatus) {
	fv := rv.FieldByIndex(fi.index)
	if fv.Len() == 0 {
		return buf, nil, aSkip
	}
	buf = append(buf, fi.jkey...)
	return buf, fv.Interface(), aJustKey
}

func appendSENString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, appendStatus) {
	v := rv.FieldByIndex(fi.index).String()
	buf = append(buf, fi.jkey...)
	buf = ojg.AppendSENString(buf, v, safe)

	return buf, nil, aWrote
}

func appendSENStringNotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, interface{}, appendStatus) {
	s := rv.FieldByIndex(fi.index).String()
	if len(s) == 0 {
		return buf, nil, aSkip
	}
	buf = append(buf, fi.jkey...)
	buf = ojg.AppendSENString(buf, s, safe)

	return buf, nil, aWrote
}

func whichAppend(rt reflect.Type, omitEmpty bool) (f appendFunc) {
	v := reflect.New(rt).Elem().Interface()
	switch v.(type) {
	case json.Marshaler:
		if omitEmpty {
			f = appendJSONMarshalerNotEmpty
		} else {
			f = appendJSONMarshaler
		}
	case encoding.TextMarshaler:
		if omitEmpty {
			f = appendTextMarshalerNotEmpty
		} else {
			f = appendTextMarshaler
		}
	case alt.Simplifier:
		if omitEmpty {
			f = appendSimplifierNotEmpty
		} else {
			f = appendSimplifier
		}
	case alt.Genericer:
		if omitEmpty {
			f = appendGenericerNotEmpty
		} else {
			f = appendGenericer
		}
	}
	vp := reflect.New(rt).Interface()
	switch vp.(type) {
	case json.Marshaler:
		f = appendJSONMarshalerAddr
	case encoding.TextMarshaler:
		f = appendTextMarshalerAddr
	case alt.Simplifier:
		f = appendSimplifierAddr
	case alt.Genericer:
		f = appendGenericerAddr
	}
	return
}

func newFinfo(f reflect.StructField, key string, omitEmpty, asString, pretty, embedded bool) *finfo {
	fi := finfo{
		rt:     f.Type,
		key:    key,
		kind:   f.Type.Kind(),
		index:  f.Index,
		offset: f.Offset,
	}
	var fx byte
	// Check for interfaces first since almost any type can implement one of
	// the supported interfaces.
	af := whichAppend(fi.rt, omitEmpty)
	if af != nil {
		fi.Append = af
		fi.iAppend = af
		goto Key
	}
	if omitEmpty {
		fx |= omitMask
	}
	if asString {
		fx |= strMask
	}
	if embedded {
		fx |= embedMask
	}
	switch fi.kind {
	case reflect.Bool:
		fi.Append = boolAppendFuncs[fx]
		fi.iAppend = boolAppendFuncs[fx|embedMask]

	case reflect.Int:
		fi.Append = intAppendFuncs[fx]
		fi.iAppend = intAppendFuncs[fx|embedMask]
	case reflect.Int8:
		fi.Append = int8AppendFuncs[fx]
		fi.iAppend = int8AppendFuncs[fx|embedMask]
	case reflect.Int16:
		fi.Append = int16AppendFuncs[fx]
		fi.iAppend = int16AppendFuncs[fx|embedMask]
	case reflect.Int32:
		fi.Append = int32AppendFuncs[fx]
		fi.iAppend = int32AppendFuncs[fx|embedMask]
	case reflect.Int64:
		fi.Append = int64AppendFuncs[fx]
		fi.iAppend = int64AppendFuncs[fx|embedMask]

	case reflect.Uint:
		fi.Append = uintAppendFuncs[fx]
		fi.iAppend = uintAppendFuncs[fx|embedMask]
	case reflect.Uint8:
		fi.Append = uint8AppendFuncs[fx]
		fi.iAppend = uint8AppendFuncs[fx|embedMask]
	case reflect.Uint16:
		fi.Append = uint16AppendFuncs[fx]
		fi.iAppend = uint16AppendFuncs[fx|embedMask]
	case reflect.Uint32:
		fi.Append = uint32AppendFuncs[fx]
		fi.iAppend = uint32AppendFuncs[fx|embedMask]
	case reflect.Uint64:
		fi.Append = uint64AppendFuncs[fx]
		fi.iAppend = uint64AppendFuncs[fx|embedMask]

	case reflect.Float32:
		fi.Append = float32AppendFuncs[fx]
		fi.iAppend = float32AppendFuncs[fx|embedMask]
	case reflect.Float64:
		fi.Append = float64AppendFuncs[fx]
		fi.iAppend = float64AppendFuncs[fx|embedMask]

	case reflect.String:
		if omitEmpty {
			fi.Append = appendSENStringNotEmpty
			fi.iAppend = appendSENStringNotEmpty
		} else {
			fi.Append = appendSENString
			fi.iAppend = appendSENString
		}
	case reflect.Struct:
		fi.elem = getTypeStruct(fi.rt, true)
		fi.Append = appendJustKey
		fi.iAppend = appendJustKey
	case reflect.Ptr:
		et := fi.rt.Elem()
		if et.Kind() == reflect.Ptr {
			et = et.Elem()
		}
		if et.Kind() == reflect.Struct {
			fi.elem = getTypeStruct(et, false)
		}
		if omitEmpty {
			fi.Append = appendPtrNotEmpty
			fi.iAppend = appendPtrNotEmpty
		} else {
			fi.Append = appendJustKey
			fi.iAppend = appendJustKey
		}
	case reflect.Interface:
		if omitEmpty {
			fi.Append = appendPtrNotEmpty
			fi.iAppend = appendPtrNotEmpty
		} else {
			fi.Append = appendJustKey
			fi.iAppend = appendJustKey
		}
	case reflect.Slice, reflect.Array, reflect.Map:
		et := fi.rt.Elem()
		embedded := true
		if et.Kind() == reflect.Ptr {
			embedded = false
			et = et.Elem()
		}
		if et.Kind() == reflect.Struct {
			fi.elem = getTypeStruct(et, embedded)
		}
		if omitEmpty {
			fi.Append = appendSliceNotEmpty
			fi.iAppend = appendSliceNotEmpty
		} else {
			fi.Append = appendJustKey
			fi.iAppend = appendJustKey
		}
	}
Key:
	fi.jkey = ojg.AppendSENString(fi.jkey, fi.key, false)
	fi.jkey = append(fi.jkey, ':')
	if pretty {
		fi.jkey = append(fi.jkey, ' ')
	}
	return &fi
}
