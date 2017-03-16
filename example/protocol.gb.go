package example

import "github.com/hotgo/gobuf"
import "encoding/binary"

func (msg *ScalarTypes) Size() int {
	var size int
	size += 1
	size += gobuf.VarintSize(int64(msg.Int))
	size += gobuf.UvarintSize(uint64(msg.Uint))
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
	size += gobuf.UvarintSize(uint64(len(msg.String))) + len(msg.String)
	size += gobuf.UvarintSize(uint64(len(msg.Bytes))) + len(msg.Bytes)
	return size
}

func (msg *ScalarTypes) Marshal(b []byte) int {
	var n int
	b[n] = byte(msg.Byte)
	n += 1
	n += binary.PutVarint(b[n:], int64(msg.Int))
	n += binary.PutUvarint(b[n:], uint64(msg.Uint))
	b[n] = byte(msg.Int8)
	n += 1
	b[n] = byte(msg.Uint8)
	n += 1
	binary.LittleEndian.PutUint16(b[n:], uint16(msg.Int16))
	n += 2
	binary.LittleEndian.PutUint16(b[n:], uint16(msg.Uint16))
	n += 2
	binary.LittleEndian.PutUint32(b[n:], uint32(msg.Int32))
	n += 4
	binary.LittleEndian.PutUint32(b[n:], uint32(msg.Uint32))
	n += 4
	binary.LittleEndian.PutUint64(b[n:], uint64(msg.Int64))
	n += 8
	binary.LittleEndian.PutUint64(b[n:], uint64(msg.Uint64))
	n += 8
	gobuf.PutFloat32(b[n:], float32(msg.Float32))
	n += 4
	gobuf.PutFloat64(b[n:], float64(msg.Float64))
	n += 8
	n += binary.PutUvarint(b[n:], uint64(len(msg.String)))
	copy(b[n:], msg.String)
	n += len(msg.String)
	n += binary.PutUvarint(b[n:], uint64(len(msg.Bytes)))
	copy(b[n:], msg.Bytes)
	n += len(msg.Bytes)
	return n
}

func (msg *ScalarTypes) Unmarshal(b []byte) int {
	var n int
	msg.Byte = uint8(b[n])
	n += 1
	{
		v, x := binary.Varint(b[n:])
		msg.Int = int(v)
		n += x
	}
	{
		v, x := binary.Uvarint(b[n:])
		msg.Uint = uint(v)
		n += x
	}
	msg.Int8 = int8(b[n])
	n += 1
	msg.Uint8 = uint8(b[n])
	n += 1
	msg.Int16 = int16(binary.LittleEndian.Uint16(b[n:]))
	n += 2
	msg.Uint16 = uint16(binary.LittleEndian.Uint16(b[n:]))
	n += 2
	msg.Int32 = int32(binary.LittleEndian.Uint32(b[n:]))
	n += 4
	msg.Uint32 = uint32(binary.LittleEndian.Uint32(b[n:]))
	n += 4
	msg.Int64 = int64(binary.LittleEndian.Uint64(b[n:]))
	n += 8
	msg.Uint64 = uint64(binary.LittleEndian.Uint64(b[n:]))
	n += 8
	msg.Float32 = float32(gobuf.GetFloat32(b[n:]))
	n += 4
	msg.Float64 = float64(gobuf.GetFloat64(b[n:]))
	n += 8
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.String = string(b[n : n+int(l)])
		n += int(l)
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Bytes = make([]byte, l)
		copy(msg.Bytes, b[n:n+int(l)])
		n += int(l)
	}
	return n
}

