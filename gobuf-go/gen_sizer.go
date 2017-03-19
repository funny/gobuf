package main

import (
	"fmt"

	"github.com/funny/gobuf/parser"
)

func genSizer(o *writer, name string, t *parser.Type, n int) {
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

func genArraySizer(o *writer, name string, t *parser.Type, n int) bool {
	if t.Kind == parser.ARRAY {
		elemSize := t.Elem.Size()
		if elemSize != 0 {
			if t.Len == 0 {
				o.Writef("size += $name$_UvarintSize(uint64(len(%s))) + len(%s) * %d", name, name, elemSize)
			} else {
				o.Writef("size += %d * %d", t.Len, elemSize)
			}
		} else {
			if t.Len == 0 {
				o.Writef("size += $name$_UvarintSize(uint64(len(%s)))", name)
			}
			o.Writef("for i%d := 0; i%d < len(%s); i%d ++ {", n, n, name, n)
			genSizer(o, fmt.Sprintf("%s[i%d]", name, n), t.Elem, n+1)
			o.Writef("}")
		}
		return true
	}
	return false
}

func genMapSizer(o *writer, name string, t *parser.Type, n int) bool {
	if t.Kind == parser.MAP {
		keySize := t.Key.Size()
		elemSize := t.Elem.Size()
		if keySize != 0 && elemSize != 0 {
			o.Writef("size += $name$_UvarintSize(uint64(len(%s))) + len(%s) * (%d + %d)", name, name, keySize, elemSize)
		} else {
			o.Writef("size += $name$_UvarintSize(uint64(len(%s)))", name)
			if t.Key.Size() == 0 && t.Elem.Size() == 0 {
				o.Writef("for key%d, val%d := range %s {", n, n, name)
			} else if t.Key.Size() != 0 {
				o.Writef("for _, val%d := range %s {", n, name)
			} else if t.Elem.Size() != 0 {
				o.Writef("for key%d, _ := range %s {", n, name)
			}
			genScalarSizer(o, fmt.Sprintf("key%d", n), t.Key)
			genSizer(o, fmt.Sprintf("val%d", n), t.Elem, n+1)
			o.Writef("}")
		}
	}
	return false
}

func genPointerSizer(o *writer, name string, t *parser.Type) bool {
	if t.Kind == parser.POINTER {
		o.Writef("size += 1")
		o.Writef("if %s != nil {", name)
		genScalarSizer(o, "*"+name, t.Elem)
		o.Writef("}")
		return true
	}
	return false
}

func genScalarSizer(o *writer, name string, t *parser.Type) {
	if t.Size() != 0 {
		o.Writef("size += %d", t.Size())
		return
	}
	switch t.Kind {
	case parser.INT:
		o.Writef("size += $name$_VarintSize(int64(%s))", name)
	case parser.UINT:
		o.Writef("size += $name$_UvarintSize(uint64(%s))", name)
	case parser.BYTES:
		if t.Len != 0 {
			o.Writef("size += %d", t.Len)
			return
		}
		fallthrough
	case parser.STRING:
		o.Writef("size += $name$_UvarintSize(uint64(len(%s))) + len(%s)", name, name)
	case parser.STRUCT:
		if name[0] == '*' {
			name = name[1:]
		}
		o.Writef("size += %s.Size()", name)
	}
}
