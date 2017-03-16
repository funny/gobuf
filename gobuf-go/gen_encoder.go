package main

import (
	"fmt"

	"github.com/hotgo/gobuf"
)

func genMarshaler(o *writer, name string, t *gobuf.Type, n int) {
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

func genArrayMarshaler(o *writer, name string, t *gobuf.Type, n int) bool {
	if t.Kind == gobuf.ARRAY {
		o.Writef("n += binary.PutUvarint(b[n:], uint64(len(%s)))", name)
		o.Writef("for i%d := 0; i%d < len(%s); i%d ++ {", n, n, name, n)
		genMarshaler(o, fmt.Sprintf("%s[i%d]", name, n), t.Elem, n+1)
		o.Writef("}")
		return true
	}
	return false
}

func genMapMarshaler(o *writer, name string, t *gobuf.Type, n int) bool {
	if t.Kind == gobuf.MAP {
		o.Writef("n += binary.PutUvarint(b[n:], uint64(len(%s)))", name)
		o.Writef("for key%d, val%d := range %s {", n, n, name)
		genScalarMarshaler(o, fmt.Sprintf("key%d", n), t.Key)
		genMarshaler(o, fmt.Sprintf("val%d", n), t.Elem, n+1)
		o.Writef("}")
	}
	return false
}

func genPointerMarshaler(o *writer, name string, t *gobuf.Type) bool {
	if t.Kind == gobuf.POINTER {
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

func genScalarMarshaler(o *writer, name string, t *gobuf.Type) {
	switch t.Kind {
	case gobuf.INT8, gobuf.UINT8:
		o.Writef("b[n] = byte(%s)", name)
		o.Writef("n += 1")
	case gobuf.INT16, gobuf.UINT16:
		o.Writef("binary.LittleEndian.PutUint16(b[n:], uint16(%s))", name)
		o.Writef("n += 2")
	case gobuf.INT32, gobuf.UINT32:
		o.Writef("binary.LittleEndian.PutUint32(b[n:], uint32(%s))", name)
		o.Writef("n += 4")
	case gobuf.INT64, gobuf.UINT64:
		o.Writef("binary.LittleEndian.PutUint64(b[n:], uint64(%s))", name)
		o.Writef("n += 8")
	case gobuf.FLOAT32:
		o.Writef("gobuf.PutFloat32(b[n:], float32(%s))", name)
		o.Writef("n += 4")
	case gobuf.FLOAT64:
		o.Writef("gobuf.PutFloat64(b[n:], float64(%s))", name)
		o.Writef("n += 8")
	case gobuf.INT:
		o.Writef("n += binary.PutVarint(b[n:], int64(%s))", name)
	case gobuf.UINT:
		o.Writef("n += binary.PutUvarint(b[n:], uint64(%s))", name)
	case gobuf.BYTES, gobuf.STRING:
		o.Writef("n += binary.PutUvarint(b[n:], uint64(len(%s)))", name)
		o.Writef("copy(b[n:], %s)", name)
		o.Writef("n += len(%s)", name)
	case gobuf.MESSAGE:
		if name[0] == '*' {
			name = name[1:]
		}
		o.Writef("n += %s.Marshal(b[n:])", name)
	}
}
