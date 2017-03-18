package example

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
	String  string
	Bytes   []byte
	Bool    bool
}

type CompositeTypes struct {
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

	Message           ScalarTypes
	MessagePtr        *ScalarTypes
	MessageArray      []ScalarTypes
	MessagePtrArray   []*ScalarTypes
	MessageArrayArray [][]ScalarTypes

	IntMap map[int]map[int][]*ScalarTypes
}
