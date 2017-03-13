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
	for i := 0; i < len(msg.IntArray); i++ {
		size += gobuf.VarintSize(int64(msg.IntArray[i]))
	}
	for i := 0; i < len(msg.UintArray); i++ {
		size += gobuf.UvarintSize(uint64(msg.UintArray[i]))
	}
	size += 1 * len(msg.Int8Array)
	size += gobuf.UvarintSize(uint64(len(msg.Uint8Array))) + len(msg.Uint8Array)
	size += 2 * len(msg.Int16Array)
	size += 2 * len(msg.Uint16Array)
	size += 4 * len(msg.Int32Array)
	size += 4 * len(msg.Uint32Array)
	size += 8 * len(msg.Int64Array)
	size += 8 * len(msg.Uint64Array)
	size += 4 * len(msg.Float32Array)
	size += 8 * len(msg.Float64Array)
	for i := 0; i < len(msg.StringArray); i++ {
		size += gobuf.UvarintSize(uint64(len(msg.StringArray[i]))) + len(msg.StringArray[i])
	}
	size += msg.Message.Size()
	size += 1
	if msg.MessagePtr != nil {
		size += msg.MessagePtr.Size()
	}
	for i := 0; i < len(msg.MessageArray); i++ {
		size += msg.MessageArray[i].Size()
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
	for i := 0; i < len(msg.IntArray); i++ {
		n += binary.PutVarint(b[n:], int64(msg.IntArray[i]))
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.UintArray)))
	for i := 0; i < len(msg.UintArray); i++ {
		n += binary.PutUvarint(b[n:], uint64(msg.UintArray[i]))
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Int8Array)))
	for i := 0; i < len(msg.Int8Array); i++ {
		b[n] = byte(msg.Int8Array[i])
		n += 1
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Uint8Array)))
	copy(b[n:], msg.Uint8Array)
	n += len(msg.Uint8Array)
	n += binary.PutUvarint(b[n:], uint64(len(msg.Int16Array)))
	for i := 0; i < len(msg.Int16Array); i++ {
		binary.LittleEndian.PutUint16(b[n:], uint16(msg.Int16Array[i]))
		n += 2
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Uint16Array)))
	for i := 0; i < len(msg.Uint16Array); i++ {
		binary.LittleEndian.PutUint16(b[n:], uint16(msg.Uint16Array[i]))
		n += 2
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Int32Array)))
	for i := 0; i < len(msg.Int32Array); i++ {
		binary.LittleEndian.PutUint32(b[n:], uint32(msg.Int32Array[i]))
		n += 4
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Uint32Array)))
	for i := 0; i < len(msg.Uint32Array); i++ {
		binary.LittleEndian.PutUint32(b[n:], uint32(msg.Uint32Array[i]))
		n += 4
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Int64Array)))
	for i := 0; i < len(msg.Int64Array); i++ {
		binary.LittleEndian.PutUint64(b[n:], uint64(msg.Int64Array[i]))
		n += 8
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Uint64Array)))
	for i := 0; i < len(msg.Uint64Array); i++ {
		binary.LittleEndian.PutUint64(b[n:], uint64(msg.Uint64Array[i]))
		n += 8
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Float32Array)))
	for i := 0; i < len(msg.Float32Array); i++ {
		gobuf.PutFloat32(b[n:], float32(msg.Float32Array[i]))
		n += 4
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.Float64Array)))
	for i := 0; i < len(msg.Float64Array); i++ {
		gobuf.PutFloat64(b[n:], float64(msg.Float64Array[i]))
		n += 8
	}
	n += binary.PutUvarint(b[n:], uint64(len(msg.StringArray)))
	for i := 0; i < len(msg.StringArray); i++ {
		n += binary.PutUvarint(b[n:], uint64(len(msg.StringArray[i])))
		copy(b[n:], msg.StringArray[i])
		n += len(msg.StringArray[i])
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
	for i := 0; i < len(msg.MessageArray); i++ {
		n += msg.MessageArray[i].Marshal(b[n:])
	}
	return n
}

