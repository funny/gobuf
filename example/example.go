package example

type Enum int

const (
	A Enum = iota
	B
	C
	D
)

type Scalar struct {
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

type Pointer struct {
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
	BoolPtr    *bool
}

type Array struct {
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
	BoolArray    []bool
}

type FixlenArray struct {
	IntArray     [1]int
	UintArray    [1]uint
	Int8Array    [1]int8
	Uint8Array   [1]uint8
	Int16Array   [1]int16
	Uint16Array  [1]uint16
	Int32Array   [1]int32
	Uint32Array  [1]uint32
	Int64Array   [1]int64
	Uint64Array  [1]uint64
	Float32Array [1]float32
	Float64Array [1]float64
	StringArray  [1]string
	BoolArray    [1]bool
}

type Map struct {
	IntMap     map[int]int
	UintMap    map[int]uint
	Int8Map    map[int]int8
	Uint8Map   map[int]uint8
	Int16Map   map[int]int16
	Uint16Map  map[int]uint16
	Int32Map   map[int]int32
	Uint32Map  map[int]uint32
	Int64Map   map[int]int64
	Uint64Map  map[int]uint64
	Float32Map map[int]float32
	Float64Map map[int]float64
	StringMap  map[int]string
	BoolMap    map[int]bool
}

type Message struct {
	Scalar      Scalar
	ScalarPtr   *Scalar
	ScalarArray []Scalar
	ScalarMap   map[int]*Scalar
}
