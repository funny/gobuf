package example

import (
	"math"
	"testing"

	"github.com/funny/utest"
)

func TestScalar(t *testing.T) {
	var msg1 = Scalar{
		Byte:    math.MaxUint8,
		Int:     math.MaxInt64,
		Uint:    math.MaxUint64,
		Int8:    math.MaxInt8,
		Uint8:   math.MaxUint8,
		Int16:   math.MaxInt16,
		Uint16:  math.MaxUint16,
		Int32:   math.MaxInt32,
		Uint32:  math.MaxUint32,
		Int64:   math.MaxInt64,
		Uint64:  math.MaxUint64,
		Float32: math.MaxFloat32,
		Float64: math.MaxFloat64,
		String:  "test string content",
		Bytes:   []byte("text bytes content"),
		Bool:    true,
	}

	var msg2 Scalar

	data := make([]byte, msg1.Size())
	size1 := msg1.Marshal(data)
	size2 := msg2.Unmarshal(data)

	utest.EqualNow(t, size1, msg1.Size())
	utest.EqualNow(t, size2, msg1.Size())
	utest.EqualNow(t, msg1.Byte, msg2.Byte)
	utest.EqualNow(t, msg1.Int, msg2.Int)
	utest.EqualNow(t, msg1.Uint, msg2.Uint)
	utest.EqualNow(t, msg1.Int8, msg2.Int8)
	utest.EqualNow(t, msg1.Uint8, msg2.Uint8)
	utest.EqualNow(t, msg1.Int16, msg2.Int16)
	utest.EqualNow(t, msg1.Uint16, msg2.Uint16)
	utest.EqualNow(t, msg1.Int32, msg2.Int32)
	utest.EqualNow(t, msg1.Uint32, msg2.Uint32)
	utest.EqualNow(t, msg1.Int64, msg2.Int64)
	utest.EqualNow(t, msg1.Uint64, msg2.Uint64)
	utest.EqualNow(t, msg1.Float32, msg2.Float32)
	utest.EqualNow(t, msg1.Float64, msg2.Float64)
	utest.EqualNow(t, msg1.String, msg2.String)
	utest.EqualNow(t, msg1.Bytes, msg2.Bytes)
	utest.EqualNow(t, msg1.Bool, msg2.Bool)
}

func TestPointer(t *testing.T) {
	var msg1 Pointer
	var msg2 Pointer

	data := make([]byte, msg1.Size())
	size1 := msg1.Marshal(data)
	size2 := msg2.Unmarshal(data)

	utest.EqualNow(t, size1, msg1.Size())
	utest.EqualNow(t, size2, msg1.Size())
	utest.EqualNow(t, msg1.IntPtr, msg2.IntPtr)
	utest.EqualNow(t, msg1.UintPtr, msg2.UintPtr)
	utest.EqualNow(t, msg1.Int8Ptr, msg2.Int8Ptr)
	utest.EqualNow(t, msg1.Uint8Ptr, msg2.Uint8Ptr)
	utest.EqualNow(t, msg1.Int16Ptr, msg2.Int16Ptr)
	utest.EqualNow(t, msg1.Uint16Ptr, msg2.Uint16Ptr)
	utest.EqualNow(t, msg1.Int32Ptr, msg2.Int32Ptr)
	utest.EqualNow(t, msg1.Uint32Ptr, msg2.Uint32Ptr)
	utest.EqualNow(t, msg1.Int64Ptr, msg2.Int64Ptr)
	utest.EqualNow(t, msg1.Uint64Ptr, msg2.Uint64Ptr)
	utest.EqualNow(t, msg1.Float32Ptr, msg2.Float32Ptr)
	utest.EqualNow(t, msg1.Float64Ptr, msg2.Float64Ptr)
	utest.EqualNow(t, msg1.StringPtr, msg2.StringPtr)
	utest.EqualNow(t, msg1.BoolPtr, msg2.BoolPtr)
}