func (msg *CompositeTypes) Unmarshal(b []byte) int {
	var n int
	if b[n] != 0 {
		n += 1
		msg.IntPtr = new(int)
		{
			v, x := binary.Varint(b[n:])
			*msg.IntPtr = int(v)
			n += x
		}
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		msg.UintPtr = new(uint)
		{
			v, x := binary.Uvarint(b[n:])
			*msg.UintPtr = uint(v)
			n += x
		}
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		msg.Int8Ptr = new(int8)
		*msg.Int8Ptr = int8(b[n])
		n += 1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		msg.Uint8Ptr = new(uint8)
		*msg.Uint8Ptr = uint8(b[n])
		n += 1
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		msg.Int16Ptr = new(int16)
		*msg.Int16Ptr = int16(binary.LittleEndian.Uint16(b[n:]))
		n += 2
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		msg.Uint16Ptr = new(uint16)
		*msg.Uint16Ptr = uint16(binary.LittleEndian.Uint16(b[n:]))
		n += 2
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		msg.Int32Ptr = new(int32)
		*msg.Int32Ptr = int32(binary.LittleEndian.Uint32(b[n:]))
		n += 4
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		msg.Uint32Ptr = new(uint32)
		*msg.Uint32Ptr = uint32(binary.LittleEndian.Uint32(b[n:]))
		n += 4
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		msg.Int64Ptr = new(int64)
		*msg.Int64Ptr = int64(binary.LittleEndian.Uint64(b[n:]))
		n += 8
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		msg.Uint64Ptr = new(uint64)
		*msg.Uint64Ptr = uint64(binary.LittleEndian.Uint64(b[n:]))
		n += 8
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		msg.Float32Ptr = new(float32)
		*msg.Float32Ptr = float32(gobuf.GetFloat32(b[n:]))
		n += 4
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		msg.Float64Ptr = new(float64)
		*msg.Float64Ptr = float64(gobuf.GetFloat64(b[n:]))
		n += 8
	} else {
		n += 1
	}
	if b[n] != 0 {
		n += 1
		msg.StringPtr = new(string)
		{
			l, x := binary.Uvarint(b[n:])
			n += x
			*msg.StringPtr = string(b[n : n+int(l)])
			n += int(l)
		}
	} else {
		n += 1
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.IntArray = make([]int, l)
		for i := 0; i < int(l); i++ {
			{
				v, x := binary.Varint(b[n:])
				msg.IntArray[i] = int(v)
				n += x
			}
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.UintArray = make([]uint, l)
		for i := 0; i < int(l); i++ {
			{
				v, x := binary.Uvarint(b[n:])
				msg.UintArray[i] = uint(v)
				n += x
			}
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Int8Array = make([]int8, l)
		for i := 0; i < int(l); i++ {
			msg.Int8Array[i] = int8(b[n])
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
		for i := 0; i < int(l); i++ {
			msg.Int16Array[i] = int16(binary.LittleEndian.Uint16(b[n:]))
			n += 2
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Uint16Array = make([]uint16, l)
		for i := 0; i < int(l); i++ {
			msg.Uint16Array[i] = uint16(binary.LittleEndian.Uint16(b[n:]))
			n += 2
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Int32Array = make([]int32, l)
		for i := 0; i < int(l); i++ {
			msg.Int32Array[i] = int32(binary.LittleEndian.Uint32(b[n:]))
			n += 4
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Uint32Array = make([]uint32, l)
		for i := 0; i < int(l); i++ {
			msg.Uint32Array[i] = uint32(binary.LittleEndian.Uint32(b[n:]))
			n += 4
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Int64Array = make([]int64, l)
		for i := 0; i < int(l); i++ {
			msg.Int64Array[i] = int64(binary.LittleEndian.Uint64(b[n:]))
			n += 8
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Uint64Array = make([]uint64, l)
		for i := 0; i < int(l); i++ {
			msg.Uint64Array[i] = uint64(binary.LittleEndian.Uint64(b[n:]))
			n += 8
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Float32Array = make([]float32, l)
		for i := 0; i < int(l); i++ {
			msg.Float32Array[i] = float32(gobuf.GetFloat32(b[n:]))
			n += 4
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.Float64Array = make([]float64, l)
		for i := 0; i < int(l); i++ {
			msg.Float64Array[i] = float64(gobuf.GetFloat64(b[n:]))
			n += 8
		}
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.StringArray = make([]string, l)
		for i := 0; i < int(l); i++ {
			{
				l, x := binary.Uvarint(b[n:])
				n += x
				msg.StringArray[i] = string(b[n : n+int(l)])
				n += int(l)
			}
		}
	}
	n += msg.Message.Unmarshal(b[n:])
	if b[n] != 0 {
		n += 1
		msg.MessagePtr = new(ScalarTypes)
		n += msg.MessagePtr.Unmarshal(b[n:])
	} else {
		n += 1
	}
	{
		l, x := binary.Uvarint(b[n:])
		n += x
		msg.MessageArray = make([]ScalarTypes, l)
		for i := 0; i < int(l); i++ {
			n += msg.MessageArray[i].Unmarshal(b[n:])
		}
	}
	return n
}
