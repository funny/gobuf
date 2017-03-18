package main

import (
	"fmt"

	"github.com/funny/gobuf/parser"
)

func genMarshaler(o *writer, name string, t *parser.Type, n int) {
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

func genArrayMarshaler(o *writer, name string, t *parser.Type, n int) bool {
	if t.Kind == parser.ARRAY {
		if t.Len == 0 {
			o.Writef("n += binary.PutUvarint(b[n:], uint64(len(%s)))", name)
		}
		o.Writef("for i%d := 0; i%d < len(%s); i%d ++ {", n, n, name, n)
		genMarshaler(o, fmt.Sprintf("%s[i%d]", name, n), t.Elem, n+1)
		o.Writef("}")
		return true
	}
	return false
}

func genMapMarshaler(o *writer, name string, t *parser.Type, n int) bool {
	if t.Kind == parser.MAP {
		o.Writef("n += binary.PutUvarint(b[n:], uint64(len(%s)))", name)
		o.Writef("for key%d, val%d := range %s {", n, n, name)
		genScalarMarshaler(o, fmt.Sprintf("key%d", n), t.Key)
		genMarshaler(o, fmt.Sprintf("val%d", n), t.Elem, n+1)
		o.Writef("}")
	}
	return false
}

func genPointerMarshaler(o *writer, name string, t *parser.Type) bool {
	if t.Kind == parser.POINTER {
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

func genScalarMarshaler(o *writer, name string, t *parser.Type) {
	switch t.Kind {
	case parser.BOOL:
		o.Writef("if %s { b[n] = 1; } else { b[n] = 0; }", name)
		o.Writef("n += 1")
	case parser.INT8, parser.UINT8:
		o.Writef("b[n] = byte(%s)", name)
		o.Writef("n += 1")
	case parser.INT16, parser.UINT16:
		o.Writef("binary.LittleEndian.PutUint16(b[n:], uint16(%s))", name)
		o.Writef("n += 2")
	case parser.INT32, parser.UINT32:
		o.Writef("binary.LittleEndian.PutUint32(b[n:], uint32(%s))", name)
		o.Writef("n += 4")
	case parser.INT64, parser.UINT64:
		o.Writef("binary.LittleEndian.PutUint64(b[n:], uint64(%s))", name)
		o.Writef("n += 8")
	case parser.FLOAT32:
		o.Writef("$name$_PutFloat32(b[n:], float32(%s))", name)
		o.Writef("n += 4")
	case parser.FLOAT64:
		o.Writef("$name$_PutFloat64(b[n:], float64(%s))", name)
		o.Writef("n += 8")
	case parser.INT:
		o.Writef("n += binary.PutVarint(b[n:], int64(%s))", name)
	case parser.UINT:
		o.Writef("n += binary.PutUvarint(b[n:], uint64(%s))", name)
	case parser.BYTES, parser.STRING:
		o.Writef("n += binary.PutUvarint(b[n:], uint64(len(%s)))", name)
		o.Writef("copy(b[n:], %s)", name)
		o.Writef("n += len(%s)", name)
	case parser.STRUCT:
		if name[0] == '*' {
			name = name[1:]
		}
		o.Writef("n += %s.Marshal(b[n:])", name)
	}
}