func (msg *CompositeTypes) Size() int {
	var size int
	size += 1
	if msg.IntPtr != nil {
		size += gobuf.VarintSize(int64(*msg.IntPtr))
	}
	size += 1
	if msg.UintPtr != nil {
		size += gobuf.UvarintSize(uint64(*msg.UintPtr))
	}
	size += 1
	if msg.Int8Ptr != nil {
		size += 1
	}
	size += 1
	if msg.Uint8Ptr != nil {
		size += 1
	}
	size += 1
	if msg.Int16Ptr != nil {
		size += 2
	}
	size += 1
	if msg.Uint16Ptr != nil {
		size += 2
	}
	size += 1
	if msg.Int32Ptr != nil {
		size += 4
	}
	size += 1
	if msg.Uint32Ptr != nil {
		size += 4
	}
	size += 1
	if msg.Int64Ptr != nil {
		size += 8
	}
	size += 1
	if msg.Uint64Ptr != nil {
		size += 8
	}
	size += 1
	if msg.Float32Ptr != nil {
		size += 4
	}
	size += 1
	if msg.Float64Ptr != nil {
		size += 8
	}
	size += 1
	if msg.StringPtr != nil {
		size += gobuf.UvarintSize(uint64(len(*msg.StringPtr))) + len(*msg.StringPtr)
	}
	gobuf.UvarintSize(uint64(len(msg.IntArray)))
	for i1 := 0; i1 < len(msg.IntArray); i1++ {
		size += gobuf.VarintSize(int64(msg.IntArray[i1]))
	}
	gobuf.UvarintSize(uint64(len(msg.UintArray)))
	for i1 := 0; i1 < len(msg.UintArray); i1++ {
		size += gobuf.UvarintSize(uint64(msg.UintArray[i1]))
	}
	size += gobuf.UvarintSize(uint64(len(msg.Int8Array))) + len(msg.Int8Array)*1
	size += gobuf.UvarintSize(uint64(len(msg.Uint8Array))) + len(msg.Uint8Array)
	size += gobuf.UvarintSize(uint64(len(msg.Int16Array))) + len(msg.Int16Array)*2
	size += gobuf.UvarintSize(uint64(len(msg.Uint16Array))) + len(msg.Uint16Array)*2
	size += gobuf.UvarintSize(uint64(len(msg.Int32Array))) + len(msg.Int32Array)*4
	size += gobuf.UvarintSize(uint64(len(msg.Uint32Array))) + len(msg.Uint32Array)*4
	size += gobuf.UvarintSize(uint64(len(msg.Int64Array))) + len(msg.Int64Array)*8
	size += gobuf.UvarintSize(uint64(len(msg.Uint64Array))) + len(msg.Uint64Array)*8
	size += gobuf.UvarintSize(uint64(len(msg.Float32Array))) + len(msg.Float32Array)*4
	size += gobuf.UvarintSize(uint64(len(msg.Float64Array))) + len(msg.Float64Array)*8
	gobuf.UvarintSize(uint64(len(msg.StringArray)))
	for i1 := 0; i1 < len(msg.StringArray); i1++ {
		size += gobuf.UvarintSize(uint64(len(msg.StringArray[i1]))) + len(msg.StringArray[i1])
	}
	size += msg.Message.Size()
	size += 1
	if msg.MessagePtr != nil {
		size += msg.MessagePtr.Size()
	}
	gobuf.UvarintSize(uint64(len(msg.MessageArray)))
	for i1 := 0; i1 < len(msg.MessageArray); i1++ {
		size += msg.MessageArray[i1].Size()
	}
	gobuf.UvarintSize(uint64(len(msg.MessagePtrArray)))
	for i1 := 0; i1 < len(msg.MessagePtrArray); i1++ {
		size += 1
		if msg.MessagePtrArray[i1] != nil {
			size += msg.MessagePtrArray[i1].Size()
		}
	}
	gobuf.UvarintSize(uint64(len(msg.MessageArrayArray)))
	for i1 := 0; i1 < len(msg.MessageArrayArray); i1++ {
		gobuf.UvarintSize(uint64(len(msg.MessageArrayArray[i1])))
		for i2 := 0; i2 < len(msg.MessageArrayArray[i1]); i2++ {
			size += msg.MessageArrayArray[i1][i2].Size()
		}
	}
	size += gobuf.UvarintSize(uint64(len(msg.IntMap)))
	for key1, val1 := range msg.IntMap {
		size += gobuf.VarintSize(int64(key1))
		size += gobuf.UvarintSize(uint64(len(val1)))
		for key2, val2 := range val1 {
			size += gobuf.VarintSize(int64(key2))
			gobuf.UvarintSize(uint64(len(val2)))
			for i3 := 0; i3 < len(val2); i3++ {
				size += 1
				if val2[i3] != nil {
					size += val2[i3].Size()
				}
			}
		}
	}
	return size
}

