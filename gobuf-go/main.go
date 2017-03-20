package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/funny/gobuf/parser"

	"go/format"
)

func main() {
	var doc parser.Doc

	decoder := json.NewDecoder(os.Stdin)

	if err := decoder.Decode(&doc); err != nil {
		log.Fatal(err)
	}

	var o writer

	o.Writef("package %s", doc.Package)

	o.Writef(`import "encoding/binary"`)
	o.Writef(`import "github.com/funny/gobuf"`)

	for _, s := range doc.Structs {
		o.Writef("var _ gobuf.Struct = (*%s)(nil)", s.Name)

		o.Writef("func (s *%s) Size() int {", s.Name)
		o.Writef("var size int")
		for _, field := range s.Fields {
			genSizer(&o, "s."+field.Name, field.Type, 1)
		}
		o.Writef("return size")
		o.Writef("}\n")

		o.Writef("func (s *%s) Marshal(b []byte) int {", s.Name)
		o.Writef("var n int")
		for _, field := range s.Fields {
			genMarshaler(&o, "s."+field.Name, field.Type, 1)
		}
		o.Writef("return n")
		o.Writef("}\n")

		o.Writef("func (s *%s) Unmarshal(b []byte) int {", s.Name)
		o.Writef("var n int")
		for _, field := range s.Fields {
			genUnmarshaler(&o, "s."+field.Name, field.Type, 1)
		}
		o.Writef("return n")
		o.Writef("}\n")
	}

	code := o.Bytes()

	source, err := format.Source(code)
	if err != nil {
		fmt.Print(string(code))
		log.Fatal(err)
	}

	if _, err := bytes.NewReader(source).WriteTo(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

type writer struct {
	bytes.Buffer
}

func (w *writer) Writef(format string, args ...interface{}) {
	w.WriteString(fmt.Sprintf(format, args...))
	w.WriteByte('\n')
}

func typeName(t *parser.Type) string {
	if t.Name != "" {
		return t.Name
	}
	switch t.Kind {
	case parser.INT:
		return "int"
	case parser.UINT:
		return "uint"
	case parser.INT8:
		return "int8"
	case parser.UINT8:
		return "uint8"
	case parser.INT16:
		return "int16"
	case parser.UINT16:
		return "uint16"
	case parser.INT32:
		return "int32"
	case parser.UINT32:
		return "uint32"
	case parser.INT64:
		return "int64"
	case parser.UINT64:
		return "uint64"
	case parser.FLOAT32:
		return "float32"
	case parser.FLOAT64:
		return "float64"
	case parser.STRING:
		return "string"
	case parser.BYTES:
		return "[]byte"
	case parser.BOOL:
		return "bool"
	case parser.MAP:
		return fmt.Sprintf("map[%s]%s", typeName(t.Key), typeName(t.Elem))
	case parser.POINTER:
		return fmt.Sprintf("*%s", typeName(t.Elem))
	case parser.ARRAY:
		if t.Len != 0 {
			return fmt.Sprintf("[%d]%s", t.Len, typeName(t.Elem))
		}
		return fmt.Sprintf("[]%s", typeName(t.Elem))
	}
	return ""
}
