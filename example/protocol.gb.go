package example

import "math"
import "encoding/binary"

func (s *FixlenArray) Size() int {
	var size int
	for i1 := 0; i1 < len(s.IntArray); i1++ {
		size += protocol_VarintSize(int64(s.IntArray[i1]))
	}
	for i1 := 0; i1 < len(s.UintArray); i1++ {
		size += protocol_UvarintSize(uint64(s.UintArray[i1]))
	}
	size += 1 * 1
	size += 1
	size += 1 * 2
	size += 1 * 2
	size += 1 * 4
	size += 1 * 4
	size += 1 * 8
	size += 1 * 8
	size += 1 * 4
	size += 1 * 8
	for i1 := 0; i1 < len(s.StringArray); i1++ {
		size += protocol_UvarintSize(uint64(len(s.StringArray[i1]))) + len(s.StringArray[i1])
	}
	size += 1 * 1
	return size
}

func (s *FixlenArray) Marshal(b []byte) int {
	var n int
	for i1 := 0; i1 < len(s.IntArray); i1++ {
		n += binary.PutVarint(b[n:], int64(s.IntArray[i1]))
	}
	for i1 := 0; i1 < len(s.UintArray); i1++ {
		n += binary.PutUvarint(b[n:], uint64(s.UintArray[i1]))
	}
	for i1 := 0; i1 < len(s.Int8Array); i1++ {
		b[n] = byte(s.Int8Array[i1])
		n += 1
	}
	copy(b[n:], s.Uint8Array[:])
	n += 1
	for i1 := 0; i1 < len(s.Int16Array); i1++ {
		binary.LittleEndian.PutUint16(b[n:], uint16(s.Int16Array[i1]))
		n += 2
	}
	for i1 := 0; i1 < len(s.Uint16Array); i1++ {
		binary.LittleEndian.PutUint16(b[n:], uint16(s.Uint16Array[i1]))
		n += 2
	}
	for i1 := 0; i1 < len(s.Int32Array); i1++ {
		binary.LittleEndian.PutUint32(b[n:], uint32(s.Int32Array[i1]))
		n += 4
	}
	for i1 := 0; i1 < len(s.Uint32Array); i1++ {
		binary.LittleEndian.PutUint32(b[n:], uint32(s.Uint32Array[i1]))
		n += 4
	}
	for i1 := 0; i1 < len(s.Int64Array); i1++ {
		binary.LittleEndian.PutUint64(b[n:], uint64(s.Int64Array[i1]))
		n += 8
	}
	for i1 := 0; i1 < len(s.Uint64Array); i1++ {
		binary.LittleEndian.PutUint64(b[n:], uint64(s.Uint64Array[i1]))
		n += 8
	}
	for i1 := 0; i1 < len(s.Float32Array); i1++ {
		protocol_PutFloat32(b[n:], float32(s.Float32Array[i1]))
		n += 4
	}
	for i1 := 0; i1 < len(s.Float64Array); i1++ {
		protocol_PutFloat64(b[n:], float64(s.Float64Array[i1]))
		n += 8
	}
	for i1 := 0; i1 < len(s.StringArray); i1++ {
		n += binary.PutUvarint(b[n:], uint64(len(s.StringArray[i1])))
		copy(b[n:], s.StringArray[i1])
		n += len(s.StringArray[i1])
	}
	for i1 := 0; i1 < len(s.BoolArray); i1++ {
		if s.BoolArray[i1] {
			b[n] = 1
		} else {
			b[n] = 0
		}
		n += 1
	}
	return n
}

