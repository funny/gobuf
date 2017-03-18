package main

import (
	"fmt"

	"github.com/funny/gobuf/gb"
)

func genSizer(o *writer, name string, t *gb.Type, n int) {
	if genArraySizer(o, name, t, n) {
		return
	}
	if genMapSizer(o, name, t, n) {
		return
	}
	if genPointerSizer(o, name, t) {
		return
	}
	genScalarSizer(o, name, t)
}

func genArraySizer(o *writer, name string, t *gb.Type, n int) bool {
	if t.Kind == gb.ARRAY {
		elemSize := t.Elem.Size()
		if elemSize != 0 {
			o.Writef("size += Gobuf.UvarintSize((ulong)%s.Length) + %s.Length * %d;", name, name, elemSize)
		} else {
			o.Writef("size += Gobuf.UvarintSize((ulong)%s.Length);", name)
			o.Writef("for (var i%d = 0; i%d < %s.Length; i%d ++) {", n, n, name, n)
			genSizer(o, fmt.Sprintf("%s[i%d]", name, n), t.Elem, n+1)
			o.Writef("}")
		}
		return true
	}
	return false
}

func genMapSizer(o *writer, name string, t *gb.Type, n int) bool {
	if t.Kind == gb.MAP {
		keySize := t.Key.Size()
		elemSize := t.Elem.Size()
		if keySize != 0 && elemSize != 0 {
			o.Writef("size += Gobuf.UvarintSize((ulong)%s.Length) + %s.Length * (%d + %d);", name, name, keySize, elemSize)
		} else {
			o.Writef("size += Gobuf.UvarintSize((ulong)%s.Length);", name)
			o.Writef("foreach (var item%d in %s) {", n, name)
			genScalarSizer(o, fmt.Sprintf("item%d.Key", n), t.Key)
			genSizer(o, fmt.Sprintf("item%d.Value", n), t.Elem, n+1)
			o.Writef("}")
		}
	}
	return false
}

func genPointerSizer(o *writer, name string, t *gb.Type) bool {
	if t.Kind == gb.POINTER {
		o.Writef("size += 1;")
		o.Writef("if (%s != null) {", name)
		genScalarSizer(o, name, t.Elem)
		o.Writef("}")
		return true
	}
	return false
}

func genScalarSizer(o *writer, name string, t *gb.Type) {
	if t.Size() != 0 {
		o.Writef("size += %d;", t.Size())
		return
	}
	switch t.Kind {
	case gb.INT:
		o.Writef("size += Gobuf.VarintSize((long)%s);", name)
	case gb.UINT:
		o.Writef("size += Gobuf.UvarintSize((ulong)%s);", name)
	case gb.STRING:
		o.Writef("size += Gobuf.StringSize(%s);", name)
	case gb.BYTES:
		o.Writef("size += Gobuf.UvarintSize((ulong)%s.Length) + %s.Length;", name, name)
	case gb.MESSAGE:
		o.Writef("size += %s.Size();", name)
	}
}