func (msg *CompositeTypes) Marshal(b []byte) int {
	var n int
	if msg.IntPtr != nil {
		b[n] = 1
		n++
		n += binary.PutVarint(b[n:], int64(*msg.IntPtr))
	} else {
		b[n] = 0
		n++
	}
	if msg.UintPtr != nil {
		b[n] = 1
		n++
		n += binary.PutUvarint(b[n:], uint64(*msg.UintPtr))
	} else {
		b[n] = 0
		n++
	}
	if msg.Int8Ptr != nil {
		b[n] = 1
		n++
		b[n] = byte(*msg.Int8Ptr)
		n += 1
	} else {
		b[n] = 0
		n++
	}
	if msg.Uint8Ptr != nil {
		b[n] = 1
		n++
		b[n] = byte(*msg.Uint8Ptr)
		n += 1
	} else {
		b[n] = 0
		n++
	}
	if msg.Int16Ptr != nil {
		b[n] = 1
		n++
		binary.LittleEndian.PutUint16(b[n:], uint16(*msg.Int16Ptr))
		n += 2
	} else {
		b[n] = 0
		n++
	}
	if msg.Uint16Ptr != nil {
		b[n] = 1
		n++
		binary.LittleEndian.PutUint16(b[n:], uint16(*msg.Uint16Ptr))
		n += 2
	} else {
		b[n] = 0
		n++
	}
	if msg.Int32Ptr != nil {
		b[n] = 1
		n++
		binary.LittleEndian.PutUint32(b[n:], uint32(*msg.Int32Ptr))
		n += 4
	} else {
		b[n] = 0
		n++
	}
	if msg.Uint32Ptr != nil {
		b[n] = 1
		n++
		binary.LittleEndian.PutUint32(b[n:], uint32(*msg.Uint32Ptr))
		n += 4
	} else {
		b[n] = 0
		n++
	}
	if msg.Int64Ptr != nil {
		b[n] = 1
		n++
		binary.LittleEndian.PutUint64(b[n:], uint64(*msg.Int64Ptr))
		n += 8
	} else {
		b[n] = 0
		n++
	}
	if msg.Uint64Ptr != nil {
		b[n] = 1
		n++
		binary.LittleEndian.PutUint64(b[n:], uint64(*msg.Uint64Ptr))
		n += 8
	} else {
		b[n] = 0
		n++
	}
	if msg.Float32Ptr != nil {
		b[n] = 1
		n++
		gobuf.PutFloat32(b[n:], float32(*msg.Float32Ptr))
		n += 4
	} else {
		b[n] = 0
		n++
	}
	if msg.Float64Ptr != nil {
		b[n] = 1
		n++
		gobuf.PutFloat64(b[n:], float64(*msg.Float64Ptr))
		n += 8
	} else {
		b[n] = 0
		n++
	}
	if msg.StringPtr != nil {
		b[n] = 1
		n++
		n += binary.PutUvarint(b[n:], uint64(len(*msg.StringPtr)))
		copy(b[n:], *msg.StringPtr)
		n += len(*msg.StringPtr)
	} else {
		b[n] = 0
		n++
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.IntArray)))
	for i1 := 0; i1 < len(msg.IntArray); i1++ {
		n += binary.PutVarint(b[n:], int64(msg.IntArray[i1]))
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.UintArray)))
	for i1 := 0; i1 < len(msg.UintArray); i1++ {
		n += binary.PutUvarint(b[n:], uint64(msg.UintArray[i1]))
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Int8Array)))
	for i1 := 0; i1 < len(msg.Int8Array); i1++ {
		b[n] = byte(msg.Int8Array[i1])
		n += 1
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Uint8Array)))
	copy(b[n:], msg.Uint8Array)
	n += len(msg.Uint8Array)
	n += binary.PutUvarint(b[n:], uint64(len(msg.Int16Array)))
	for i1 := 0; i1 < len(msg.Int16Array); i1++ {
		binary.LittleEndian.PutUint16(b[n:], uint16(msg.Int16Array[i1]))
		n += 2
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Uint16Array)))
	for i1 := 0; i1 < len(msg.Uint16Array); i1++ {
		binary.LittleEndian.PutUint16(b[n:], uint16(msg.Uint16Array[i1]))
		n += 2
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Int32Array)))
	for i1 := 0; i1 < len(msg.Int32Array); i1++ {
		binary.LittleEndian.PutUint32(b[n:], uint32(msg.Int32Array[i1]))
		n += 4
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Uint32Array)))
	for i1 := 0; i1 < len(msg.Uint32Array); i1++ {
		binary.LittleEndian.PutUint32(b[n:], uint32(msg.Uint32Array[i1]))
		n += 4
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Int64Array)))
	for i1 := 0; i1 < len(msg.Int64Array); i1++ {
		binary.LittleEndian.PutUint64(b[n:], uint64(msg.Int64Array[i1]))
		n += 8
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Uint64Array)))
	for i1 := 0; i1 < len(msg.Uint64Array); i1++ {
		binary.LittleEndian.PutUint64(b[n:], uint64(msg.Uint64Array[i1]))
		n += 8
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Float32Array)))
	for i1 := 0; i1 < len(msg.Float32Array); i1++ {
		gobuf.PutFloat32(b[n:], float32(msg.Float32Array[i1]))
		n += 4
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Float64Array)))
	for i1 := 0; i1 < len(msg.Float64Array); i1++ {
		gobuf.PutFloat64(b[n:], float64(msg.Float64Array[i1]))
		n += 8
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.StringArray)))
	for i1 := 0; i1 < len(msg.StringArray); i1++ {
		n += binary.PutUvarint(b[n:], uint64(len(msg.StringArray[i1])))
		copy(b[n:], msg.StringArray[i1])
		n += len(msg.StringArray[i1])
	}
	n += msg.Message.Marshal(b[n:])
	if msg.MessagePtr != nil {
		b[n] = 1
		n++
		n += msg.MessagePtr.Marshal(b[n:])
	} else {
		b[n] = 0
		n++
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.MessageArray)))
	for i1 := 0; i1 < len(msg.MessageArray); i1++ {
		n += msg.MessageArray[i1].Marshal(b[n:])
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.MessagePtrArray)))
	for i1 := 0; i1 < len(msg.MessagePtrArray); i1++ {
		if msg.MessagePtrArray[i1] != nil {
			b[n] = 1
			n++
			n += msg.MessagePtrArray[i1].Marshal(b[n:])
		} else {
			b[n] = 0
			n++
		}
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.MessageArrayArray)))
	for i1 := 0; i1 < len(msg.MessageArrayArray); i1++ {
		n += binary.PutUvarint(b[n:], uint64(len(msg.MessageArrayArray[i1])))
		for i2 := 0; i2 < len(msg.MessageArrayArray[i1]); i2++ {
			n += msg.MessageArrayArray[i1][i2].Marshal(b[n:])
		}
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.IntMap)))
	for key1, val1 := range msg.IntMap {
		n += binary.PutVarint(b[n:], int64(key1))
		n += binary.PutUvarint(b[n:], uint64(len(val1)))
		for key2, val2 := range val1 {
			n += binary.PutVarint(b[n:], int64(key2))
			n += binary.PutUvarint(b[n:], uint64(len(val2)))
			for i3 := 0; i3 < len(val2); i3++ {
				if val2[i3] != nil {
					b[n] = 1
					n++
					n += val2[i3].Marshal(b[n:])
				} else {
					b[n] = 0
					n++
				}
			}
		}
	}
	return n
}