func (s *FixlenArray) Unmarshal(b []byte) int {
	var n int
	{
		for i1 := 0; i1 < 1; i1++ {
			{
				v, x := binary.Varint(b[n:])
				s.IntArray[i1] = int(v)
				n += x
			}
		}
	}
	{
		for i1 := 0; i1 < 1; i1++ {
			{
				v, x := binary.Uvarint(b[n:])
				s.UintArray[i1] = uint(v)
				n += x
			}
		}
	}
	{
		for i1 := 0; i1 < 1; i1++ {
			s.Int8Array[i1] = int8(b[n])
			n += 1
		}
	}
	{
		copy(s.Uint8Array[:], b[n:n+1])
		n += 1
	}
	{
		for i1 := 0; i1 < 1; i1++ {
			s.Int16Array[i1] = int16(binary.LittleEndian.Uint16(b[n:]))
			n += 2
		}
	}
	{
		for i1 := 0; i1 < 1; i1++ {
			s.Uint16Array[i1] = uint16(binary.LittleEndian.Uint16(b[n:]))
			n += 2
		}
	}
	{
		for i1 := 0; i1 < 1; i1++ {
			s.Int32Array[i1] = int32(binary.LittleEndian.Uint32(b[n:]))
			n += 4
		}
	}
	{
		for i1 := 0; i1 < 1; i1++ {
			s.Uint32Array[i1] = uint32(binary.LittleEndian.Uint32(b[n:]))
			n += 4
		}
	}
	{
		for i1 := 0; i1 < 1; i1++ {
			s.Int64Array[i1] = int64(binary.LittleEndian.Uint64(b[n:]))
			n += 8
		}
	}
	{
		for i1 := 0; i1 < 1; i1++ {
			s.Uint64Array[i1] = uint64(binary.LittleEndian.Uint64(b[n:]))
			n += 8
		}
	}
	{
		for i1 := 0; i1 < 1; i1++ {
			s.Float32Array[i1] = float32(protocol_GetFloat32(b[n:]))
			n += 4
		}
	}
	{
		for i1 := 0; i1 < 1; i1++ {
			s.Float64Array[i1] = float64(protocol_GetFloat64(b[n:]))
			n += 8
		}
	}
	{
		for i1 := 0; i1 < 1; i1++ {
			{
				l, x := binary.Uvarint(b[n:])
				n += x
				s.StringArray[i1] = string(b[n : n+int(l)])
				n += int(l)
			}
		}
	}
	{
		for i1 := 0; i1 < 1; i1++ {
			s.BoolArray[i1] = bool(b[n] == 1)
			n += 1
		}
	}
	return n
}

func (s *Map) Size() int {
	var size int
	size += protocol_UvarintSize(uint64(len(s.IntMap)))
	for key1, val1 := range s.IntMap {
		size += protocol_VarintSize(int64(key1))
		size += protocol_VarintSize(int64(val1))
	}
	size += protocol_UvarintSize(uint64(len(s.UintMap)))
	for key1, val1 := range s.UintMap {
		size += protocol_VarintSize(int64(key1))
		size += protocol_UvarintSize(uint64(val1))
	}
	size += protocol_UvarintSize(uint64(len(s.Int8Map)))
	for key1, _ := range s.Int8Map {
		size += protocol_VarintSize(int64(key1))
		size += 1
	}
	size += protocol_UvarintSize(uint64(len(s.Uint8Map)))
	for key1, _ := range s.Uint8Map {
		size += protocol_VarintSize(int64(key1))
		size += 1
	}
	size += protocol_UvarintSize(uint64(len(s.Int16Map)))
	for key1, _ := range s.Int16Map {
		size += protocol_VarintSize(int64(key1))
		size += 2
	}
	size += protocol_UvarintSize(uint64(len(s.Uint16Map)))
	for key1, _ := range s.Uint16Map {
		size += protocol_VarintSize(int64(key1))
		size += 2
	}
	size += protocol_UvarintSize(uint64(len(s.Int32Map)))
	for key1, _ := range s.Int32Map {
		size += protocol_VarintSize(int64(key1))
		size += 4
	}
	size += protocol_UvarintSize(uint64(len(s.Uint32Map)))
	for key1, _ := range s.Uint32Map {
		size += protocol_VarintSize(int64(key1))
		size += 4
	}
	size += protocol_UvarintSize(uint64(len(s.Int64Map)))
	for key1, _ := range s.Int64Map {
		size += protocol_VarintSize(int64(key1))
		size += 8
	}
	size += protocol_UvarintSize(uint64(len(s.Uint64Map)))
	for key1, _ := range s.Uint64Map {
		size += protocol_VarintSize(int64(key1))
		size += 8
	}
	size += protocol_UvarintSize(uint64(len(s.Float32Map)))
	for key1, _ := range s.Float32Map {
		size += protocol_VarintSize(int64(key1))
		size += 4
	}
	size += protocol_UvarintSize(uint64(len(s.Float64Map)))
	for key1, _ := range s.Float64Map {
		size += protocol_VarintSize(int64(key1))
		size += 8
	}
	size += protocol_UvarintSize(uint64(len(s.StringMap)))
	for key1, val1 := range s.StringMap {
		size += protocol_VarintSize(int64(key1))
		size += protocol_UvarintSize(uint64(len(val1))) + len(val1)
	}
	size += protocol_UvarintSize(uint64(len(s.BoolMap)))
	for key1, _ := range s.BoolMap {
		size += protocol_VarintSize(int64(key1))
		size += 1
	}
	return size
}

