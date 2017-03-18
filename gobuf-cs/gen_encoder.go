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
		o.Writef("Gobuf.WriteUvarint((ulong)%s.Count, b, ref n);", name)
		o.Writef("for (var i%d = 0; i%d < %s.Count; i%d ++) {", n, n, name, n)
		genMarshaler(o, fmt.Sprintf("%s[i%d]", name, n), t.Elem, n+1)
		o.Writef("}")
		return true
	}
	return false
}

func genMapMarshaler(o *writer, name string, t *parser.Type, n int) bool {
	if t.Kind == parser.MAP {
		o.Writef("Gobuf.WriteUvarint((ulong)%s.Count, b, ref n);", name)
		o.Writef("foreach (var item%d in %s) {", n, name)
		genScalarMarshaler(o, fmt.Sprintf("item%d.Key", n), t.Key)
		genMarshaler(o, fmt.Sprintf("item%d.Value", n), t.Elem, n+1)
		o.Writef("}")
	}
	return false
}

func genPointerMarshaler(o *writer, name string, t *parser.Type) bool {
	if t.Kind == parser.POINTER {
		valName := name
		if isNullable(t) {
			valName += ".Value"
		}
		o.Writef("if (%s != null) {", name)
		o.Writef("	b[n++] = 1;")
		genScalarMarshaler(o, valName, t.Elem)
		o.Writef("} else {")
		o.Writef("	b[n++] = 0;")
		o.Writef("}")
		return true
	}
	return false
}

func genScalarMarshaler(o *writer, name string, t *parser.Type) {
	switch t.Kind {
	case parser.BOOL:
		o.Writef("b[n++] = %s ? (byte)1 : (byte)0;", name)
	case parser.INT8, parser.UINT8:
		o.Writef("b[n++] = (byte)%s;", name)
	case parser.INT16, parser.UINT16:
		o.Writef("Gobuf.WriteUint16((ushort)%s, b, ref n);", name)
	case parser.INT32, parser.UINT32:
		o.Writef("Gobuf.WriteUint32((uint)%s, b, ref n);", name)
	case parser.INT64, parser.UINT64:
		o.Writef("Gobuf.WriteUint64((ulong)%s, b, ref n);", name)
	case parser.INT:
		o.Writef("Gobuf.WriteVarint(%s, b, ref n);", name)
	case parser.UINT:
		o.Writef("Gobuf.WriteUvarint(%s, b, ref n);", name)
	case parser.FLOAT32:
		o.Writef("Gobuf.WriteFloat32(%s, b, ref n);", name)
	case parser.FLOAT64:
		o.Writef("Gobuf.WriteFloat64(%s, b, ref n);", name)
	case parser.BYTES:
		o.Writef("Gobuf.WriteBytes(%s, b, ref n);", name)
	case parser.STRING:
		o.Writef("Gobuf.WriteString(%s, b, ref n);", name)
	case parser.STRUCT:
		o.Writef("n = %s.Marshal(b, n);", name)
	}
}
