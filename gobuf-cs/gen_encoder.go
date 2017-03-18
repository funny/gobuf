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
		o.Writef("n = Gobuf.WriteUvarint((ulong)%s.Length, b, n);", name)
		o.Writef("for (var i%d = 0; i%d < %s.Length; i%d ++) {", n, n, name, n)
		genMarshaler(o, fmt.Sprintf("%s[i%d]", name, n), t.Elem, n+1)
		o.Writef("}")
		return true
	}
	return false
}

func genMapMarshaler(o *writer, name string, t *gb.Type, n int) bool {
	if t.Kind == gb.MAP {
		o.Writef("n = Gobuf.WriteUvarint((ulong)%s.Length, b, n);", name)
		o.Writef("foreach (var item%d in %s) {", n, name)
		genScalarMarshaler(o, fmt.Sprintf("item%d.Key", n), t.Key)
		genMarshaler(o, fmt.Sprintf("item%d.Value", n), t.Elem, n+1)
		o.Writef("}")
	}
	return false
}

func genPointerMarshaler(o *writer, name string, t *gb.Type) bool {
	if t.Kind == gb.POINTER {
		o.Writef("if (%s != null) {", name)
		o.Writef("	b[n] = 1;")
		o.Writef("	n ++;")
		genScalarMarshaler(o, name, t.Elem)
		o.Writef("} else {")
		o.Writef("	b[n] = 0;")
		o.Writef("	n ++;")
		o.Writef("}")
		return true
	}
	return false
}

func genScalarMarshaler(o *writer, name string, t *gb.Type) {
	switch t.Kind {
	case gb.INT8, gb.UINT8:
		o.Writef("b[n] = (byte)%s;", name)
		o.Writef("n += 1;")
	case gb.INT16, gb.UINT16:
		o.Writef("n = Gobuf.WriteUint16((ushort)%s, b, n);", name)
	case gb.INT32, gb.UINT32:
		o.Writef("n = Gobuf.WriteUint32((long)%s, b, n);", name)
	case gb.INT64, gb.UINT64:
		o.Writef("n = Gobuf.WriteUint64((ulong)%s, b, n);", name)
	case gb.FLOAT32:
		o.Writef("n = Gobuf.WriteFloat32((float)%s, b, n);", name)
	case gb.FLOAT64:
		o.Writef("n = Gobuf.WriteFloat64((double)%s, b, n);", name)
	case gb.INT:
		o.Writef("n = Gobuf.WriteVarint((long)%s, b, n);", name)
	case gb.UINT:
		o.Writef("n = Gobuf.WriteUvarint((ulong)%s, b, n);", name)
	case gb.BYTES:
		o.Writef("n = Gobuf.WriteBytes(%s, b, n);", name)
	case gb.STRING:
		o.Writef("n = Gobuf.WriteString(%s, b, n);", name)
	case gb.STRUCT:
		o.Writef("n = %s.Marshal(b, n);", name)
	}
}
