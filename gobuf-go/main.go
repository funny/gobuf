package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/funny/gobuf/parser"

	"go/format"
	"path"
	"strings"
)

func main() {
	var doc parser.Doc

	decoder := json.NewDecoder(os.Stdin)

	if err := decoder.Decode(&doc); err != nil {
		log.Fatal(err)
	}

	var name = strings.TrimSuffix(path.Base(doc.File), path.Ext(doc.File))

	var o writer

	o.Writef("package %s", doc.Package)

	o.Writef(`import "math"`)
	o.Writef(`import "encoding/binary"`)

	for _, s := range doc.Structs {
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

	o.WriteString(`
func $name$_UvarintSize(x uint64) int {
	i := 0
	for x >= 0x80 {
		x >>= 7
		i++
	}
	return i + 1
}

func $name$_VarintSize(x int64) int {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return $name$_UvarintSize(ux)
}

func $name$_GetFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(b))
}

func $name$_PutFloat32(b []byte, v float32) {
	binary.LittleEndian.PutUint32(b, math.Float32bits(v))
}

func $name$_GetFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(b))
}

func $name$_PutFloat64(b []byte, v float64) {
	binary.LittleEndian.PutUint64(b, math.Float64bits(v))
}
`)

	code := bytes.Replace(o.Bytes(), []byte("$name$"), []byte(name), -1)

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