func TestArray(t *testing.T) {
	var msg1 = Array{
		IntArray:     []int{math.MaxInt64},
		UintArray:    []uint{math.MaxUint64},
		Int8Array:    []int8{math.MaxInt8},
		Uint8Array:   []uint8{math.MaxUint8},
		Int16Array:   []int16{math.MaxInt16},
		Uint16Array:  []uint16{math.MaxUint16},
		Int32Array:   []int32{math.MaxInt32},
		Uint32Array:  []uint32{math.MaxUint32},
		Int64Array:   []int64{math.MaxInt64},
		Uint64Array:  []uint64{math.MaxUint64},
		Float32Array: []float32{math.MaxFloat32},
		Float64Array: []float64{math.MaxFloat64},
		StringArray:  []string{"test string content"},
		BoolArray:    []bool{true},
	}

	var msg2 Array

	data := make([]byte, msg1.Size())
	size1 := msg1.Marshal(data)
	size2 := msg2.Unmarshal(data)

	utest.EqualNow(t, size1, msg1.Size())
	utest.EqualNow(t, size2, msg1.Size())
	utest.EqualNow(t, msg1.IntArray, msg2.IntArray)
	utest.EqualNow(t, msg1.UintArray, msg2.UintArray)
	utest.EqualNow(t, msg1.Int8Array, msg2.Int8Array)
	utest.EqualNow(t, msg1.Uint8Array, msg2.Uint8Array)
	utest.EqualNow(t, msg1.Int16Array, msg2.Int16Array)
	utest.EqualNow(t, msg1.Uint16Array, msg2.Uint16Array)
	utest.EqualNow(t, msg1.Int32Array, msg2.Int32Array)
	utest.EqualNow(t, msg1.Uint32Array, msg2.Uint32Array)
	utest.EqualNow(t, msg1.Int64Array, msg2.Int64Array)
	utest.EqualNow(t, msg1.Uint64Array, msg2.Uint64Array)
	utest.EqualNow(t, msg1.Float32Array, msg2.Float32Array)
	utest.EqualNow(t, msg1.Float64Array, msg2.Float64Array)
	utest.EqualNow(t, msg1.StringArray, msg2.StringArray)
	utest.EqualNow(t, msg1.BoolArray, msg2.BoolArray)
}

func TestFixlenArray(t *testing.T) {
	var msg1 = FixlenArray{
		IntArray:     [1]int{math.MaxInt64},
		UintArray:    [1]uint{math.MaxUint64},
		Int8Array:    [1]int8{math.MaxInt8},
		Uint8Array:   [1]uint8{math.MaxUint8},
		Int16Array:   [1]int16{math.MaxInt16},
		Uint16Array:  [1]uint16{math.MaxUint16},
		Int32Array:   [1]int32{math.MaxInt32},
		Uint32Array:  [1]uint32{math.MaxUint32},
		Int64Array:   [1]int64{math.MaxInt64},
		Uint64Array:  [1]uint64{math.MaxUint64},
		Float32Array: [1]float32{math.MaxFloat32},
		Float64Array: [1]float64{math.MaxFloat64},
		StringArray:  [1]string{"test string content"},
		BoolArray:    [1]bool{true},
	}

	var msg2 FixlenArray

	data := make([]byte, msg1.Size())
	size1 := msg1.Marshal(data)
	size2 := msg2.Unmarshal(data)

	utest.EqualNow(t, size1, msg1.Size())
	utest.EqualNow(t, size2, msg1.Size())
	utest.EqualNow(t, msg1.IntArray, msg2.IntArray)
	utest.EqualNow(t, msg1.UintArray, msg2.UintArray)
	utest.EqualNow(t, msg1.Int8Array, msg2.Int8Array)
	utest.EqualNow(t, msg1.Uint8Array, msg2.Uint8Array)
	utest.EqualNow(t, msg1.Int16Array, msg2.Int16Array)
	utest.EqualNow(t, msg1.Uint16Array, msg2.Uint16Array)
	utest.EqualNow(t, msg1.Int32Array, msg2.Int32Array)
	utest.EqualNow(t, msg1.Uint32Array, msg2.Uint32Array)
	utest.EqualNow(t, msg1.Int64Array, msg2.Int64Array)
	utest.EqualNow(t, msg1.Uint64Array, msg2.Uint64Array)
	utest.EqualNow(t, msg1.Float32Array, msg2.Float32Array)
	utest.EqualNow(t, msg1.Float64Array, msg2.Float64Array)
	utest.EqualNow(t, msg1.StringArray, msg2.StringArray)
	utest.EqualNow(t, msg1.BoolArray, msg2.BoolArray)
}

