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
		o.Writef("n = Gobuf.WriteUvarint((ulong)%s.Length, b, n);", name)
		o.Writef("for (var i%d = 0; i%d < %s.Length; i%d ++) {", n, n, name, n)
		genMarshaler(o, fmt.Sprintf("%s[i%d]", name, n), t.Elem, n+1)
		o.Writef("}")
		return true
	}
	return false
}

func genMapMarshaler(o *writer, name string, t *gobuf.Type, n int) bool {
	if t.Kind == gobuf.MAP {
		o.Writef("n = Gobuf.WriteUvarint((ulong)%s.Length, b, n);", name)
		o.Writef("foreach (var item%d in %s) {", n, name)
		genScalarMarshaler(o, fmt.Sprintf("item%d.Key", n), t.Key)
		genMarshaler(o, fmt.Sprintf("item%d.Value", n), t.Elem, n+1)
		o.Writef("}")
	}
	return false
}

func genPointerMarshaler(o *writer, name string, t *gobuf.Type) bool {
	if t.Kind == gobuf.POINTER {
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

func genScalarMarshaler(o *writer, name string, t *gobuf.Type) {
	switch t.Kind {
	case gobuf.INT8, gobuf.UINT8:
		o.Writef("b[n] = (byte)%s;", name)
		o.Writef("n += 1;")
	case gobuf.INT16, gobuf.UINT16:
		o.Writef("n = Gobuf.WriteUint16((ushort)%s, b, n);", name)
	case gobuf.INT32, gobuf.UINT32:
		o.Writef("n = Gobuf.WriteUint32((long)%s, b, n);", name)
	case gobuf.INT64, gobuf.UINT64:
		o.Writef("n = Gobuf.WriteUint64((ulong)%s, b, n);", name)
	case gobuf.FLOAT32:
		o.Writef("n = Gobuf.WriteFloat32((float)%s, b, n);", name)
	case gobuf.FLOAT64:
		o.Writef("n = Gobuf.WriteFloat64((double)%s, b, n);", name)
	case gobuf.INT:
		o.Writef("n = Gobuf.WriteVarint((long)%s, b, n);", name)
	case gobuf.UINT:
		o.Writef("n = Gobuf.WriteUvarint((ulong)%s, b, n);", name)
	case gobuf.BYTES:
		o.Writef("n = Gobuf.WriteBytes(%s, b, n);", name)
	case gobuf.STRING:
		o.Writef("n = Gobuf.WriteString(%s, b, n);", name)
	case gobuf.MESSAGE:
		o.Writef("n = %s.Marshal(b, n);", name)
	}
}
