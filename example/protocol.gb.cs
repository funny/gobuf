using System;
using System.Collections.Generic;
class ScalarTypes {
	public byte Byte;
	public long Int;
	public ulong Uint;
	public sbyte Int8;
	public byte Uint8;
	public short Int16;
	public ushort Uint16;
	public int Int32;
	public uint Uint32;
	public long Int64;
	public ulong Uint64;
	public float Float32;
	public double Float64;
	public string String;
	public byte[] Bytes;
	public bool Bool;
	public int Size() {
		int size = 0;
		size += 1;
		size += Gobuf.VarintSize(this.Int);
		size += Gobuf.UvarintSize(this.Uint);
		size += 1;
		size += 1;
		size += 2;
		size += 2;
		size += 4;
		size += 4;
		size += 8;
		size += 8;
		size += 4;
		size += 8;
		size += Gobuf.StringSize(this.String);
		size += Gobuf.UvarintSize((ulong)this.Bytes.Length) + this.Bytes.Length;
		size += 1;
		return size;
	}
	public int Marshal(byte[] b, int n) {
		b[n++] = (byte)this.Byte;
		Gobuf.WriteVarint(this.Int, b, ref n);
		Gobuf.WriteUvarint(this.Uint, b, ref n);
		b[n++] = (byte)this.Int8;
		b[n++] = (byte)this.Uint8;
		Gobuf.WriteUint16((ushort)this.Int16, b, ref n);
		Gobuf.WriteUint16((ushort)this.Uint16, b, ref n);
		Gobuf.WriteUint32((uint)this.Int32, b, ref n);
		Gobuf.WriteUint32((uint)this.Uint32, b, ref n);
		Gobuf.WriteUint64((ulong)this.Int64, b, ref n);
		Gobuf.WriteUint64((ulong)this.Uint64, b, ref n);
		Gobuf.WriteFloat32(this.Float32, b, ref n);
		Gobuf.WriteFloat64(this.Float64, b, ref n);
		Gobuf.WriteString(this.String, b, ref n);
		Gobuf.WriteBytes(this.Bytes, b, ref n);
		b[n++] = this.Bool ? (byte)1 : (byte)0;
		return n;
	}
	public int Unmarshal(byte[] b, int n) {
		this.Byte = (byte)b[n++];
		this.Int = Gobuf.ReadVarint(b, ref n);
		this.Uint = Gobuf.ReadUvarint(b, ref n);
		this.Int8 = (sbyte)b[n++];
		this.Uint8 = (byte)b[n++];
		this.Int16 = (short)Gobuf.ReadUint16(b, ref n);
		this.Uint16 = (ushort)Gobuf.ReadUint16(b, ref n);
		this.Int32 = (int)Gobuf.ReadUint32(b, ref n);
		this.Uint32 = (uint)Gobuf.ReadUint32(b, ref n);
		this.Int64 = (long)Gobuf.ReadUint64(b, ref n);
		this.Uint64 = (ulong)Gobuf.ReadUint64(b, ref n);
		this.Float32 = Gobuf.ReadFloat32(b, ref n);
		this.Float64 = Gobuf.ReadFloat64(b, ref n);
		this.String = Gobuf.ReadString(b, ref n);
		this.Bytes = Gobuf.ReadBytes(b, ref n);
		this.Bool = b[n++] == 1;
		return n;
	}
}
class CompositeTypes {
	public Nullable<long> IntPtr;
	public Nullable<ulong> UintPtr;
	public Nullable<sbyte> Int8Ptr;
	public Nullable<byte> Uint8Ptr;
	public Nullable<short> Int16Ptr;
	public Nullable<ushort> Uint16Ptr;
	public Nullable<int> Int32Ptr;
	public Nullable<uint> Uint32Ptr;
	public Nullable<long> Int64Ptr;
	public Nullable<ulong> Uint64Ptr;
	public Nullable<float> Float32Ptr;
	public Nullable<double> Float64Ptr;
	public string StringPtr;
	public List<long> IntArray;
	public List<ulong> UintArray;
	public List<sbyte> Int8Array;
	public byte[] Uint8Array;
	public List<short> Int16Array;
	public List<ushort> Uint16Array;
	public List<int> Int32Array;
	public List<uint> Uint32Array;
	public List<long> Int64Array;
	public List<ulong> Uint64Array;
	public List<float> Float32Array;
	public List<double> Float64Array;
	public List<string> StringArray;
	public List<long> FixLenIntArray = new List<long>(10);
	public List<int> FixLenInt32Array = new List<int>(10);
	public ScalarTypes Message;
	public ScalarTypes MessagePtr;
	public List<ScalarTypes> MessageArray;
	public List<ScalarTypes> MessagePtrArray;
	public List<List<ScalarTypes>> MessageArrayArray;
	public Dictionary<long, Dictionary<long, List<ScalarTypes>>> IntMap;
	public int Size() {
		int size = 0;
		size += 1;
		if (this.IntPtr != null) {
			size += Gobuf.VarintSize(this.IntPtr.Value);
		}
		size += 1;
		if (this.UintPtr != null) {
			size += Gobuf.UvarintSize(this.UintPtr.Value);
		}
		size += 1;
		if (this.Int8Ptr != null) {
			size += 1;
		}
		size += 1;
		if (this.Uint8Ptr != null) {
			size += 1;
		}
		size += 1;
		if (this.Int16Ptr != null) {
			size += 2;
		}
		size += 1;
		if (this.Uint16Ptr != null) {
			size += 2;
		}
		size += 1;
		if (this.Int32Ptr != null) {
			size += 4;
		}
		size += 1;
		if (this.Uint32Ptr != null) {
			size += 4;
		}
		size += 1;
		if (this.Int64Ptr != null) {
			size += 8;
		}
		size += 1;
		if (this.Uint64Ptr != null) {
			size += 8;
		}
		size += 1;
		if (this.Float32Ptr != null) {
			size += 4;
		}
		size += 1;
		if (this.Float64Ptr != null) {
			size += 8;
		}
		size += 1;
		if (this.StringPtr != null) {
			size += Gobuf.StringSize(this.StringPtr);
		}
		size += Gobuf.UvarintSize((ulong)this.IntArray.Count);
		for (var i1 = 0; i1 < this.IntArray.Count; i1 ++) {
			size += Gobuf.VarintSize(this.IntArray[i1]);
		}
		size += Gobuf.UvarintSize((ulong)this.UintArray.Count);
		for (var i1 = 0; i1 < this.UintArray.Count; i1 ++) {
			size += Gobuf.UvarintSize(this.UintArray[i1]);
		}
		size += Gobuf.UvarintSize((ulong)this.Int8Array.Count) + this.Int8Array.Count * 1;
		size += Gobuf.UvarintSize((ulong)this.Uint8Array.Length) + this.Uint8Array.Length;
		size += Gobuf.UvarintSize((ulong)this.Int16Array.Count) + this.Int16Array.Count * 2;
		size += Gobuf.UvarintSize((ulong)this.Uint16Array.Count) + this.Uint16Array.Count * 2;
		size += Gobuf.UvarintSize((ulong)this.Int32Array.Count) + this.Int32Array.Count * 4;
		size += Gobuf.UvarintSize((ulong)this.Uint32Array.Count) + this.Uint32Array.Count * 4;
		size += Gobuf.UvarintSize((ulong)this.Int64Array.Count) + this.Int64Array.Count * 8;
		size += Gobuf.UvarintSize((ulong)this.Uint64Array.Count) + this.Uint64Array.Count * 8;
		size += Gobuf.UvarintSize((ulong)this.Float32Array.Count) + this.Float32Array.Count * 4;
		size += Gobuf.UvarintSize((ulong)this.Float64Array.Count) + this.Float64Array.Count * 8;
		size += Gobuf.UvarintSize((ulong)this.StringArray.Count);
		for (var i1 = 0; i1 < this.StringArray.Count; i1 ++) {
			size += Gobuf.StringSize(this.StringArray[i1]);
		}
		for (var i1 = 0; i1 < this.FixLenIntArray.Count; i1 ++) {
			size += Gobuf.VarintSize(this.FixLenIntArray[i1]);
		}
		size += 10 * 4;
		size += this.Message.Size();
		size += 1;
		if (this.MessagePtr != null) {
			size += this.MessagePtr.Size();
		}
		size += Gobuf.UvarintSize((ulong)this.MessageArray.Count);
		for (var i1 = 0; i1 < this.MessageArray.Count; i1 ++) {
			size += this.MessageArray[i1].Size();
		}
		size += Gobuf.UvarintSize((ulong)this.MessagePtrArray.Count);
		for (var i1 = 0; i1 < this.MessagePtrArray.Count; i1 ++) {
			size += 1;
			if (this.MessagePtrArray[i1] != null) {
				size += this.MessagePtrArray[i1].Size();
			}
		}
		size += Gobuf.UvarintSize((ulong)this.MessageArrayArray.Count);
		for (var i1 = 0; i1 < this.MessageArrayArray.Count; i1 ++) {
			size += Gobuf.UvarintSize((ulong)this.MessageArrayArray[i1].Count);
			for (var i2 = 0; i2 < this.MessageArrayArray[i1].Count; i2 ++) {
				size += this.MessageArrayArray[i1][i2].Size();
			}
		}
		size += Gobuf.UvarintSize((ulong)this.IntMap.Count);
		foreach (var item1 in this.IntMap) {
			size += Gobuf.VarintSize(item1.Key);
			size += Gobuf.UvarintSize((ulong)item1.Value.Count);
			foreach (var item2 in item1.Value) {
				size += Gobuf.VarintSize(item2.Key);
				size += Gobuf.UvarintSize((ulong)item2.Value.Count);
				for (var i3 = 0; i3 < item2.Value.Count; i3 ++) {
					size += 1;
					if (item2.Value[i3] != null) {
						size += item2.Value[i3].Size();
					}
				}
			}
		}
		return size;
	}
	public int Marshal(byte[] b, int n) {
		if (this.IntPtr != null) {
			b[n++] = 1;
			Gobuf.WriteVarint(this.IntPtr.Value, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.UintPtr != null) {
			b[n++] = 1;
			Gobuf.WriteUvarint(this.UintPtr.Value, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Int8Ptr != null) {
			b[n++] = 1;
			b[n++] = (byte)this.Int8Ptr.Value;
		} else {
			b[n++] = 0;
		}
		if (this.Uint8Ptr != null) {
			b[n++] = 1;
			b[n++] = (byte)this.Uint8Ptr.Value;
		} else {
			b[n++] = 0;
		}
		if (this.Int16Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteUint16((ushort)this.Int16Ptr.Value, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Uint16Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteUint16((ushort)this.Uint16Ptr.Value, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Int32Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteUint32((uint)this.Int32Ptr.Value, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Uint32Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteUint32((uint)this.Uint32Ptr.Value, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Int64Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteUint64((ulong)this.Int64Ptr.Value, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Uint64Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteUint64((ulong)this.Uint64Ptr.Value, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Float32Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteFloat32(this.Float32Ptr.Value, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Float64Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteFloat64(this.Float64Ptr.Value, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.StringPtr != null) {
			b[n++] = 1;
			Gobuf.WriteString(this.StringPtr, b, ref n);
		} else {
			b[n++] = 0;
		}
		Gobuf.WriteUvarint((ulong)this.IntArray.Count, b, ref n);
		for (var i1 = 0; i1 < this.IntArray.Count; i1 ++) {
			Gobuf.WriteVarint(this.IntArray[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.UintArray.Count, b, ref n);
		for (var i1 = 0; i1 < this.UintArray.Count; i1 ++) {
			Gobuf.WriteUvarint(this.UintArray[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Int8Array.Count, b, ref n);
		for (var i1 = 0; i1 < this.Int8Array.Count; i1 ++) {
			b[n++] = (byte)this.Int8Array[i1];
		}
		Gobuf.WriteBytes(this.Uint8Array, b, ref n);
		Gobuf.WriteUvarint((ulong)this.Int16Array.Count, b, ref n);
		for (var i1 = 0; i1 < this.Int16Array.Count; i1 ++) {
			Gobuf.WriteUint16((ushort)this.Int16Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Uint16Array.Count, b, ref n);
		for (var i1 = 0; i1 < this.Uint16Array.Count; i1 ++) {
			Gobuf.WriteUint16((ushort)this.Uint16Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Int32Array.Count, b, ref n);
		for (var i1 = 0; i1 < this.Int32Array.Count; i1 ++) {
			Gobuf.WriteUint32((uint)this.Int32Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Uint32Array.Count, b, ref n);
		for (var i1 = 0; i1 < this.Uint32Array.Count; i1 ++) {
			Gobuf.WriteUint32((uint)this.Uint32Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Int64Array.Count, b, ref n);
		for (var i1 = 0; i1 < this.Int64Array.Count; i1 ++) {
			Gobuf.WriteUint64((ulong)this.Int64Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Uint64Array.Count, b, ref n);
		for (var i1 = 0; i1 < this.Uint64Array.Count; i1 ++) {
			Gobuf.WriteUint64((ulong)this.Uint64Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Float32Array.Count, b, ref n);
		for (var i1 = 0; i1 < this.Float32Array.Count; i1 ++) {
			Gobuf.WriteFloat32(this.Float32Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Float64Array.Count, b, ref n);
		for (var i1 = 0; i1 < this.Float64Array.Count; i1 ++) {
			Gobuf.WriteFloat64(this.Float64Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.StringArray.Count, b, ref n);
		for (var i1 = 0; i1 < this.StringArray.Count; i1 ++) {
			Gobuf.WriteString(this.StringArray[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.FixLenIntArray.Count; i1 ++) {
			Gobuf.WriteVarint(this.FixLenIntArray[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.FixLenInt32Array.Count; i1 ++) {
			Gobuf.WriteUint32((uint)this.FixLenInt32Array[i1], b, ref n);
		}
		n = this.Message.Marshal(b, n);
		if (this.MessagePtr != null) {
			b[n++] = 1;
			n = this.MessagePtr.Marshal(b, n);
		} else {
			b[n++] = 0;
		}
		Gobuf.WriteUvarint((ulong)this.MessageArray.Count, b, ref n);
		for (var i1 = 0; i1 < this.MessageArray.Count; i1 ++) {
			n = this.MessageArray[i1].Marshal(b, n);
		}
		Gobuf.WriteUvarint((ulong)this.MessagePtrArray.Count, b, ref n);
		for (var i1 = 0; i1 < this.MessagePtrArray.Count; i1 ++) {
			if (this.MessagePtrArray[i1] != null) {
				b[n++] = 1;
				n = this.MessagePtrArray[i1].Marshal(b, n);
			} else {
				b[n++] = 0;
			}
		}
		Gobuf.WriteUvarint((ulong)this.MessageArrayArray.Count, b, ref n);
		for (var i1 = 0; i1 < this.MessageArrayArray.Count; i1 ++) {
			Gobuf.WriteUvarint((ulong)this.MessageArrayArray[i1].Count, b, ref n);
			for (var i2 = 0; i2 < this.MessageArrayArray[i1].Count; i2 ++) {
				n = this.MessageArrayArray[i1][i2].Marshal(b, n);
			}
		}
		Gobuf.WriteUvarint((ulong)this.IntMap.Count, b, ref n);
		foreach (var item1 in this.IntMap) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteUvarint((ulong)item1.Value.Count, b, ref n);
			foreach (var item2 in item1.Value) {
				Gobuf.WriteVarint(item2.Key, b, ref n);
				Gobuf.WriteUvarint((ulong)item2.Value.Count, b, ref n);
				for (var i3 = 0; i3 < item2.Value.Count; i3 ++) {
					if (item2.Value[i3] != null) {
						b[n++] = 1;
						n = item2.Value[i3].Marshal(b, n);
					} else {
						b[n++] = 0;
					}
				}
			}
		}
		return n;
	}
	public int Unmarshal(byte[] b, int n) {
		if (b[n++] != 0) {
			this.IntPtr = Gobuf.ReadVarint(b, ref n);
		}
		if (b[n++] != 0) {
			this.UintPtr = Gobuf.ReadUvarint(b, ref n);
		}
		if (b[n++] != 0) {
			this.Int8Ptr = (sbyte)b[n++];
		}
		if (b[n++] != 0) {
			this.Uint8Ptr = (byte)b[n++];
		}
		if (b[n++] != 0) {
			this.Int16Ptr = (short)Gobuf.ReadUint16(b, ref n);
		}
		if (b[n++] != 0) {
			this.Uint16Ptr = (ushort)Gobuf.ReadUint16(b, ref n);
		}
		if (b[n++] != 0) {
			this.Int32Ptr = (int)Gobuf.ReadUint32(b, ref n);
		}
		if (b[n++] != 0) {
			this.Uint32Ptr = (uint)Gobuf.ReadUint32(b, ref n);
		}
		if (b[n++] != 0) {
			this.Int64Ptr = (long)Gobuf.ReadUint64(b, ref n);
		}
		if (b[n++] != 0) {
			this.Uint64Ptr = (ulong)Gobuf.ReadUint64(b, ref n);
		}
		if (b[n++] != 0) {
			this.Float32Ptr = Gobuf.ReadFloat32(b, ref n);
		}
		if (b[n++] != 0) {
			this.Float64Ptr = Gobuf.ReadFloat64(b, ref n);
		}
		if (b[n++] != 0) {
			this.StringPtr = Gobuf.ReadString(b, ref n);
		}
		{
			this.IntArray = new List<long>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.IntArray.Count; i1 ++) {
				this.IntArray[i1] = Gobuf.ReadVarint(b, ref n);
			}
		}
		{
			this.UintArray = new List<ulong>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.UintArray.Count; i1 ++) {
				this.UintArray[i1] = Gobuf.ReadUvarint(b, ref n);
			}
		}
		{
			this.Int8Array = new List<sbyte>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Int8Array.Count; i1 ++) {
				this.Int8Array[i1] = (sbyte)b[n++];
			}
		}
		this.Uint8Array = Gobuf.ReadBytes(b, ref n);
		{
			this.Int16Array = new List<short>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Int16Array.Count; i1 ++) {
				this.Int16Array[i1] = (short)Gobuf.ReadUint16(b, ref n);
			}
		}
		{
			this.Uint16Array = new List<ushort>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Uint16Array.Count; i1 ++) {
				this.Uint16Array[i1] = (ushort)Gobuf.ReadUint16(b, ref n);
			}
		}
		{
			this.Int32Array = new List<int>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Int32Array.Count; i1 ++) {
				this.Int32Array[i1] = (int)Gobuf.ReadUint32(b, ref n);
			}
		}
		{
			this.Uint32Array = new List<uint>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Uint32Array.Count; i1 ++) {
				this.Uint32Array[i1] = (uint)Gobuf.ReadUint32(b, ref n);
			}
		}
		{
			this.Int64Array = new List<long>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Int64Array.Count; i1 ++) {
				this.Int64Array[i1] = (long)Gobuf.ReadUint64(b, ref n);
			}
		}
		{
			this.Uint64Array = new List<ulong>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Uint64Array.Count; i1 ++) {
				this.Uint64Array[i1] = (ulong)Gobuf.ReadUint64(b, ref n);
			}
		}
		{
			this.Float32Array = new List<float>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Float32Array.Count; i1 ++) {
				this.Float32Array[i1] = Gobuf.ReadFloat32(b, ref n);
			}
		}
		{
			this.Float64Array = new List<double>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Float64Array.Count; i1 ++) {
				this.Float64Array[i1] = Gobuf.ReadFloat64(b, ref n);
			}
		}
		{
			this.StringArray = new List<string>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.StringArray.Count; i1 ++) {
				this.StringArray[i1] = Gobuf.ReadString(b, ref n);
			}
		}
		{
			for (var i1 = 0; i1 < this.FixLenIntArray.Count; i1 ++) {
				this.FixLenIntArray[i1] = Gobuf.ReadVarint(b, ref n);
			}
		}
		{
			for (var i1 = 0; i1 < this.FixLenInt32Array.Count; i1 ++) {
				this.FixLenInt32Array[i1] = (int)Gobuf.ReadUint32(b, ref n);
			}
		}
		this.Message = new ScalarTypes();
		n = this.Message.Unmarshal(b, n);
		if (b[n++] != 0) {
			this.MessagePtr = new ScalarTypes();
			n = this.MessagePtr.Unmarshal(b, n);
		}
		{
			this.MessageArray = new List<ScalarTypes>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.MessageArray.Count; i1 ++) {
				this.MessageArray[i1] = new ScalarTypes();
				n = this.MessageArray[i1].Unmarshal(b, n);
			}
		}
		{
			this.MessagePtrArray = new List<ScalarTypes>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.MessagePtrArray.Count; i1 ++) {
				if (b[n++] != 0) {
					this.MessagePtrArray[i1] = new ScalarTypes();
					n = this.MessagePtrArray[i1].Unmarshal(b, n);
				}
			}
		}
		{
			this.MessageArrayArray = new List<List<ScalarTypes>>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.MessageArrayArray.Count; i1 ++) {
				{
					this.MessageArrayArray[i1] = new List<ScalarTypes>((int)Gobuf.ReadUvarint(b, ref n));
					for (var i2 = 0; i2 < this.MessageArrayArray[i1].Count; i2 ++) {
						this.MessageArrayArray[i1][i2] = new ScalarTypes();
						n = this.MessageArrayArray[i1][i2].Unmarshal(b, n);
					}
				}
			}
		}
		{
			this.IntMap = new Dictionary<long, Dictionary<long, List<ScalarTypes>>>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.IntMap.Count; i1 ++) {
				long key1;
				Dictionary<long, List<ScalarTypes>> val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				{
					val1 = new Dictionary<long, List<ScalarTypes>>((int)Gobuf.ReadUvarint(b, ref n));
					for (var i2 = 0; i2 < val1.Count; i2 ++) {
						long key2;
						List<ScalarTypes> val2;
						key2 = Gobuf.ReadVarint(b, ref n);
						{
							val2 = new List<ScalarTypes>((int)Gobuf.ReadUvarint(b, ref n));
							for (var i3 = 0; i3 < val2.Count; i3 ++) {
								if (b[n++] != 0) {
									val2[i3] = new ScalarTypes();
									n = val2[i3].Unmarshal(b, n);
								}
							}
						}
						val1[key2] = val2;
					}
				}
				this.IntMap[key1] = val1;
			}
		}
		return n;
	}
}