func (s *Map) Marshal(b []byte) int {
	var n int
	n += binary.PutUvarint(b[n:], uint64(len(s.IntMap)))
	for key1, val1 := range s.IntMap {
		n += binary.PutVarint(b[n:], int64(key1))
		n += binary.PutVarint(b[n:], int64(val1))
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.UintMap)))
	for key1, val1 := range s.UintMap {
		n += binary.PutVarint(b[n:], int64(key1))
		n += binary.PutUvarint(b[n:], uint64(val1))
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Int8Map)))
	for key1, val1 := range s.Int8Map {
		n += binary.PutVarint(b[n:], int64(key1))
		b[n] = byte(val1)
		n += 1
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Uint8Map)))
	for key1, val1 := range s.Uint8Map {
		n += binary.PutVarint(b[n:], int64(key1))
		b[n] = byte(val1)
		n += 1
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Int16Map)))
	for key1, val1 := range s.Int16Map {
		n += binary.PutVarint(b[n:], int64(key1))
		binary.LittleEndian.PutUint16(b[n:], uint16(val1))
		n += 2
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Uint16Map)))
	for key1, val1 := range s.Uint16Map {
		n += binary.PutVarint(b[n:], int64(key1))
		binary.LittleEndian.PutUint16(b[n:], uint16(val1))
		n += 2
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Int32Map)))
	for key1, val1 := range s.Int32Map {
		n += binary.PutVarint(b[n:], int64(key1))
		binary.LittleEndian.PutUint32(b[n:], uint32(val1))
		n += 4
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Uint32Map)))
	for key1, val1 := range s.Uint32Map {
		n += binary.PutVarint(b[n:], int64(key1))
		binary.LittleEndian.PutUint32(b[n:], uint32(val1))
		n += 4
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Int64Map)))
	for key1, val1 := range s.Int64Map {
		n += binary.PutVarint(b[n:], int64(key1))
		binary.LittleEndian.PutUint64(b[n:], uint64(val1))
		n += 8
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Uint64Map)))
	for key1, val1 := range s.Uint64Map {
		n += binary.PutVarint(b[n:], int64(key1))
		binary.LittleEndian.PutUint64(b[n:], uint64(val1))
		n += 8
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Float32Map)))
	for key1, val1 := range s.Float32Map {
		n += binary.PutVarint(b[n:], int64(key1))
		protocol_PutFloat32(b[n:], float32(val1))
		n += 4
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Float64Map)))
	for key1, val1 := range s.Float64Map {
		n += binary.PutVarint(b[n:], int64(key1))
		protocol_PutFloat64(b[n:], float64(val1))
		n += 8
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.StringMap)))
	for key1, val1 := range s.StringMap {
		n += binary.PutVarint(b[n:], int64(key1))
		n += binary.PutUvarint(b[n:], uint64(len(val1)))
		copy(b[n:], val1)
		n += len(val1)
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.BoolMap)))
	for key1, val1 := range s.BoolMap {
		n += binary.PutVarint(b[n:], int64(key1))
		if val1 {
			b[n] = 1
		} else {
			b[n] = 0
		}
		n += 1
	}
	return n
}

