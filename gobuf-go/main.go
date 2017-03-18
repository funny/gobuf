package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/funny/gobuf/gb"

	"go/format"
)

func main() {
	var doc gb.Doc

	decoder := json.NewDecoder(os.Stdin)

	if err := decoder.Decode(&doc); err != nil {
		log.Fatal(err)
	}

	var o writer

	o.Writef("package %s", doc.Package)

	o.Writef(`import "github.com/funny/gobuf"`)
	o.Writef(`import "encoding/binary"`)

	for _, msg := range doc.Messages {
		o.Writef("func (msg *%s) Size() int {", msg.Name)
		o.Writef("var size int")
		for _, field := range msg.Fields {
			genSizer(&o, "msg."+field.Name, field.Type, 1)
		}
		o.Writef("return size")
		o.Writef("}\n")

		o.Writef("func (msg *%s) Marshal(b []byte) int {", msg.Name)
		o.Writef("var n int")
		for _, field := range msg.Fields {
			genMarshaler(&o, "msg."+field.Name, field.Type, 1)
		}
		o.Writef("return n")
		o.Writef("}\n")

		o.Writef("func (msg *%s) Unmarshal(b []byte) int {", msg.Name)
		o.Writef("var n int")
		for _, field := range msg.Fields {
			genUnmarshaler(&o, "msg."+field.Name, field.Type, 1)
		}
		o.Writef("return n")
		o.Writef("}\n")
	}

	source, err := format.Source(o.Bytes())
	if err != nil {
		if _, err := o.WriteTo(os.Stdout); err != nil {
			log.Fatal(err)
		}
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

func typeName(t *gb.Type) string {
	if t.Name != "" {
		return t.Name
	}
	switch t.Kind {
	case gb.INT:
		return "int"
	case gb.UINT:
		return "uint"
	case gb.INT8:
		return "int8"
	case gb.UINT8:
		return "uint8"
	case gb.INT16:
		return "int16"
	case gb.UINT16:
		return "uint16"
	case gb.INT32:
		return "int32"
	case gb.UINT32:
		return "uint32"
	case gb.INT64:
		return "int64"
	case gb.UINT64:
		return "uint64"
	case gb.FLOAT32:
		return "float32"
	case gb.FLOAT64:
		return "float64"
	case gb.STRING:
		return "string"
	case gb.BYTES:
		return "[]byte"
	case gb.BOOL:
		return "bool"
	case gb.MAP:
		return fmt.Sprintf("map[%s]%s", typeName(t.Key), typeName(t.Elem))
	case gb.POINTER:
		return fmt.Sprintf("*%s", typeName(t.Elem))
	case gb.ARRAY:
		if t.Len != 0 {
			return fmt.Sprintf("[%d]%s", t.Len, typeName(t.Elem))
		}
		return fmt.Sprintf("[]%s", typeName(t.Elem))
	}
	return ""
}
