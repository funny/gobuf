package parser

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
	STRUCT  = "Struct"
)

type Doc struct {
	File    string
	Package string
	Enums   []*Enum
	Structs []*Struct
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

type Struct struct {
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