func (s *Map) Unmarshal(b []byte) int {
	var n int
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.IntMap = make(map[int]int, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 int
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			{
				v, x := binary.Varint(b[n:])
				val1 = int(v)
				n += x
			}
			s.IntMap[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.UintMap = make(map[int]uint, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 uint
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			{
				v, x := binary.Uvarint(b[n:])
				val1 = uint(v)
				n += x
			}
			s.UintMap[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Int8Map = make(map[int]int8, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 int8
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			val1 = int8(b[n])
			n += 1
			s.Int8Map[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Uint8Map = make(map[int]uint8, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 uint8
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			val1 = uint8(b[n])
			n += 1
			s.Uint8Map[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Int16Map = make(map[int]int16, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 int16
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			val1 = int16(binary.LittleEndian.Uint16(b[n:]))
			n += 2
			s.Int16Map[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Uint16Map = make(map[int]uint16, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 uint16
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			val1 = uint16(binary.LittleEndian.Uint16(b[n:]))
			n += 2
			s.Uint16Map[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Int32Map = make(map[int]int32, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 int32
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			val1 = int32(binary.LittleEndian.Uint32(b[n:]))
			n += 4
			s.Int32Map[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Uint32Map = make(map[int]uint32, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 uint32
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			val1 = uint32(binary.LittleEndian.Uint32(b[n:]))
			n += 4
			s.Uint32Map[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Int64Map = make(map[int]int64, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 int64
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			val1 = int64(binary.LittleEndian.Uint64(b[n:]))
			n += 8
			s.Int64Map[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Uint64Map = make(map[int]uint64, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 uint64
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			val1 = uint64(binary.LittleEndian.Uint64(b[n:]))
			n += 8
			s.Uint64Map[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Float32Map = make(map[int]float32, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 float32
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			val1 = float32(protocol_GetFloat32(b[n:]))
			n += 4
			s.Float32Map[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Float64Map = make(map[int]float64, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 float64
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			val1 = float64(protocol_GetFloat64(b[n:]))
			n += 8
			s.Float64Map[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.StringMap = make(map[int]string, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 string
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			{
				l, x := binary.Uvarint(b[n:])
				n += x
				val1 = string(b[n : n+int(l)])
				n += int(l)
			}
			s.StringMap[key1] = val1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.BoolMap = make(map[int]bool, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 bool
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			val1 = bool(b[n] == 1)
			n += 1
			s.BoolMap[key1] = val1
		}
	}
	return n
}

func (s *Message) Size() int {
	var size int
	size += s.Scalar.Size()
	size += 1
	if s.ScalarPtr != nil {
		size += s.ScalarPtr.Size()
	}
	size += protocol_UvarintSize(uint64(len(s.ScalarArray)))
	for i1 := 0; i1 < len(s.ScalarArray); i1++ {
		size += s.ScalarArray[i1].Size()
	}
	size += protocol_UvarintSize(uint64(len(s.ScalarMap)))
	for key1, val1 := range s.ScalarMap {
		size += protocol_VarintSize(int64(key1))
		size += 1
		if val1 != nil {
			size += val1.Size()
		}
	}
	return size
}

func (s *Message) Marshal(b []byte) int {
	var n int
	n += s.Scalar.Marshal(b[n:])
	if s.ScalarPtr != nil {
		b[n] = 1
		n++
		n += s.ScalarPtr.Marshal(b[n:])
	} else {
		b[n] = 0
		n++
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.ScalarArray)))
	for i1 := 0; i1 < len(s.ScalarArray); i1++ {
		n += s.ScalarArray[i1].Marshal(b[n:])
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.ScalarMap)))
	for key1, val1 := range s.ScalarMap {
		n += binary.PutVarint(b[n:], int64(key1))
		if val1 != nil {
			b[n] = 1
			n++
			n += val1.Marshal(b[n:])
		} else {
			b[n] = 0
			n++
		}
	}
	return n
}

func (s *Message) Unmarshal(b []byte) int {
	var n int
	n += s.Scalar.Unmarshal(b[n:])
	if b[n] != 0 {
		n += 1
		val1 := new(Scalar)
		n += val1.Unmarshal(b[n:])
		s.ScalarPtr = val1
	} else {
		n += 1
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.ScalarArray = make([]Scalar, l)
		for i1 := 0; i1 < int(l); i1++ {
			n += s.ScalarArray[i1].Unmarshal(b[n:])
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.ScalarMap = make(map[int]*Scalar, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 *Scalar
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			if b[n] != 0 {
				n += 1
				val2 := new(Scalar)
				n += val2.Unmarshal(b[n:])
				val1 = val2
			} else {
				n += 1
			}
			s.ScalarMap[key1] = val1
		}
	}
	return n
}

func (s *Scalar) Size() int {
	var size int
	size += 1
	size += protocol_VarintSize(int64(s.Int))
	size += protocol_UvarintSize(uint64(s.Uint))
	size += 1
	size += 1
	size += 2
	size += 2
	size += 4
	size += 4
	size += 8
	size += 8
	size += 4
	size += 8
	size += protocol_UvarintSize(uint64(len(s.String))) + len(s.String)
	size += protocol_UvarintSize(uint64(len(s.Bytes))) + len(s.Bytes)
	size += 1
	return size
}

func (s *Scalar) Marshal(b []byte) int {
	var n int
	b[n] = byte(s.Byte)
	n += 1
	n += binary.PutVarint(b[n:], int64(s.Int))
	n += binary.PutUvarint(b[n:], uint64(s.Uint))
	b[n] = byte(s.Int8)
	n += 1
	b[n] = byte(s.Uint8)
	n += 1
	binary.LittleEndian.PutUint16(b[n:], uint16(s.Int16))
	n += 2
	binary.LittleEndian.PutUint16(b[n:], uint16(s.Uint16))
	n += 2
	binary.LittleEndian.PutUint32(b[n:], uint32(s.Int32))
	n += 4
	binary.LittleEndian.PutUint32(b[n:], uint32(s.Uint32))
	n += 4
	binary.LittleEndian.PutUint64(b[n:], uint64(s.Int64))
	n += 8
	binary.LittleEndian.PutUint64(b[n:], uint64(s.Uint64))
	n += 8
	protocol_PutFloat32(b[n:], float32(s.Float32))
	n += 4
	protocol_PutFloat64(b[n:], float64(s.Float64))
	n += 8
	n += binary.PutUvarint(b[n:], uint64(len(s.String)))
	copy(b[n:], s.String)
	n += len(s.String)
	n += binary.PutUvarint(b[n:], uint64(len(s.Bytes)))
	copy(b[n:], s.Bytes)
	n += len(s.Bytes)
	if s.Bool {
		b[n] = 1
	} else {
		b[n] = 0
	}
	n += 1
	return n
}

func (s *Scalar) Unmarshal(b []byte) int {
	var n int
	s.Byte = uint8(b[n])
	n += 1
	{
		v, x := binary.Varint(b[n:])
		s.Int = int(v)
		n += x
	}
	{
		v, x := binary.Uvarint(b[n:])
		s.Uint = uint(v)
		n += x
	}
	s.Int8 = int8(b[n])
	n += 1
	s.Uint8 = uint8(b[n])
	n += 1
	s.Int16 = int16(binary.LittleEndian.Uint16(b[n:]))
	n += 2
	s.Uint16 = uint16(binary.LittleEndian.Uint16(b[n:]))
	n += 2
	s.Int32 = int32(binary.LittleEndian.Uint32(b[n:]))
	n += 4
	s.Uint32 = uint32(binary.LittleEndian.Uint32(b[n:]))
	n += 4
	s.Int64 = int64(binary.LittleEndian.Uint64(b[n:]))
	n += 8
	s.Uint64 = uint64(binary.LittleEndian.Uint64(b[n:]))
	n += 8
	s.Float32 = float32(protocol_GetFloat32(b[n:]))
	n += 4
	s.Float64 = float64(protocol_GetFloat64(b[n:]))
	n += 8
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.String = string(b[n : n+int(l)])
		n += int(l)
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Bytes = make([]byte, l)
		copy(s.Bytes, b[n:n+int(l)])
		n += int(l)
	}
	s.Bool = bool(b[n] == 1)
	n += 1
	return n
}

func (s *Pointer) Size() int {
	var size int
	size += 1
	if s.IntPtr != nil {
		size += protocol_VarintSize(int64(*s.IntPtr))
	}
	size += 1
	if s.UintPtr != nil {
		size += protocol_UvarintSize(uint64(*s.UintPtr))
	}
	size += 1
	if s.Int8Ptr != nil {
		size += 1
	}
	size += 1
	if s.Uint8Ptr != nil {
		size += 1
	}
	size += 1
	if s.Int16Ptr != nil {
		size += 2
	}
	size += 1
	if s.Uint16Ptr != nil {
		size += 2
	}
	size += 1
	if s.Int32Ptr != nil {
		size += 4
	}
	size += 1
	if s.Uint32Ptr != nil {
		size += 4
	}
	size += 1
	if s.Int64Ptr != nil {
		size += 8
	}
	size += 1
	if s.Uint64Ptr != nil {
		size += 8
	}
	size += 1
	if s.Float32Ptr != nil {
		size += 4
	}
	size += 1
	if s.Float64Ptr != nil {
		size += 8
	}
	size += 1
	if s.StringPtr != nil {
		size += protocol_UvarintSize(uint64(len(*s.StringPtr))) + len(*s.StringPtr)
	}
	size += 1
	if s.BoolPtr != nil {
		size += 1
	}
	return size
}

func (s *Pointer) Marshal(b []byte) int {
	var n int
	if s.IntPtr != nil {
		b[n] = 1
		n++
		n += binary.PutVarint(b[n:], int64(*s.IntPtr))
	} else {
		b[n] = 0
		n++
	}
	if s.UintPtr != nil {
		b[n] = 1
		n++
		n += binary.PutUvarint(b[n:], uint64(*s.UintPtr))
	} else {
		b[n] = 0
		n++
	}
	if s.Int8Ptr != nil {
		b[n] = 1
		n++
		b[n] = byte(*s.Int8Ptr)
		n += 1
	} else {
		b[n] = 0
		n++
	}
	if s.Uint8Ptr != nil {
		b[n] = 1
		n++
		b[n] = byte(*s.Uint8Ptr)
		n += 1
	} else {
		b[n] = 0
		n++
	}
	if s.Int16Ptr != nil {
		b[n] = 1
		n++
		binary.LittleEndian.PutUint16(b[n:], uint16(*s.Int16Ptr))
		n += 2
	} else {
		b[n] = 0
		n++
	}
	if s.Uint16Ptr != nil {
		b[n] = 1
		n++
		binary.LittleEndian.PutUint16(b[n:], uint16(*s.Uint16Ptr))
		n += 2
	} else {
		b[n] = 0
		n++
	}
	if s.Int32Ptr != nil {
		b[n] = 1
		n++
		binary.LittleEndian.PutUint32(b[n:], uint32(*s.Int32Ptr))
		n += 4
	} else {
		b[n] = 0
		n++
	}
	if s.Uint32Ptr != nil {
		b[n] = 1
		n++
		binary.LittleEndian.PutUint32(b[n:], uint32(*s.Uint32Ptr))
		n += 4
	} else {
		b[n] = 0
		n++
	}
	if s.Int64Ptr != nil {
		b[n] = 1
		n++
		binary.LittleEndian.PutUint64(b[n:], uint64(*s.Int64Ptr))
		n += 8
	} else {
		b[n] = 0
		n++
	}
	if s.Uint64Ptr != nil {
		b[n] = 1
		n++
		binary.LittleEndian.PutUint64(b[n:], uint64(*s.Uint64Ptr))
		n += 8
	} else {
		b[n] = 0
		n++
	}
	if s.Float32Ptr != nil {
		b[n] = 1
		n++
		protocol_PutFloat32(b[n:], float32(*s.Float32Ptr))
		n += 4
	} else {
		b[n] = 0
		n++
	}
	if s.Float64Ptr != nil {
		b[n] = 1
		n++
		protocol_PutFloat64(b[n:], float64(*s.Float64Ptr))
		n += 8
	} else {
		b[n] = 0
		n++
	}
	if s.StringPtr != nil {
		b[n] = 1
		n++
		n += binary.PutUvarint(b[n:], uint64(len(*s.StringPtr)))
		copy(b[n:], *s.StringPtr)
		n += len(*s.StringPtr)
	} else {
		b[n] = 0
		n++
	}
	if s.BoolPtr != nil {
		b[n] = 1
		n++
		if *s.BoolPtr {
			b[n] = 1
		} else {
			b[n] = 0
		}
		n += 1
	} else {
		b[n] = 0
		n++
	}
	return n
}

func (s *Pointer) Unmarshal(b []byte) int {
	var n int
	if b[n] != 0 {
		n += 1
		val1 := new(int)
		{
			v, x := binary.Varint(b[n:])
			*val1 = int(v)
			n += x
		}
		s.IntPtr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(uint)
		{
			v, x := binary.Uvarint(b[n:])
			*val1 = uint(v)
			n += x
		}
		s.UintPtr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(int8)
		*val1 = int8(b[n])
		n += 1
		s.Int8Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(uint8)
		*val1 = uint8(b[n])
		n += 1
		s.Uint8Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(int16)
		*val1 = int16(binary.LittleEndian.Uint16(b[n:]))
		n += 2
		s.Int16Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(uint16)
		*val1 = uint16(binary.LittleEndian.Uint16(b[n:]))
		n += 2
		s.Uint16Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(int32)
		*val1 = int32(binary.LittleEndian.Uint32(b[n:]))
		n += 4
		s.Int32Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(uint32)
		*val1 = uint32(binary.LittleEndian.Uint32(b[n:]))
		n += 4
		s.Uint32Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(int64)
		*val1 = int64(binary.LittleEndian.Uint64(b[n:]))
		n += 8
		s.Int64Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(uint64)
		*val1 = uint64(binary.LittleEndian.Uint64(b[n:]))
		n += 8
		s.Uint64Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(float32)
		*val1 = float32(protocol_GetFloat32(b[n:]))
		n += 4
		s.Float32Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(float64)
		*val1 = float64(protocol_GetFloat64(b[n:]))
		n += 8
		s.Float64Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(string)
		{
			l, x := binary.Uvarint(b[n:])
			n += x
			*val1 = string(b[n : n+int(l)])
			n += int(l)
		}
		s.StringPtr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(bool)
		*val1 = bool(b[n] == 1)
		n += 1
		s.BoolPtr = val1
	} else {
		n += 1
	}
	return n
}

func (s *Array) Size() int {
	var size int
	size += protocol_UvarintSize(uint64(len(s.IntArray)))
	for i1 := 0; i1 < len(s.IntArray); i1++ {
		size += protocol_VarintSize(int64(s.IntArray[i1]))
	}
	size += protocol_UvarintSize(uint64(len(s.UintArray)))
	for i1 := 0; i1 < len(s.UintArray); i1++ {
		size += protocol_UvarintSize(uint64(s.UintArray[i1]))
	}
	size += protocol_UvarintSize(uint64(len(s.Int8Array))) + len(s.Int8Array)*1
	size += protocol_UvarintSize(uint64(len(s.Uint8Array))) + len(s.Uint8Array)
	size += protocol_UvarintSize(uint64(len(s.Int16Array))) + len(s.Int16Array)*2
	size += protocol_UvarintSize(uint64(len(s.Uint16Array))) + len(s.Uint16Array)*2
	size += protocol_UvarintSize(uint64(len(s.Int32Array))) + len(s.Int32Array)*4
	size += protocol_UvarintSize(uint64(len(s.Uint32Array))) + len(s.Uint32Array)*4
	size += protocol_UvarintSize(uint64(len(s.Int64Array))) + len(s.Int64Array)*8
	size += protocol_UvarintSize(uint64(len(s.Uint64Array))) + len(s.Uint64Array)*8
	size += protocol_UvarintSize(uint64(len(s.Float32Array))) + len(s.Float32Array)*4
	size += protocol_UvarintSize(uint64(len(s.Float64Array))) + len(s.Float64Array)*8
	size += protocol_UvarintSize(uint64(len(s.StringArray)))
	for i1 := 0; i1 < len(s.StringArray); i1++ {
		size += protocol_UvarintSize(uint64(len(s.StringArray[i1]))) + len(s.StringArray[i1])
	}
	size += protocol_UvarintSize(uint64(len(s.BoolArray))) + len(s.BoolArray)*1
	return size
}

func (s *Array) Marshal(b []byte) int {
	var n int
	n += binary.PutUvarint(b[n:], uint64(len(s.IntArray)))
	for i1 := 0; i1 < len(s.IntArray); i1++ {
		n += binary.PutVarint(b[n:], int64(s.IntArray[i1]))
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.UintArray)))
	for i1 := 0; i1 < len(s.UintArray); i1++ {
		n += binary.PutUvarint(b[n:], uint64(s.UintArray[i1]))
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Int8Array)))
	for i1 := 0; i1 < len(s.Int8Array); i1++ {
		b[n] = byte(s.Int8Array[i1])
		n += 1
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Uint8Array)))
	copy(b[n:], s.Uint8Array)
	n += len(s.Uint8Array)
	n += binary.PutUvarint(b[n:], uint64(len(s.Int16Array)))
	for i1 := 0; i1 < len(s.Int16Array); i1++ {
		binary.LittleEndian.PutUint16(b[n:], uint16(s.Int16Array[i1]))
		n += 2
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Uint16Array)))
	for i1 := 0; i1 < len(s.Uint16Array); i1++ {
		binary.LittleEndian.PutUint16(b[n:], uint16(s.Uint16Array[i1]))
		n += 2
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Int32Array)))
	for i1 := 0; i1 < len(s.Int32Array); i1++ {
		binary.LittleEndian.PutUint32(b[n:], uint32(s.Int32Array[i1]))
		n += 4
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Uint32Array)))
	for i1 := 0; i1 < len(s.Uint32Array); i1++ {
		binary.LittleEndian.PutUint32(b[n:], uint32(s.Uint32Array[i1]))
		n += 4
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Int64Array)))
	for i1 := 0; i1 < len(s.Int64Array); i1++ {
		binary.LittleEndian.PutUint64(b[n:], uint64(s.Int64Array[i1]))
		n += 8
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Uint64Array)))
	for i1 := 0; i1 < len(s.Uint64Array); i1++ {
		binary.LittleEndian.PutUint64(b[n:], uint64(s.Uint64Array[i1]))
		n += 8
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Float32Array)))
	for i1 := 0; i1 < len(s.Float32Array); i1++ {
		protocol_PutFloat32(b[n:], float32(s.Float32Array[i1]))
		n += 4
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.Float64Array)))
	for i1 := 0; i1 < len(s.Float64Array); i1++ {
		protocol_PutFloat64(b[n:], float64(s.Float64Array[i1]))
		n += 8
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.StringArray)))
	for i1 := 0; i1 < len(s.StringArray); i1++ {
		n += binary.PutUvarint(b[n:], uint64(len(s.StringArray[i1])))
		copy(b[n:], s.StringArray[i1])
		n += len(s.StringArray[i1])
	}
	n += binary.PutUvarint(b[n:], uint64(len(s.BoolArray)))
	for i1 := 0; i1 < len(s.BoolArray); i1++ {
		if s.BoolArray[i1] {
			b[n] = 1
		} else {
			b[n] = 0
		}
		n += 1
	}
	return n
}

