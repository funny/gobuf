package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/funny/gobuf/parser"

	"strings"
)

func main() {
	var doc parser.Doc

	decoder := json.NewDecoder(os.Stdin)

	if err := decoder.Decode(&doc); err != nil {
		log.Fatal(err)
	}

	var o writer

	o.Writef("using System;")
	o.Writef("using System.Collections.Generic;")

	for _, s := range doc.Structs {
		o.Writef("class %s {", s.Name)

		for _, field := range s.Fields {
			if field.Type.Kind == parser.ARRAY && field.Type.Len != 0 {
				o.Writef("public %s %s = new List<%s>(%d);",
					typeName(field.Type), field.Name, typeName(field.Type.Elem), field.Type.Len)
			} else {
				o.Writef("public %s %s;", typeName(field.Type), field.Name)
			}
		}

		o.Writef("public int Size() {")
		o.Writef("int size = 0;")
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

func isNullable(t *parser.Type) bool {
	return t.Kind == parser.POINTER && t.Elem.Kind != parser.STRUCT && t.Elem.Kind != parser.STRING
}

func typeName(t *parser.Type) string {
	if t.Name != "" {
		return t.Name
	}
	switch t.Kind {
	case parser.INT:
		return "long"
	case parser.UINT:
		return "ulong"
	case parser.INT8:
		return "sbyte"
	case parser.UINT8:
		return "byte"
	case parser.INT16:
		return "short"
	case parser.UINT16:
		return "ushort"
	case parser.INT32:
		return "int"
	case parser.UINT32:
		return "uint"
	case parser.INT64:
		return "long"
	case parser.UINT64:
		return "ulong"
	case parser.FLOAT32:
		return "float"
	case parser.FLOAT64:
		return "double"
	case parser.STRING:
		return "string"
	case parser.BYTES:
		return "byte[]"
	case parser.BOOL:
		return "bool"
	case parser.MAP:
		return fmt.Sprintf("Dictionary<%s, %s>", typeName(t.Key), typeName(t.Elem))
	case parser.POINTER:
		if t.Elem.Kind == parser.STRUCT {
			return typeName(t.Elem)
		}
		if t.Elem.Kind == parser.STRING {
			return "string"
		}
		return fmt.Sprintf("Nullable<%s>", typeName(t.Elem))
	case parser.ARRAY:
		return fmt.Sprintf("List<%s>", typeName(t.Elem))
	}
	return ""
}
