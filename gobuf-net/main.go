package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"strings"

	"github.com/hotgo/gobuf"
)

func main() {
	var doc gobuf.Doc

	decoder := json.NewDecoder(os.Stdin)

	if err := decoder.Decode(&doc); err != nil {
		log.Fatal(err)
	}

	var o writer

	for _, msg := range doc.Messages {
		o.Writef("class %s {", msg.Name)

		for _, field := range msg.Fields {
			o.Writef("public %s %s;", typeName(field.Type), field.Name)
		}

		o.Writef("public int Size() {")
		o.Writef("int size;")
		for _, field := range msg.Fields {
			genSizer(&o, "this."+field.Name, field.Type, 1)
		}
		o.Writef("return size;")
		o.Writef("}")

		o.Writef("public int Marshal(byte[] b, int n) {")
		for _, field := range msg.Fields {
			genMarshaler(&o, "msg."+field.Name, field.Type, 1)
		}
		o.Writef("return n;")
		o.Writef("}")

		o.Writef("public int Unmarshal(byte[] b, int n) {")
		for _, field := range msg.Fields {
			genUnmarshaler(&o, "msg."+field.Name, field.Type, 1)
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

func typeName(t *gobuf.Type) string {
	if t.Name != "" {
		return t.Name
	}
	switch t.Kind {
	case gobuf.INT:
		return "long"
	case gobuf.UINT:
		return "ulong"
	case gobuf.INT8:
		return "sbyte"
	case gobuf.UINT8:
		return "byte"
	case gobuf.INT16:
		return "short"
	case gobuf.UINT16:
		return "ushort"
	case gobuf.INT32:
		return "int"
	case gobuf.UINT32:
		return "uint"
	case gobuf.INT64:
		return "long"
	case gobuf.UINT64:
		return "ulong"
	case gobuf.FLOAT32:
		return "float"
	case gobuf.FLOAT64:
		return "double"
	case gobuf.STRING:
		return "string"
	case gobuf.BYTES:
		return "byte[]"
	case gobuf.BOOL:
		return "bool"
	case gobuf.MAP:
		return fmt.Sprintf("Dictionary<%s, %s>", typeName(t.Key), typeName(t.Elem))
	case gobuf.POINTER:
		if t.Elem.Kind == gobuf.MESSAGE {
			return typeName(t.Elem)
		}
		return fmt.Sprintf("Nullable<%s>", typeName(t.Elem))
	case gobuf.ARRAY:
		if t.Len != 0 {
			return fmt.Sprintf("%s[%d]", t.Len, typeName(t.Elem))
		}
		return fmt.Sprintf("%s[]", typeName(t.Elem))
	}
	return ""
}