func (msg *CompositeTypes) Unmarshal(b []byte) int {
	var n int
	if b[n] != 0 {
		n += 1
		val1 := new(int)
		{
			v, x := binary.Varint(b[n:])
			*val1 = int(v)
			n += x
		}
		msg.IntPtr = val1
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
		msg.UintPtr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(int8)
		*val1 = int8(b[n])
		n += 1
		msg.Int8Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(uint8)
		*val1 = uint8(b[n])
		n += 1
		msg.Uint8Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(int16)
		*val1 = int16(binary.LittleEndian.Uint16(b[n:]))
		n += 2
		msg.Int16Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(uint16)
		*val1 = uint16(binary.LittleEndian.Uint16(b[n:]))
		n += 2
		msg.Uint16Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(int32)
		*val1 = int32(binary.LittleEndian.Uint32(b[n:]))
		n += 4
		msg.Int32Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(uint32)
		*val1 = uint32(binary.LittleEndian.Uint32(b[n:]))
		n += 4
		msg.Uint32Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(int64)
		*val1 = int64(binary.LittleEndian.Uint64(b[n:]))
		n += 8
		msg.Int64Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(uint64)
		*val1 = uint64(binary.LittleEndian.Uint64(b[n:]))
		n += 8
		msg.Uint64Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(float32)
		*val1 = float32(gobuf.GetFloat32(b[n:]))
		n += 4
		msg.Float32Ptr = val1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		val1 := new(float64)
		*val1 = float64(gobuf.GetFloat64(b[n:]))
		n += 8
		msg.Float64Ptr = val1
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
		msg.StringPtr = val1
	} else {
		n += 1
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.IntArray = make([]int, l)
		for i1 := 0; i1 < int(l); i1++ {
			{
				v, x := binary.Varint(b[n:])
				msg.IntArray[i1] = int(v)
				n += x
			}
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.UintArray = make([]uint, l)
		for i1 := 0; i1 < int(l); i1++ {
			{
				v, x := binary.Uvarint(b[n:])
				msg.UintArray[i1] = uint(v)
				n += x
			}
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Int8Array = make([]int8, l)
		for i1 := 0; i1 < int(l); i1++ {
			msg.Int8Array[i1] = int8(b[n])
			n += 1
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Uint8Array = make([]byte, l)
		copy(msg.Uint8Array, b[n:n+int(l)])
		n += int(l)
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Int16Array = make([]int16, l)
		for i1 := 0; i1 < int(l); i1++ {
			msg.Int16Array[i1] = int16(binary.LittleEndian.Uint16(b[n:]))
			n += 2
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Uint16Array = make([]uint16, l)
		for i1 := 0; i1 < int(l); i1++ {
			msg.Uint16Array[i1] = uint16(binary.LittleEndian.Uint16(b[n:]))
			n += 2
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Int32Array = make([]int32, l)
		for i1 := 0; i1 < int(l); i1++ {
			msg.Int32Array[i1] = int32(binary.LittleEndian.Uint32(b[n:]))
			n += 4
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Uint32Array = make([]uint32, l)
		for i1 := 0; i1 < int(l); i1++ {
			msg.Uint32Array[i1] = uint32(binary.LittleEndian.Uint32(b[n:]))
			n += 4
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Int64Array = make([]int64, l)
		for i1 := 0; i1 < int(l); i1++ {
			msg.Int64Array[i1] = int64(binary.LittleEndian.Uint64(b[n:]))
			n += 8
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Uint64Array = make([]uint64, l)
		for i1 := 0; i1 < int(l); i1++ {
			msg.Uint64Array[i1] = uint64(binary.LittleEndian.Uint64(b[n:]))
			n += 8
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Float32Array = make([]float32, l)
		for i1 := 0; i1 < int(l); i1++ {
			msg.Float32Array[i1] = float32(gobuf.GetFloat32(b[n:]))
			n += 4
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Float64Array = make([]float64, l)
		for i1 := 0; i1 < int(l); i1++ {
			msg.Float64Array[i1] = float64(gobuf.GetFloat64(b[n:]))
			n += 8
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.StringArray = make([]string, l)
		for i1 := 0; i1 < int(l); i1++ {
			{
				l, x := binary.Uvarint(b[n:])
				n += x
				msg.StringArray[i1] = string(b[n : n+int(l)])
				n += int(l)
			}
		}
	}
	n += msg.Message.Unmarshal(b[n:])
	if b[n] != 0 {
		n += 1
		val1 := new(ScalarTypes)
		n += val1.Unmarshal(b[n:])
		msg.MessagePtr = val1
	} else {
		n += 1
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.MessageArray = make([]ScalarTypes, l)
		for i1 := 0; i1 < int(l); i1++ {
			n += msg.MessageArray[i1].Unmarshal(b[n:])
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.MessagePtrArray = make([]*ScalarTypes, l)
		for i1 := 0; i1 < int(l); i1++ {
			if b[n] != 0 {
				n += 1
				val2 := new(ScalarTypes)
				n += val2.Unmarshal(b[n:])
				msg.MessagePtrArray[i1] = val2
			} else {
				n += 1
			}
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.MessageArrayArray = make([][]ScalarTypes, l)
		for i1 := 0; i1 < int(l); i1++ {
			{
				l, x := binary.Uvarint(b[n:])
				n += x
				msg.MessageArrayArray[i1] = make([]ScalarTypes, l)
				for i2 := 0; i2 < int(l); i2++ {
					n += msg.MessageArrayArray[i1][i2].Unmarshal(b[n:])
				}
			}
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.IntMap = make(map[int]map[int][]*ScalarTypes, l)
		for i1 := 0; i1 < int(l); i1++ {
			var key1 int
			var val1 map[int][]*ScalarTypes
			{
				v, x := binary.Varint(b[n:])
				key1 = int(v)
				n += x
			}
			{
				l, x := binary.Uvarint(b[n:])
				n += x
				val1 = make(map[int][]*ScalarTypes, l)
				for i2 := 0; i2 < int(l); i2++ {
					var key2 int
					var val2 []*ScalarTypes
					{
						v, x := binary.Varint(b[n:])
						key2 = int(v)
						n += x
					}
					{
						l, x := binary.Uvarint(b[n:])
						n += x
						val2 = make([]*ScalarTypes, l)
						for i3 := 0; i3 < int(l); i3++ {
							if b[n] != 0 {
								n += 1
								val4 := new(ScalarTypes)
								n += val4.Unmarshal(b[n:])
								val2[i3] = val4
							} else {
								n += 1
							}
						}
					}
					val1[key2] = val2
				}
			}
			msg.IntMap[key1] = val1
		}
	}
	return n
}
