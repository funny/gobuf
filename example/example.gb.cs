using System;
using System.Collections.Generic;
class Pointer {
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
	public Nullable<bool> BoolPtr;
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
		size += 1;
		if (this.BoolPtr != null) {
			size += 1;
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
		if (this.BoolPtr != null) {
			b[n++] = 1;
			b[n++] = this.BoolPtr.Value ? (byte)1 : (byte)0;
		} else {
			b[n++] = 0;
		}
		return n;
	}
	public int Unmarshal(byte[] b, int n) {
		if (b[n++] != 0) {
			this.IntPtr = Gobuf.ReadVarint(b, ref n);
		} else {
			this.IntPtr = null;
		}
		if (b[n++] != 0) {
			this.UintPtr = Gobuf.ReadUvarint(b, ref n);
		} else {
			this.UintPtr = null;
		}
		if (b[n++] != 0) {
			this.Int8Ptr = (sbyte)b[n++];
		} else {
			this.Int8Ptr = null;
		}
		if (b[n++] != 0) {
			this.Uint8Ptr = (byte)b[n++];
		} else {
			this.Uint8Ptr = null;
		}
		if (b[n++] != 0) {
			this.Int16Ptr = (short)Gobuf.ReadUint16(b, ref n);
		} else {
			this.Int16Ptr = null;
		}
		if (b[n++] != 0) {
			this.Uint16Ptr = (ushort)Gobuf.ReadUint16(b, ref n);
		} else {
			this.Uint16Ptr = null;
		}
		if (b[n++] != 0) {
			this.Int32Ptr = (int)Gobuf.ReadUint32(b, ref n);
		} else {
			this.Int32Ptr = null;
		}
		if (b[n++] != 0) {
			this.Uint32Ptr = (uint)Gobuf.ReadUint32(b, ref n);
		} else {
			this.Uint32Ptr = null;
		}
		if (b[n++] != 0) {
			this.Int64Ptr = (long)Gobuf.ReadUint64(b, ref n);
		} else {
			this.Int64Ptr = null;
		}
		if (b[n++] != 0) {
			this.Uint64Ptr = (ulong)Gobuf.ReadUint64(b, ref n);
		} else {
			this.Uint64Ptr = null;
		}
		if (b[n++] != 0) {
			this.Float32Ptr = Gobuf.ReadFloat32(b, ref n);
		} else {
			this.Float32Ptr = null;
		}
		if (b[n++] != 0) {
			this.Float64Ptr = Gobuf.ReadFloat64(b, ref n);
		} else {
			this.Float64Ptr = null;
		}
		if (b[n++] != 0) {
			this.StringPtr = Gobuf.ReadString(b, ref n);
		} else {
			this.StringPtr = null;
		}
		if (b[n++] != 0) {
			this.BoolPtr = b[n++] == 1;
		} else {
			this.BoolPtr = null;
		}
		return n;
	}
}
class Array {
	public List<long> IntArray = new List<long>();
	public List<ulong> UintArray = new List<ulong>();
	public List<sbyte> Int8Array = new List<sbyte>();
	public byte[] Uint8Array;
	public List<short> Int16Array = new List<short>();
	public List<ushort> Uint16Array = new List<ushort>();
	public List<int> Int32Array = new List<int>();
	public List<uint> Uint32Array = new List<uint>();
	public List<long> Int64Array = new List<long>();
	public List<ulong> Uint64Array = new List<ulong>();
	public List<float> Float32Array = new List<float>();
	public List<double> Float64Array = new List<double>();
	public List<string> StringArray = new List<string>();
	public List<bool> BoolArray = new List<bool>();
	public int Size() {
		int size = 0;
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
		size += Gobuf.UvarintSize((ulong)this.BoolArray.Count) + this.BoolArray.Count * 1;
		return size;
	}
	public int Marshal(byte[] b, int n) {
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
		Gobuf.WriteUvarint((ulong)this.BoolArray.Count, b, ref n);
		for (var i1 = 0; i1 < this.BoolArray.Count; i1 ++) {
			b[n++] = this.BoolArray[i1] ? (byte)1 : (byte)0;
		}
		return n;
	}
	public int Unmarshal(byte[] b, int n) {
		{
			this.IntArray = new List<long>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.IntArray.Capacity; i1 ++) {
				long v1;
				v1 = Gobuf.ReadVarint(b, ref n);
				this.IntArray.Add(v1);
			}
		}
		{
			this.UintArray = new List<ulong>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.UintArray.Capacity; i1 ++) {
				ulong v1;
				v1 = Gobuf.ReadUvarint(b, ref n);
				this.UintArray.Add(v1);
			}
		}
		{
			this.Int8Array = new List<sbyte>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Int8Array.Capacity; i1 ++) {
				sbyte v1;
				v1 = (sbyte)b[n++];
				this.Int8Array.Add(v1);
			}
		}
		this.Uint8Array = Gobuf.ReadBytes(b, ref n);
		{
			this.Int16Array = new List<short>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Int16Array.Capacity; i1 ++) {
				short v1;
				v1 = (short)Gobuf.ReadUint16(b, ref n);
				this.Int16Array.Add(v1);
			}
		}
		{
			this.Uint16Array = new List<ushort>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Uint16Array.Capacity; i1 ++) {
				ushort v1;
				v1 = (ushort)Gobuf.ReadUint16(b, ref n);
				this.Uint16Array.Add(v1);
			}
		}
		{
			this.Int32Array = new List<int>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Int32Array.Capacity; i1 ++) {
				int v1;
				v1 = (int)Gobuf.ReadUint32(b, ref n);
				this.Int32Array.Add(v1);
			}
		}
		{
			this.Uint32Array = new List<uint>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Uint32Array.Capacity; i1 ++) {
				uint v1;
				v1 = (uint)Gobuf.ReadUint32(b, ref n);
				this.Uint32Array.Add(v1);
			}
		}
		{
			this.Int64Array = new List<long>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Int64Array.Capacity; i1 ++) {
				long v1;
				v1 = (long)Gobuf.ReadUint64(b, ref n);
				this.Int64Array.Add(v1);
			}
		}
		{
			this.Uint64Array = new List<ulong>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Uint64Array.Capacity; i1 ++) {
				ulong v1;
				v1 = (ulong)Gobuf.ReadUint64(b, ref n);
				this.Uint64Array.Add(v1);
			}
		}
		{
			this.Float32Array = new List<float>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Float32Array.Capacity; i1 ++) {
				float v1;
				v1 = Gobuf.ReadFloat32(b, ref n);
				this.Float32Array.Add(v1);
			}
		}
		{
			this.Float64Array = new List<double>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.Float64Array.Capacity; i1 ++) {
				double v1;
				v1 = Gobuf.ReadFloat64(b, ref n);
				this.Float64Array.Add(v1);
			}
		}
		{
			this.StringArray = new List<string>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.StringArray.Capacity; i1 ++) {
				string v1;
				v1 = Gobuf.ReadString(b, ref n);
				this.StringArray.Add(v1);
			}
		}
		{
			this.BoolArray = new List<bool>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.BoolArray.Capacity; i1 ++) {
				bool v1;
				v1 = b[n++] == 1;
				this.BoolArray.Add(v1);
			}
		}
		return n;
	}
}
class FixlenArray {
	public long[] IntArray = new long[1];
	public ulong[] UintArray = new ulong[1];
	public sbyte[] Int8Array = new sbyte[1];
	public byte[] Uint8Array = new byte[1];
	public short[] Int16Array = new short[1];
	public ushort[] Uint16Array = new ushort[1];
	public int[] Int32Array = new int[1];
	public uint[] Uint32Array = new uint[1];
	public long[] Int64Array = new long[1];
	public ulong[] Uint64Array = new ulong[1];
	public float[] Float32Array = new float[1];
	public double[] Float64Array = new double[1];
	public string[] StringArray = new string[1];
	public bool[] BoolArray = new bool[1];
	public int Size() {
		int size = 0;
		for (var i1 = 0; i1 < this.IntArray.Length; i1 ++) {
			size += Gobuf.VarintSize(this.IntArray[i1]);
		}
		for (var i1 = 0; i1 < this.UintArray.Length; i1 ++) {
			size += Gobuf.UvarintSize(this.UintArray[i1]);
		}
		size += 1 * 1;
		size += Gobuf.UvarintSize((ulong)this.Uint8Array.Length) + this.Uint8Array.Length;
		size += 1 * 2;
		size += 1 * 2;
		size += 1 * 4;
		size += 1 * 4;
		size += 1 * 8;
		size += 1 * 8;
		size += 1 * 4;
		size += 1 * 8;
		for (var i1 = 0; i1 < this.StringArray.Length; i1 ++) {
			size += Gobuf.StringSize(this.StringArray[i1]);
		}
		size += 1 * 1;
		return size;
	}
	public int Marshal(byte[] b, int n) {
		for (var i1 = 0; i1 < this.IntArray.Length; i1 ++) {
			Gobuf.WriteVarint(this.IntArray[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.UintArray.Length; i1 ++) {
			Gobuf.WriteUvarint(this.UintArray[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.Int8Array.Length; i1 ++) {
			b[n++] = (byte)this.Int8Array[i1];
		}
		Gobuf.WriteBytes(this.Uint8Array, b, ref n);
		for (var i1 = 0; i1 < this.Int16Array.Length; i1 ++) {
			Gobuf.WriteUint16((ushort)this.Int16Array[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.Uint16Array.Length; i1 ++) {
			Gobuf.WriteUint16((ushort)this.Uint16Array[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.Int32Array.Length; i1 ++) {
			Gobuf.WriteUint32((uint)this.Int32Array[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.Uint32Array.Length; i1 ++) {
			Gobuf.WriteUint32((uint)this.Uint32Array[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.Int64Array.Length; i1 ++) {
			Gobuf.WriteUint64((ulong)this.Int64Array[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.Uint64Array.Length; i1 ++) {
			Gobuf.WriteUint64((ulong)this.Uint64Array[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.Float32Array.Length; i1 ++) {
			Gobuf.WriteFloat32(this.Float32Array[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.Float64Array.Length; i1 ++) {
			Gobuf.WriteFloat64(this.Float64Array[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.StringArray.Length; i1 ++) {
			Gobuf.WriteString(this.StringArray[i1], b, ref n);
		}
		for (var i1 = 0; i1 < this.BoolArray.Length; i1 ++) {
			b[n++] = this.BoolArray[i1] ? (byte)1 : (byte)0;
		}
		return n;
	}
	public int Unmarshal(byte[] b, int n) {
		this.IntArray = new long[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.IntArray[i1] = Gobuf.ReadVarint(b, ref n);
		}
		this.UintArray = new ulong[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.UintArray[i1] = Gobuf.ReadUvarint(b, ref n);
		}
		this.Int8Array = new sbyte[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.Int8Array[i1] = (sbyte)b[n++];
		}
		this.Uint8Array = Gobuf.ReadBytes(b, ref n);
		this.Int16Array = new short[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.Int16Array[i1] = (short)Gobuf.ReadUint16(b, ref n);
		}
		this.Uint16Array = new ushort[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.Uint16Array[i1] = (ushort)Gobuf.ReadUint16(b, ref n);
		}
		this.Int32Array = new int[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.Int32Array[i1] = (int)Gobuf.ReadUint32(b, ref n);
		}
		this.Uint32Array = new uint[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.Uint32Array[i1] = (uint)Gobuf.ReadUint32(b, ref n);
		}
		this.Int64Array = new long[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.Int64Array[i1] = (long)Gobuf.ReadUint64(b, ref n);
		}
		this.Uint64Array = new ulong[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.Uint64Array[i1] = (ulong)Gobuf.ReadUint64(b, ref n);
		}
		this.Float32Array = new float[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.Float32Array[i1] = Gobuf.ReadFloat32(b, ref n);
		}
		this.Float64Array = new double[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.Float64Array[i1] = Gobuf.ReadFloat64(b, ref n);
		}
		this.StringArray = new string[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.StringArray[i1] = Gobuf.ReadString(b, ref n);
		}
		this.BoolArray = new bool[1];
		for (var i1 = 0; i1 < 1; i1 ++) {
			this.BoolArray[i1] = b[n++] == 1;
		}
		return n;
	}
}
class Map {
	public Dictionary<long, long> IntMap = new Dictionary<long, long>();
	public Dictionary<long, ulong> UintMap = new Dictionary<long, ulong>();
	public Dictionary<long, sbyte> Int8Map = new Dictionary<long, sbyte>();
	public Dictionary<long, byte> Uint8Map = new Dictionary<long, byte>();
	public Dictionary<long, short> Int16Map = new Dictionary<long, short>();
	public Dictionary<long, ushort> Uint16Map = new Dictionary<long, ushort>();
	public Dictionary<long, int> Int32Map = new Dictionary<long, int>();
	public Dictionary<long, uint> Uint32Map = new Dictionary<long, uint>();
	public Dictionary<long, long> Int64Map = new Dictionary<long, long>();
	public Dictionary<long, ulong> Uint64Map = new Dictionary<long, ulong>();
	public Dictionary<long, float> Float32Map = new Dictionary<long, float>();
	public Dictionary<long, double> Float64Map = new Dictionary<long, double>();
	public Dictionary<long, string> StringMap = new Dictionary<long, string>();
	public Dictionary<long, bool> BoolMap = new Dictionary<long, bool>();
	public int Size() {
		int size = 0;
		size += Gobuf.UvarintSize((ulong)this.IntMap.Count);
		foreach (var item1 in this.IntMap) {
			size += Gobuf.VarintSize(item1.Key);
			size += Gobuf.VarintSize(item1.Value);
		}
		size += Gobuf.UvarintSize((ulong)this.UintMap.Count);
		foreach (var item1 in this.UintMap) {
			size += Gobuf.VarintSize(item1.Key);
			size += Gobuf.UvarintSize(item1.Value);
		}
		size += Gobuf.UvarintSize((ulong)this.Int8Map.Count);
		foreach (var item1 in this.Int8Map) {
			size += Gobuf.VarintSize(item1.Key);
			size += 1;
		}
		size += Gobuf.UvarintSize((ulong)this.Uint8Map.Count);
		foreach (var item1 in this.Uint8Map) {
			size += Gobuf.VarintSize(item1.Key);
			size += 1;
		}
		size += Gobuf.UvarintSize((ulong)this.Int16Map.Count);
		foreach (var item1 in this.Int16Map) {
			size += Gobuf.VarintSize(item1.Key);
			size += 2;
		}
		size += Gobuf.UvarintSize((ulong)this.Uint16Map.Count);
		foreach (var item1 in this.Uint16Map) {
			size += Gobuf.VarintSize(item1.Key);
			size += 2;
		}
		size += Gobuf.UvarintSize((ulong)this.Int32Map.Count);
		foreach (var item1 in this.Int32Map) {
			size += Gobuf.VarintSize(item1.Key);
			size += 4;
		}
		size += Gobuf.UvarintSize((ulong)this.Uint32Map.Count);
		foreach (var item1 in this.Uint32Map) {
			size += Gobuf.VarintSize(item1.Key);
			size += 4;
		}
		size += Gobuf.UvarintSize((ulong)this.Int64Map.Count);
		foreach (var item1 in this.Int64Map) {
			size += Gobuf.VarintSize(item1.Key);
			size += 8;
		}
		size += Gobuf.UvarintSize((ulong)this.Uint64Map.Count);
		foreach (var item1 in this.Uint64Map) {
			size += Gobuf.VarintSize(item1.Key);
			size += 8;
		}
		size += Gobuf.UvarintSize((ulong)this.Float32Map.Count);
		foreach (var item1 in this.Float32Map) {
			size += Gobuf.VarintSize(item1.Key);
			size += 4;
		}
		size += Gobuf.UvarintSize((ulong)this.Float64Map.Count);
		foreach (var item1 in this.Float64Map) {
			size += Gobuf.VarintSize(item1.Key);
			size += 8;
		}
		size += Gobuf.UvarintSize((ulong)this.StringMap.Count);
		foreach (var item1 in this.StringMap) {
			size += Gobuf.VarintSize(item1.Key);
			size += Gobuf.StringSize(item1.Value);
		}
		size += Gobuf.UvarintSize((ulong)this.BoolMap.Count);
		foreach (var item1 in this.BoolMap) {
			size += Gobuf.VarintSize(item1.Key);
			size += 1;
		}
		return size;
	}
	public int Marshal(byte[] b, int n) {
		Gobuf.WriteUvarint((ulong)this.IntMap.Count, b, ref n);
		foreach (var item1 in this.IntMap) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteVarint(item1.Value, b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.UintMap.Count, b, ref n);
		foreach (var item1 in this.UintMap) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteUvarint(item1.Value, b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Int8Map.Count, b, ref n);
		foreach (var item1 in this.Int8Map) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			b[n++] = (byte)item1.Value;
		}
		Gobuf.WriteUvarint((ulong)this.Uint8Map.Count, b, ref n);
		foreach (var item1 in this.Uint8Map) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			b[n++] = (byte)item1.Value;
		}
		Gobuf.WriteUvarint((ulong)this.Int16Map.Count, b, ref n);
		foreach (var item1 in this.Int16Map) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteUint16((ushort)item1.Value, b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Uint16Map.Count, b, ref n);
		foreach (var item1 in this.Uint16Map) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteUint16((ushort)item1.Value, b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Int32Map.Count, b, ref n);
		foreach (var item1 in this.Int32Map) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteUint32((uint)item1.Value, b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Uint32Map.Count, b, ref n);
		foreach (var item1 in this.Uint32Map) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteUint32((uint)item1.Value, b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Int64Map.Count, b, ref n);
		foreach (var item1 in this.Int64Map) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteUint64((ulong)item1.Value, b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Uint64Map.Count, b, ref n);
		foreach (var item1 in this.Uint64Map) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteUint64((ulong)item1.Value, b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Float32Map.Count, b, ref n);
		foreach (var item1 in this.Float32Map) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteFloat32(item1.Value, b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Float64Map.Count, b, ref n);
		foreach (var item1 in this.Float64Map) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteFloat64(item1.Value, b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.StringMap.Count, b, ref n);
		foreach (var item1 in this.StringMap) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteString(item1.Value, b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.BoolMap.Count, b, ref n);
		foreach (var item1 in this.BoolMap) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			b[n++] = item1.Value ? (byte)1 : (byte)0;
		}
		return n;
	}
	public int Unmarshal(byte[] b, int n) {
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.IntMap = new Dictionary<long, long>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				long val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = Gobuf.ReadVarint(b, ref n);
				this.IntMap.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.UintMap = new Dictionary<long, ulong>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				ulong val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = Gobuf.ReadUvarint(b, ref n);
				this.UintMap.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.Int8Map = new Dictionary<long, sbyte>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				sbyte val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = (sbyte)b[n++];
				this.Int8Map.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.Uint8Map = new Dictionary<long, byte>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				byte val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = (byte)b[n++];
				this.Uint8Map.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.Int16Map = new Dictionary<long, short>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				short val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = (short)Gobuf.ReadUint16(b, ref n);
				this.Int16Map.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.Uint16Map = new Dictionary<long, ushort>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				ushort val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = (ushort)Gobuf.ReadUint16(b, ref n);
				this.Uint16Map.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.Int32Map = new Dictionary<long, int>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				int val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = (int)Gobuf.ReadUint32(b, ref n);
				this.Int32Map.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.Uint32Map = new Dictionary<long, uint>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				uint val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = (uint)Gobuf.ReadUint32(b, ref n);
				this.Uint32Map.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.Int64Map = new Dictionary<long, long>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				long val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = (long)Gobuf.ReadUint64(b, ref n);
				this.Int64Map.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.Uint64Map = new Dictionary<long, ulong>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				ulong val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = (ulong)Gobuf.ReadUint64(b, ref n);
				this.Uint64Map.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.Float32Map = new Dictionary<long, float>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				float val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = Gobuf.ReadFloat32(b, ref n);
				this.Float32Map.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.Float64Map = new Dictionary<long, double>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				double val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = Gobuf.ReadFloat64(b, ref n);
				this.Float64Map.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.StringMap = new Dictionary<long, string>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				string val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = Gobuf.ReadString(b, ref n);
				this.StringMap.Add(key1, val1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.BoolMap = new Dictionary<long, bool>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				bool val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				val1 = b[n++] == 1;
				this.BoolMap.Add(key1, val1);
			}
		}
		return n;
	}
}
class Message {
	public Scalar Scalar;
	public Scalar ScalarPtr;
	public List<Scalar> ScalarArray = new List<Scalar>();
	public Dictionary<long, Scalar> ScalarMap = new Dictionary<long, Scalar>();
	public int Size() {
		int size = 0;
		size += this.Scalar.Size();
		size += 1;
		if (this.ScalarPtr != null) {
			size += this.ScalarPtr.Size();
		}
		size += Gobuf.UvarintSize((ulong)this.ScalarArray.Count);
		for (var i1 = 0; i1 < this.ScalarArray.Count; i1 ++) {
			size += this.ScalarArray[i1].Size();
		}
		size += Gobuf.UvarintSize((ulong)this.ScalarMap.Count);
		foreach (var item1 in this.ScalarMap) {
			size += Gobuf.VarintSize(item1.Key);
			size += 1;
			if (item1.Value != null) {
				size += item1.Value.Size();
			}
		}
		return size;
	}
	public int Marshal(byte[] b, int n) {
		n = this.Scalar.Marshal(b, n);
		if (this.ScalarPtr != null) {
			b[n++] = 1;
			n = this.ScalarPtr.Marshal(b, n);
		} else {
			b[n++] = 0;
		}
		Gobuf.WriteUvarint((ulong)this.ScalarArray.Count, b, ref n);
		for (var i1 = 0; i1 < this.ScalarArray.Count; i1 ++) {
			n = this.ScalarArray[i1].Marshal(b, n);
		}
		Gobuf.WriteUvarint((ulong)this.ScalarMap.Count, b, ref n);
		foreach (var item1 in this.ScalarMap) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			if (item1.Value != null) {
				b[n++] = 1;
				n = item1.Value.Marshal(b, n);
			} else {
				b[n++] = 0;
			}
		}
		return n;
	}
	public int Unmarshal(byte[] b, int n) {
		this.Scalar = new Scalar();
		n = this.Scalar.Unmarshal(b, n);
		if (b[n++] != 0) {
			this.ScalarPtr = new Scalar();
			n = this.ScalarPtr.Unmarshal(b, n);
		} else {
			this.ScalarPtr = null;
		}
		{
			this.ScalarArray = new List<Scalar>((int)Gobuf.ReadUvarint(b, ref n));
			for (var i1 = 0; i1 < this.ScalarArray.Capacity; i1 ++) {
				Scalar v1;
				v1 = new Scalar();
				n = v1.Unmarshal(b, n);
				this.ScalarArray.Add(v1);
			}
		}
		{
			var cap1 = (int)Gobuf.ReadUvarint(b, ref n);
			this.ScalarMap = new Dictionary<long, Scalar>(cap1);
			for (var i1 = 0; i1 < cap1; i1 ++) {
				long key1;
				Scalar val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				if (b[n++] != 0) {
					val1 = new Scalar();
					n = val1.Unmarshal(b, n);
				} else {
					val1 = null;
				}
				this.ScalarMap.Add(key1, val1);
			}
		}
		return n;
	}
}
class Scalar {
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
