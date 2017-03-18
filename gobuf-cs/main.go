package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/funny/gobuf/gb"

	"strings"
)

func main() {
	var doc gb.Doc

	decoder := json.NewDecoder(os.Stdin)

	if err := decoder.Decode(&doc); err != nil {
		log.Fatal(err)
	}

	var o writer

	for _, s := range doc.Structs {
		o.Writef("class %s {", s.Name)

		for _, field := range s.Fields {
			o.Writef("public %s %s;", typeName(field.Type), field.Name)
		}

		o.Writef("public int Size() {")
		o.Writef("int size;")
		for _, field := range s.Fields {
			genSizer(&o, "this."+field.Name, field.Type, 1)
		}
		o.Writef("return size;")
		o.Writef("}")

		o.Writef("public int Marshal(byte[] b, int n) {")
		for _, field := range s.Fields {
			genMarshaler(&o, "this."+field.Name, field.Type, 1)
		}
		o.Writef("return n;")
		o.Writef("}")

		o.Writef("public int Unmarshal(byte[] b, int n) {")
		for _, field := range s.Fields {
			genUnmarshaler(&o, "this."+field.Name, field.Type, 1)
		}
		o.Writef("return n;")
		o.Writef("}")

		o.Writef("}")
	}

	if _, err := o.WriteTo(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

type writer struct {
	deepth int
	bytes.Buffer
}

func (w *writer) Writef(format string, args ...interface{}) {
	format = strings.TrimLeft(format, "\t ")

	if format[0] == '}' {
		w.deepth--
	}

	for i := 0; i < w.deepth; i++ {
		w.WriteByte('\t')
	}

	if format[len(format)-1] == '{' {
		w.deepth++
	}

	w.WriteString(fmt.Sprintf(format, args...))
	w.WriteByte('\n')
}

func typeName(t *gb.Type) string {
	if t.Name != "" {
		return t.Name
	}
	switch t.Kind {
	case gb.INT:
		return "long"
	case gb.UINT:
		return "ulong"
	case gb.INT8:
		return "sbyte"
	case gb.UINT8:
		return "byte"
	case gb.INT16:
		return "short"
	case gb.UINT16:
		return "ushort"
	case gb.INT32:
		return "int"
	case gb.UINT32:
		return "uint"
	case gb.INT64:
		return "long"
	case gb.UINT64:
		return "ulong"
	case gb.FLOAT32:
		return "float"
	case gb.FLOAT64:
		return "double"
	case gb.STRING:
		return "string"
	case gb.BYTES:
		return "byte[]"
	case gb.BOOL:
		return "bool"
	case gb.MAP:
		return fmt.Sprintf("Dictionary<%s, %s>", typeName(t.Key), typeName(t.Elem))
	case gb.POINTER:
		if t.Elem.Kind == gb.STRUCT {
			return typeName(t.Elem)
		}
		return fmt.Sprintf("Nullable<%s>", typeName(t.Elem))
	case gb.ARRAY:
		if t.Len != 0 {
			return fmt.Sprintf("%s[%d]", t.Len, typeName(t.Elem))
		}
		return fmt.Sprintf("%s[]", typeName(t.Elem))
	}
	return ""
}
