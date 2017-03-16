package main

import (
	"fmt"

	"github.com/hotgo/gobuf"
)

func genUnmarshaler(o *writer, name string, t *gobuf.Type, n int) {
	if genArrayUnmarshaler(o, name, t, n) {
		return
	}

	if genMapUnmarshaler(o, name, t, n) {
		return
	}

	if genPointerUnmarshaler(o, name, t, n) {
		return
	}

	genScalarUnmarshaler(o, name, t)
}

func genArrayUnmarshaler(o *writer, name string, t *gobuf.Type, n int) bool {
	if t.Kind == gobuf.ARRAY {
		o.Writef("{")
		o.Writef("	l, x := binary.Uvarint(b[n:])")
		o.Writef("	n += x")
		o.Writef("	%s = make(%s, l)", name, typeName(t))
		o.Writef("	for i%d := 0; i%d < int(l); i%d ++ {", n, n, n)
		genUnmarshaler(o, fmt.Sprintf("%s[i%d]", name, n), t.Elem, n+1)
		o.Writef("	}")
		o.Writef("}")
		return true
	}
	return false
}

func genMapUnmarshaler(o *writer, name string, t *gobuf.Type, n int) bool {
	if t.Kind == gobuf.MAP {
		o.Writef("{")
		o.Writef("	l, x := binary.Uvarint(b[n:])")
		o.Writef("	n += x")
		o.Writef("	%s = make(%s, l)", name, typeName(t))
		o.Writef("	for i%d := 0; i%d < int(l); i%d ++ {", n, n, n)
		o.Writef("		var key%d %s", n, typeName(t.Key))
		o.Writef("		var val%d %s", n, typeName(t.Elem))
		genScalarUnmarshaler(o, fmt.Sprintf("key%d", n), t.Key)
		genUnmarshaler(o, fmt.Sprintf("val%d", n), t.Elem, n+1)
		o.Writef("		%s[key%d] = val%d", name, n, n)
		o.Writef("	}")
		o.Writef("}")
	}
	return false
}

func genPointerUnmarshaler(o *writer, name string, t *gobuf.Type, n int) bool {
	if t.Kind == gobuf.POINTER {
		o.Writef("if b[n] != 0 {")
		o.Writef("	n += 1")
		o.Writef("	val%d := new(%s)", n, typeName(t.Elem))
		genScalarUnmarshaler(o, fmt.Sprintf("*val%d", n), t.Elem)
		o.Writef("	%s = val%d", name, n)
		o.Writef("} else {")
		o.Writef("	n += 1")
		o.Writef("}")
		return true
	}
	return false
}

func genScalarUnmarshaler(o *writer, name string, t *gobuf.Type) {
	switch t.Kind {
	case gobuf.INT8, gobuf.UINT8:
		o.Writef("%s = %s(b[n])", name, typeName(t))
		o.Writef("n += 1")
	case gobuf.INT16, gobuf.UINT16:
		o.Writef("%s = %s(binary.LittleEndian.Uint16(b[n:]))", name, typeName(t))
		o.Writef("n += 2")
	case gobuf.INT32, gobuf.UINT32:
		o.Writef("%s = %s(binary.LittleEndian.Uint32(b[n:]))", name, typeName(t))
		o.Writef("n += 4")
	case gobuf.INT64, gobuf.UINT64:
		o.Writef("%s = %s(binary.LittleEndian.Uint64(b[n:]))", name, typeName(t))
		o.Writef("n += 8")
	case gobuf.FLOAT32:
		o.Writef("%s = %s(gobuf.GetFloat32(b[n:]))", name, typeName(t))
		o.Writef("n += 4")
	case gobuf.FLOAT64:
		o.Writef("%s = %s(gobuf.GetFloat64(b[n:]))", name, typeName(t))
		o.Writef("n += 8")
	case gobuf.INT:
		o.Writef("{")
		o.Writef("	v, x := binary.Varint(b[n:])")
		o.Writef("	%s = %s(v)", name, typeName(t))
		o.Writef("	n += x")
		o.Writef("}")
	case gobuf.UINT:
		o.Writef("{")
		o.Writef("	v, x := binary.Uvarint(b[n:])")
		o.Writef("	%s = %s(v)", name, typeName(t))
		o.Writef("	n += x")
		o.Writef("}")
	case gobuf.BYTES:
		o.Writef("{")
		o.Writef("	l, x := binary.Uvarint(b[n:])")
		o.Writef("	n += x")
		o.Writef("	%s = make([]byte, l)", name)
		o.Writef("	copy(%s, b[n:n+int(l)])", name)
		o.Writef("	n += int(l)")
		o.Writef("}")
	case gobuf.STRING:
		o.Writef("{")
		o.Writef("	l, x := binary.Uvarint(b[n:])")
		o.Writef("	n += x")
		o.Writef("	%s = string(b[n:n+int(l)])", name)
		o.Writef("	n += int(l)")
		o.Writef("}")
	case gobuf.MESSAGE:
		if name[0] == '*' {
			name = name[1:]
		}
		o.Writef("n += %s.Unmarshal(b[n:])", name)
	}
}
