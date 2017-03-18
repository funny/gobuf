package gb

import (
	"encoding/binary"
	"math"
)

const (
	INT     = "Int"
	UINT    = "Uint"
	INT8    = "Int8"
	UINT8   = "Uint8"
	INT16   = "Int16"
	UINT16  = "Uint16"
	INT32   = "Int32"
	UINT32  = "Uint32"
	INT64   = "Int64"
	UINT64  = "Uint64"
	FLOAT32 = "Float32"
	FLOAT64 = "Float64"
	BOOL    = "Bool"
	MAP     = "Map"
	ARRAY   = "Array"
	BYTES   = "Bytes"
	STRING  = "String"
	POINTER = "Pointer"
	MESSAGE = "Message"
)

type Doc struct {
	Package  string
	Enums    []*Enum
	Messages []*Message
}

type Enum struct {
	Name   string
	Kind   string
	Values []*Value
}

type Value struct {
	Name  string
	Value string
}

type Message struct {
	Name   string
	Fields []*Field
}

type Field struct {
	Name string
	Type *Type
}

type Type struct {
	Kind string `json:",omitempty"`
	Name string `json:",omitempty"`
	Key  *Type  `json:",omitempty"`
	Elem *Type  `json:",omitempty"`
	Len  int    `json:",omitempty"`
}

func (t *Type) Size() int {
	switch t.Kind {
	case INT8, UINT8, BOOL:
		return 1
	case INT16, UINT16:
		return 2
	case INT32, UINT32, FLOAT32:
		return 4
	case INT64, UINT64, FLOAT64:
		return 8
	}
	return 0
}

func Parse(filename string) (*Doc, error) {
	file, err := parseFile(filename)
	if err != nil {
		return nil, err
	}
	return analyzeFile(file)
}

func UvarintSize(x uint64) int {
	i := 0
	for x >= 0x80 {
		x >>= 7
		i++
	}
	return i + 1
}

func VarintSize(x int64) int {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return UvarintSize(ux)
}

func GetFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(b))
}

func PutFloat32(b []byte, v float32) {
	binary.LittleEndian.PutUint32(b, math.Float32bits(v))
}

func GetFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(b))
}

func PutFloat64(b []byte, v float64) {
	binary.LittleEndian.PutUint64(b, math.Float64bits(v))
}
