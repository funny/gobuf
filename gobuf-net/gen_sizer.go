package main

import (
	"fmt"

	"github.com/funny/gobuf"
)

func genSizer(o *writer, name string, t *gobuf.Type, n int) {
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

func genArraySizer(o *writer, name string, t *gobuf.Type, n int) bool {
	if t.Kind == gobuf.ARRAY {
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

func genMapSizer(o *writer, name string, t *gobuf.Type, n int) bool {
	if t.Kind == gobuf.MAP {
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

func genPointerSizer(o *writer, name string, t *gobuf.Type) bool {
	if t.Kind == gobuf.POINTER {
		o.Writef("size += 1;")
		o.Writef("if (%s != null) {", name)
		genScalarSizer(o, name, t.Elem)
		o.Writef("}")
		return true
	}
	return false
}

func genScalarSizer(o *writer, name string, t *gobuf.Type) {
	if t.Size() != 0 {
		o.Writef("size += %d;", t.Size())
		return
	}
	switch t.Kind {
	case gobuf.INT:
		o.Writef("size += Gobuf.VarintSize((long)%s);", name)
	case gobuf.UINT:
		o.Writef("size += Gobuf.UvarintSize((ulong)%s);", name)
	case gobuf.STRING:
		o.Writef("size += Gobuf.StringSize(%s);", name, name)
	case gobuf.BYTES:
		o.Writef("size += Gobuf.UvarintSize((ulong)%s.Length) + %s.Length;", name, name)
	case gobuf.MESSAGE:
		o.Writef("size += %s.Size();", name)
	}
}
