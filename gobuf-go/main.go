package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"fmt"

	"go/format"

	"github.com/hotgo/gobuf"
)

func main() {
	var doc gobuf.Doc

	decoder := json.NewDecoder(os.Stdin)

	if err := decoder.Decode(&doc); err != nil {
		log.Fatal(err)
	}

	var o writer

	o.Writef("package %s", doc.Package)

	o.Writef(`import "github.com/hotgo/gobuf"`)
	o.Writef(`import "encoding/binary"`)

	for _, msg := range doc.Messages {
		o.Writef("func (msg *%s) Size() int {", msg.Name)
		o.Writef("var size int")
		for _, field := range msg.Fields {
			genSizer(&o, "msg."+field.Name, field.Type)
		}
		o.Writef("return size")
		o.Writef("}\n")

		o.Writef("func (msg *%s) Marshal(b []byte) int {", msg.Name)
		o.Writef("var n int")
		for _, field := range msg.Fields {
			genMarshaler(&o, "msg."+field.Name, field.Type)
		}
		o.Writef("return n")
		o.Writef("}\n")

		o.Writef("func (msg *%s) Unmarshal(b []byte) int {", msg.Name)
		o.Writef("var n int")
		for _, field := range msg.Fields {
			genUnmarshaler(&o, "msg."+field.Name, field.Type)
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

func typeName(t *gobuf.Type) string {
	if t.Name != "" {
		return t.Name
	}
	switch t.Kind {
	case gobuf.INT:
		return "int"
	case gobuf.UINT:
		return "uint"
	case gobuf.INT8:
		return "int8"
	case gobuf.UINT8:
		return "uint8"
	case gobuf.INT16:
		return "int16"
	case gobuf.UINT16:
		return "uint16"
	case gobuf.INT32:
		return "int32"
	case gobuf.UINT32:
		return "uint32"
	case gobuf.INT64:
		return "int64"
	case gobuf.UINT64:
		return "uint64"
	case gobuf.FLOAT32:
		return "float32"
	case gobuf.FLOAT64:
		return "float64"
	case gobuf.STRING:
		return "string"
	}
	return ""
}

func genSizer(o *writer, name string, t *gobuf.Type) {
	switch {
	case t.IsArray:
		if t.Size() != 0 {
			o.Writef("size += %d * len(%s)", t.Size(), name)
		} else {
			o.Writef("for i := 0; i < len(%s); i ++ {", name)
			genScalarSizer(o, name+"[i]", t)
			o.Writef("}")
		}
	case t.IsPointer:
		o.Writef("size += 1")
		o.Writef("if %s != nil {", name)
		genScalarSizer(o, "*"+name, t)
		o.Writef("}")
	default:
		genScalarSizer(o, name, t)
	}
}

func genScalarSizer(o *writer, name string, t *gobuf.Type) {
	if t.Size() != 0 {
		o.Writef("size += %d", t.Size())
		return
	}
	switch t.Kind {
	case gobuf.INT:
		o.Writef("size += gobuf.VarintSize(int64(%s))", name)
	case gobuf.UINT:
		o.Writef("size += gobuf.UvarintSize(uint64(%s))", name)
	case gobuf.BYTES, gobuf.STRING:
		o.Writef("size += gobuf.UvarintSize(uint64(len(%s))) + len(%s)", name, name)
	case gobuf.MESSAGE:
		if name[0] == '*' {
			name = name[1:]
		}
		o.Writef("size += %s.Size()", name)
	}
}

func genMarshaler(o *writer, name string, t *gobuf.Type) {
	switch {
	case t.IsArray:
		o.Writef("n += binary.PutUvarint(b[n:], uint64(len(%s)))", name)
		o.Writef("for i := 0; i < len(%s); i ++ {", name)
		genScalarMarshaler(o, name+"[i]", t)
		o.Writef("}")
	case t.IsPointer:
		o.Writef("if %s != nil {", name)
		o.Writef("	b[n] = 1")
		o.Writef("	n ++")
		genScalarMarshaler(o, "*"+name, t)
		o.Writef("} else {")
		o.Writef("	b[n] = 0")
		o.Writef("	n ++")
		o.Writef("}")
	default:
		genScalarMarshaler(o, name, t)
	}
}

func genScalarMarshaler(o *writer, name string, t *gobuf.Type) {
	switch t.Kind {
	case gobuf.INT8, gobuf.UINT8:
		o.Writef("b[n] = byte(%s)", name)
		o.Writef("n += 1")
	case gobuf.INT16, gobuf.UINT16:
		o.Writef("binary.LittleEndian.PutUint16(b[n:], uint16(%s))", name)
		o.Writef("n += 2")
	case gobuf.INT32, gobuf.UINT32:
		o.Writef("binary.LittleEndian.PutUint32(b[n:], uint32(%s))", name)
		o.Writef("n += 4")
	case gobuf.INT64, gobuf.UINT64:
		o.Writef("binary.LittleEndian.PutUint64(b[n:], uint64(%s))", name)
		o.Writef("n += 8")
	case gobuf.FLOAT32:
		o.Writef("gobuf.PutFloat32(b[n:], float32(%s))", name)
		o.Writef("n += 4")
	case gobuf.FLOAT64:
		o.Writef("gobuf.PutFloat64(b[n:], float64(%s))", name)
		o.Writef("n += 8")
	case gobuf.INT:
		o.Writef("n += binary.PutVarint(b[n:], int64(%s))", name)
	case gobuf.UINT:
		o.Writef("n += binary.PutUvarint(b[n:], uint64(%s))", name)
	case gobuf.BYTES, gobuf.STRING:
		o.Writef("n += binary.PutUvarint(b[n:], uint64(len(%s)))", name)
		o.Writef("copy(b[n:], %s)", name)
		o.Writef("n += len(%s)", name)
	case gobuf.MESSAGE:
		if name[0] == '*' {
			name = name[1:]
		}
		o.Writef("n += %s.Marshal(b[n:])", name)
	}
}

func genUnmarshaler(o *writer, name string, t *gobuf.Type) {
	switch {
	case t.IsArray:
		o.Writef("{")
		o.Writef("	l, x := binary.Uvarint(b[n:])")
		o.Writef("	n += x")
		o.Writef("	%s = make([]%s, l)", name, typeName(t))
		o.Writef("	for i := 0; i < int(l); i ++ {")
		genScalarUnmarshaler(o, name+"[i]", t)
		o.Writef("	}")
		o.Writef("}")
	case t.IsPointer:
		o.Writef("if b[n] != 0 {")
		o.Writef("	n += 1")
		o.Writef("	%s = new(%s)", name, typeName(t))
		genScalarUnmarshaler(o, "*"+name, t)
		o.Writef("} else {")
		o.Writef("	n += 1")
		o.Writef("}")
	default:
		genScalarUnmarshaler(o, name, t)
	}
}

func genScalarUnmarshaler(o *writer, name string, t *gobuf.Type) {
	switch t.Kind {
	case gobuf.INT8, gobuf.UINT8:
		o.Writef("%s = %s(b[n])", name, typeName(t))
		o.Writef("n += 1")
	case gobuf.INT16, gobuf.UINT16:
		o.Writef("%s = %s(binary.LittleEndian.Uint16(b[n:]))", name, typeName(t))
		o.Writef("n += 2")
	case gobuf.INT32, gobuf.UINT32:
		o.Writef("%s = %s(binary.LittleEndian.Uint32(b[n:]))", name, typeName(t))
		o.Writef("n += 4")
	case gobuf.INT64, gobuf.UINT64:
		o.Writef("%s = %s(binary.LittleEndian.Uint64(b[n:]))", name, typeName(t))
		o.Writef("n += 8")
	case gobuf.FLOAT32:
		o.Writef("%s = %s(gobuf.GetFloat32(b[n:]))", name, typeName(t))
		o.Writef("n += 4")
	case gobuf.FLOAT64:
		o.Writef("%s = %s(gobuf.GetFloat64(b[n:]))", name, typeName(t))
		o.Writef("n += 8")
	case gobuf.INT:
		o.Writef("{")
		o.Writef("	v, x := binary.Varint(b[n:])")
		o.Writef("	%s = %s(v)", name, typeName(t))
		o.Writef("	n += x")
		o.Writef("}")
	case gobuf.UINT:
		o.Writef("{")
		o.Writef("	v, x := binary.Uvarint(b[n:])")
		o.Writef("	%s = %s(v)", name, typeName(t))
		o.Writef("	n += x")
		o.Writef("}")
	case gobuf.BYTES:
		o.Writef("{")
		o.Writef("	l, x := binary.Uvarint(b[n:])")
		o.Writef("	n += x")
		o.Writef("	%s = make([]byte, l)", name)
		o.Writef("	copy(%s, b[n:n+int(l)])", name)
		o.Writef("	n += int(l)")
		o.Writef("}")
	case gobuf.STRING:
		o.Writef("{")
		o.Writef("	l, x := binary.Uvarint(b[n:])")
		o.Writef("	n += x")
		o.Writef("	%s = string(b[n:n+int(l)])", name)
		o.Writef("	n += int(l)")
		o.Writef("}")
	case gobuf.MESSAGE:
		if name[0] == '*' {
			name = name[1:]
		}
		o.Writef("n += %s.Unmarshal(b[n:])", name)
	}
}
