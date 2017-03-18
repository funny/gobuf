package main

import (
	"fmt"

	"github.com/funny/gobuf/gb"
)

func genMarshaler(o *writer, name string, t *gb.Type, n int) {
	if genArrayMarshaler(o, name, t, n) {
		return
	}

	if genMapMarshaler(o, name, t, n) {
		return
	}

	if genPointerMarshaler(o, name, t) {
		return
	}

	genScalarMarshaler(o, name, t)
}

func genArrayMarshaler(o *writer, name string, t *gb.Type, n int) bool {
	if t.Kind == gb.ARRAY {
		o.Writef("n += binary.PutUvarint(b[n:], uint64(len(%s)))", name)
		o.Writef("for i%d := 0; i%d < len(%s); i%d ++ {", n, n, name, n)
		genMarshaler(o, fmt.Sprintf("%s[i%d]", name, n), t.Elem, n+1)
		o.Writef("}")
		return true
	}
	return false
}

func genMapMarshaler(o *writer, name string, t *gb.Type, n int) bool {
	if t.Kind == gb.MAP {
		o.Writef("n += binary.PutUvarint(b[n:], uint64(len(%s)))", name)
		o.Writef("for key%d, val%d := range %s {", n, n, name)
		genScalarMarshaler(o, fmt.Sprintf("key%d", n), t.Key)
		genMarshaler(o, fmt.Sprintf("val%d", n), t.Elem, n+1)
		o.Writef("}")
	}
	return false
}

func genPointerMarshaler(o *writer, name string, t *gb.Type) bool {
	if t.Kind == gb.POINTER {
		o.Writef("if %s != nil {", name)
		o.Writef("	b[n] = 1")
		o.Writef("	n ++")
		genScalarMarshaler(o, "*"+name, t.Elem)
		o.Writef("} else {")
		o.Writef("	b[n] = 0")
		o.Writef("	n ++")
		o.Writef("}")
		return true
	}
	return false
}

func genScalarMarshaler(o *writer, name string, t *gb.Type) {
	switch t.Kind {
	case gb.INT8, gb.UINT8:
		o.Writef("b[n] = byte(%s)", name)
		o.Writef("n += 1")
	case gb.INT16, gb.UINT16:
		o.Writef("binary.LittleEndian.PutUint16(b[n:], uint16(%s))", name)
		o.Writef("n += 2")
	case gb.INT32, gb.UINT32:
		o.Writef("binary.LittleEndian.PutUint32(b[n:], uint32(%s))", name)
		o.Writef("n += 4")
	case gb.INT64, gb.UINT64:
		o.Writef("binary.LittleEndian.PutUint64(b[n:], uint64(%s))", name)
		o.Writef("n += 8")
	case gb.FLOAT32:
		o.Writef("gb.PutFloat32(b[n:], float32(%s))", name)
		o.Writef("n += 4")
	case gb.FLOAT64:
		o.Writef("gb.PutFloat64(b[n:], float64(%s))", name)
		o.Writef("n += 8")
	case gb.INT:
		o.Writef("n += binary.PutVarint(b[n:], int64(%s))", name)
	case gb.UINT:
		o.Writef("n += binary.PutUvarint(b[n:], uint64(%s))", name)
	case gb.BYTES, gb.STRING:
		o.Writef("n += binary.PutUvarint(b[n:], uint64(len(%s)))", name)
		o.Writef("copy(b[n:], %s)", name)
		o.Writef("n += len(%s)", name)
	case gb.STRUCT:
		if name[0] == '*' {
			name = name[1:]
		}
		o.Writef("n += %s.Marshal(b[n:])", name)
	}
}