func (s *Array) Unmarshal(b []byte) int {
	var n int
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.IntArray = make([]int, l)
		for i1 := 0; i1 < int(l); i1++ {
			{
				v, x := binary.Varint(b[n:])
				s.IntArray[i1] = int(v)
				n += x
			}
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.UintArray = make([]uint, l)
		for i1 := 0; i1 < int(l); i1++ {
			{
				v, x := binary.Uvarint(b[n:])
				s.UintArray[i1] = uint(v)
				n += x
			}
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Int8Array = make([]int8, l)
		for i1 := 0; i1 < int(l); i1++ {
			s.Int8Array[i1] = int8(b[n])
			n += 1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Uint8Array = make([]byte, l)
		copy(s.Uint8Array, b[n:n+int(l)])
		n += int(l)
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Int16Array = make([]int16, l)
		for i1 := 0; i1 < int(l); i1++ {
			s.Int16Array[i1] = int16(binary.LittleEndian.Uint16(b[n:]))
			n += 2
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Uint16Array = make([]uint16, l)
		for i1 := 0; i1 < int(l); i1++ {
			s.Uint16Array[i1] = uint16(binary.LittleEndian.Uint16(b[n:]))
			n += 2
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Int32Array = make([]int32, l)
		for i1 := 0; i1 < int(l); i1++ {
			s.Int32Array[i1] = int32(binary.LittleEndian.Uint32(b[n:]))
			n += 4
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Uint32Array = make([]uint32, l)
		for i1 := 0; i1 < int(l); i1++ {
			s.Uint32Array[i1] = uint32(binary.LittleEndian.Uint32(b[n:]))
			n += 4
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Int64Array = make([]int64, l)
		for i1 := 0; i1 < int(l); i1++ {
			s.Int64Array[i1] = int64(binary.LittleEndian.Uint64(b[n:]))
			n += 8
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Uint64Array = make([]uint64, l)
		for i1 := 0; i1 < int(l); i1++ {
			s.Uint64Array[i1] = uint64(binary.LittleEndian.Uint64(b[n:]))
			n += 8
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Float32Array = make([]float32, l)
		for i1 := 0; i1 < int(l); i1++ {
			s.Float32Array[i1] = float32(protocol_GetFloat32(b[n:]))
			n += 4
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.Float64Array = make([]float64, l)
		for i1 := 0; i1 < int(l); i1++ {
			s.Float64Array[i1] = float64(protocol_GetFloat64(b[n:]))
			n += 8
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.StringArray = make([]string, l)
		for i1 := 0; i1 < int(l); i1++ {
			{
				l, x := binary.Uvarint(b[n:])
				n += x
				s.StringArray[i1] = string(b[n : n+int(l)])
				n += int(l)
			}
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		s.BoolArray = make([]bool, l)
		for i1 := 0; i1 < int(l); i1++ {
			s.BoolArray[i1] = bool(b[n] == 1)
			n += 1
		}
	}
	return n
}

func protocol_UvarintSize(x uint64) int {
	i := 0
	for x >= 0x80 {
		x >>= 7
		i++
	}
	return i + 1
}

func protocol_VarintSize(x int64) int {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return protocol_UvarintSize(ux)
}

func protocol_GetFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(b))
}

func protocol_PutFloat32(b []byte, v float32) {
	binary.LittleEndian.PutUint32(b, math.Float32bits(v))
}

func protocol_GetFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(b))
}

func protocol_PutFloat64(b []byte, v float64) {
	binary.LittleEndian.PutUint64(b, math.Float64bits(v))
}
