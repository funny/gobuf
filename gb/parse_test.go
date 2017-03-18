package gb

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"encoding/json"

	"github.com/funny/utest"
)

func Test_ParseStructs(t *testing.T) {
	src := `
package main

type Enum1 int

const (
	A Enum1 = iota
	B
	C
	D
)

type Enum2 Enum1

type ScalarTypes struct {
	Byte    byte
	Int     int
	Uint    uint
	Int8    int8
	Uint8   uint8
	Int16   int16
	Uint16  uint16
	Int32   int32
	Uint32  uint32
	Int64   int64
	Uint64  uint64
	Float32 float32
	Float64 float64
}

type CompositeTypes struct {
	String  string
	Bytes   []byte
	
	Message    ScalarTypes
	MessagePtr *ScalarTypes

	IntPtr     *int
	UintPtr    *uint
	Int8Ptr    *int8
	Uint8Ptr   *uint8
	Int16Ptr   *int16
	Uint16Ptr  *uint16
	Int32Ptr   *int32
	Uint32Ptr  *uint32
	Int64Ptr   *int64
	Uint64Ptr  *uint64
	Float32Ptr *float32
	Float64Ptr *float64
	StringPtr  *string

	IntArray     []int
	UintArray    []uint
	Int8Array    []int8
	Uint8Array   []uint8
	Int16Array   []int16
	Uint16Array  []uint16
	Int32Array   []int32
	Uint32Array  []uint32
	Int64Array   []int64
	Uint64Array  []uint64
	Float32Array []float32
	Float64Array []float64
	StringArray  []string
}
`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("", fset, []*ast.File{f}, nil)
	if err != nil {
		t.Fatal(err)
	}

	structs := scanMessages(f, pkg)

	utest.EqualNow(t, 2, len(structs))
	utest.Assert(t, structs["ScalarTypes"] != nil)
	utest.Assert(t, structs["CompositeTypes"] != nil)

	for name, s := range structs {
		t.Log(name)
		for i := 0; i < s.NumFields(); i++ {
			data, err := json.Marshal(analyzeType(s.Field(i).Type()))
			utest.IsNilNow(t, err)
			t.Logf("	%s %s %s", s.Field(i).Name(), s.Field(i).Type(), data)
		}
	}

	consts := scanConsts(f, pkg)

	for name, c := range consts {
		t.Log(name, c.Type().String(), c.Type().Underlying().String(), c.Val().String())
	}
}
