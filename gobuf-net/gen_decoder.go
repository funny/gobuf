package main

import (
	"fmt"

	"github.com/funny/gobuf"
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
		o.Writef("	ulong l;")
		o.Writef("	n = Gobuf.ReadUvarint(out l, b, n);")
		o.Writef("	%s = new System.Collections.Generic.List<%s>();", name, typeName(t.Elem))
		o.Writef("	for (var i%d = 0; i%d < l; i%d ++) {", n, n, n)
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
		o.Writef("	ulong l;")
		o.Writef("	n = Gobuf.ReadUvarint(out l, b, n);")
		o.Writef("	%s = System.Collections.Generic.Dictionary<%s, %s>();", name, typeName(t.Key), typeName(t.Elem))
		o.Writef("	for (var i%d = 0; i%d < l; i%d ++) {", n, n, n)
		o.Writef("		%s key%d;", typeName(t.Key), n)
		o.Writef("		%s val%d;", typeName(t.Elem), n)
		genScalarUnmarshaler(o, fmt.Sprintf("key%d", n), t.Key)
		genUnmarshaler(o, fmt.Sprintf("val%d", n), t.Elem, n+1)
		o.Writef("		%s[key%d] = val%d;", name, n, n)
		o.Writef("	}")
		o.Writef("}")
	}
	return false
}

func genPointerUnmarshaler(o *writer, name string, t *gobuf.Type, n int) bool {
	if t.Kind == gobuf.POINTER {
		o.Writef("if (b[n] != 0) {")
		o.Writef("	n += 1;")
		o.Writef("	var val%d = new %s();", n, typeName(t.Elem))
		genScalarUnmarshaler(o, fmt.Sprintf("val%d", n), t.Elem)
		o.Writef("	%s = val%d;", name, n)
		o.Writef("} else {")
		o.Writef("	n += 1;")
		o.Writef("}")
		return true
	}
	return false
}

func genScalarUnmarshaler(o *writer, name string, t *gobuf.Type) {
	switch t.Kind {
	case gobuf.INT8, gobuf.UINT8:
		o.Writef("%s = (%s)b[n];", name, typeName(t))
		o.Writef("n += 1;")
	case gobuf.INT16, gobuf.UINT16:
		o.Writef("n = (%s)Gobuf.ReadUint16(out %s, b, n);", typeName(t), name)
	case gobuf.INT32, gobuf.UINT32:
		o.Writef("n = (%s)Gobuf.ReadUint32(out %s, b, n);", typeName(t), name)
	case gobuf.INT64, gobuf.UINT64:
		o.Writef("n = (%s)Gobuf.ReadUint64(out %s, b, n);", typeName(t), name)
	case gobuf.FLOAT32:
		o.Writef("n = (%s)Gobuf.ReadFloat32(out %s, b, n);", typeName(t), name)
	case gobuf.FLOAT64:
		o.Writef("n = (%s)Gobuf.ReadFloat64(out %s, b, n);", typeName(t), name)
	case gobuf.INT:
		o.Writef("n = (%s)Gobuf.ReadVarint(out %s, b, n);", typeName(t), name)
	case gobuf.UINT:
		o.Writef("n = (%s)Gobuf.ReadUvarint(out %s, b, n);", typeName(t), name)
	case gobuf.BYTES:
		o.Writef("n = (%s)Gobuf.ReadBytes(out %s, b, n);", typeName(t), name)
	case gobuf.STRING:
		o.Writef("n = (%s)Gobuf.ReadString(out %s, b, n);", typeName(t), name)
	case gobuf.MESSAGE:
		o.Writef("n = %s.Unmarshal(b, n);", name)
	}
}