func TestMap(t *testing.T) {
	var msg1 = Map{
		IntMap:     map[int]int{1: math.MaxInt64},
		UintMap:    map[int]uint{1: math.MaxUint64},
		Int8Map:    map[int]int8{1: math.MaxInt8},
		Uint8Map:   map[int]uint8{1: math.MaxUint8},
		Int16Map:   map[int]int16{1: math.MaxInt16},
		Uint16Map:  map[int]uint16{1: math.MaxUint16},
		Int32Map:   map[int]int32{1: math.MaxInt32},
		Uint32Map:  map[int]uint32{1: math.MaxUint32},
		Int64Map:   map[int]int64{1: math.MaxInt64},
		Uint64Map:  map[int]uint64{1: math.MaxUint64},
		Float32Map: map[int]float32{1: math.MaxFloat32},
		Float64Map: map[int]float64{1: math.MaxFloat64},
		StringMap:  map[int]string{1: "test string content"},
		BoolMap:    map[int]bool{1: true},
	}

	var msg2 Map

	data := make([]byte, msg1.Size())
	size1 := msg1.Marshal(data)
	size2 := msg2.Unmarshal(data)

	utest.EqualNow(t, size1, msg1.Size())
	utest.EqualNow(t, size2, msg1.Size())
	utest.EqualNow(t, msg1.IntMap, msg2.IntMap)
	utest.EqualNow(t, msg1.UintMap, msg2.UintMap)
	utest.EqualNow(t, msg1.Int8Map, msg2.Int8Map)
	utest.EqualNow(t, msg1.Uint8Map, msg2.Uint8Map)
	utest.EqualNow(t, msg1.Int16Map, msg2.Int16Map)
	utest.EqualNow(t, msg1.Uint16Map, msg2.Uint16Map)
	utest.EqualNow(t, msg1.Int32Map, msg2.Int32Map)
	utest.EqualNow(t, msg1.Uint32Map, msg2.Uint32Map)
	utest.EqualNow(t, msg1.Int64Map, msg2.Int64Map)
	utest.EqualNow(t, msg1.Uint64Map, msg2.Uint64Map)
	utest.EqualNow(t, msg1.Float32Map, msg2.Float32Map)
	utest.EqualNow(t, msg1.Float64Map, msg2.Float64Map)
	utest.EqualNow(t, msg1.StringMap, msg2.StringMap)
	utest.EqualNow(t, msg1.BoolMap, msg2.BoolMap)
}

func TestMessage(t *testing.T) {
	var scalar = Scalar{
		Byte:    math.MaxUint8,
		Int:     math.MaxInt64,
		Uint:    math.MaxUint64,
		Int8:    math.MaxInt8,
		Uint8:   math.MaxUint8,
		Int16:   math.MaxInt16,
		Uint16:  math.MaxUint16,
		Int32:   math.MaxInt32,
		Uint32:  math.MaxUint32,
		Int64:   math.MaxInt64,
		Uint64:  math.MaxUint64,
		Float32: math.MaxFloat32,
		Float64: math.MaxFloat64,
		String:  "test string content",
		Bytes:   []byte("text bytes content"),
		Bool:    true,
	}

	var msg1 = Message{
		Scalar:      scalar,
		ScalarPtr:   &scalar,
		ScalarArray: []Scalar{scalar},
		ScalarMap:   map[int]*Scalar{1: &scalar},
	}

	var msg2 Message

	data := make([]byte, msg1.Size())
	size1 := msg1.Marshal(data)
	size2 := msg2.Unmarshal(data)

	utest.EqualNow(t, size1, msg1.Size())
	utest.EqualNow(t, size2, msg1.Size())

	check := func(msg1, msg2 *Scalar) {
		utest.EqualNow(t, msg1.Byte, msg2.Byte)
		utest.EqualNow(t, msg1.Int, msg2.Int)
		utest.EqualNow(t, msg1.Uint, msg2.Uint)
		utest.EqualNow(t, msg1.Int8, msg2.Int8)
		utest.EqualNow(t, msg1.Uint8, msg2.Uint8)
		utest.EqualNow(t, msg1.Int16, msg2.Int16)
		utest.EqualNow(t, msg1.Uint16, msg2.Uint16)
		utest.EqualNow(t, msg1.Int32, msg2.Int32)
		utest.EqualNow(t, msg1.Uint32, msg2.Uint32)
		utest.EqualNow(t, msg1.Int64, msg2.Int64)
		utest.EqualNow(t, msg1.Uint64, msg2.Uint64)
		utest.EqualNow(t, msg1.Float32, msg2.Float32)
		utest.EqualNow(t, msg1.Float64, msg2.Float64)
		utest.EqualNow(t, msg1.String, msg2.String)
		utest.EqualNow(t, msg1.Bytes, msg2.Bytes)
		utest.EqualNow(t, msg1.Bool, msg2.Bool)
	}

	check(&msg1.Scalar, &msg2.Scalar)
	check(msg1.ScalarPtr, msg2.ScalarPtr)
	check(&msg1.ScalarArray[0], &msg2.ScalarArray[0])
	check(msg1.ScalarMap[1], msg2.ScalarMap[1])
}
