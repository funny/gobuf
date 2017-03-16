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
	public int Size() {
		int size;
		size += 1;
		size += Gobuf.VarintSize((long)this.Int);
		size += Gobuf.UvarintSize((ulong)this.Uint);
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
		size += Gobuf.UvarintSize((ulong)this.String.Length) + this.String.Length;
		size += Gobuf.UvarintSize((ulong)this.Bytes.Length) + this.Bytes.Length;
		return size;
	}
	public int Marshal(byte[] b, int n) {
		b[n] = (byte)msg.Byte;
		n += 1;
		n = Gobuf.WriteVarint((long)msg.Int, b, n);
		n = Gobuf.WriteUvarint((ulong)msg.Uint, b, n);
		b[n] = (byte)msg.Int8;
		n += 1;
		b[n] = (byte)msg.Uint8;
		n += 1;
		n = Gobuf.WriteUint16((ushort)msg.Int16, b, n);
		n = Gobuf.WriteUint16((ushort)msg.Uint16, b, n);
		n = Gobuf.WriteUint32((long)msg.Int32, b, n);
		n = Gobuf.WriteUint32((long)msg.Uint32, b, n);
		n = Gobuf.WriteUint64((ulong)msg.Int64, b, n);
		n = Gobuf.WriteUint64((ulong)msg.Uint64, b, n);
		n = Gobuf.WriteFloat32((float)msg.Float32, b, n);
		n = Gobuf.WriteFloat64((double)msg.Float64, b, n);
		n = Gobuf.WriteString(msg.String, b, n);
		n = Gobuf.WriteBytes(msg.Bytes, b, n);
		return n;
	}
	public int Unmarshal(byte[] b, int n) {
		msg.Byte = (byte)b[n];
		n += 1;
		n = (long)Gobuf.ReadVarint(out msg.Int, b, n);
		n = (ulong)Gobuf.ReadUvarint(out msg.Uint, b, n);
		msg.Int8 = (sbyte)b[n];
		n += 1;
		msg.Uint8 = (byte)b[n];
		n += 1;
		n = (short)Gobuf.ReadUint16(out msg.Int16, b, n);
		n = (ushort)Gobuf.ReadUint16(out msg.Uint16, b, n);
		n = (int)Gobuf.ReadUint32(out msg.Int32, b, n);
		n = (uint)Gobuf.ReadUint32(out msg.Uint32, b, n);
		n = (long)Gobuf.ReadUint64(out msg.Int64, b, n);
		n = (ulong)Gobuf.ReadUint64(out msg.Uint64, b, n);
		n = (float)Gobuf.ReadFloat32(out msg.Float32, b, n);
		n = (double)Gobuf.ReadFloat64(out msg.Float64, b, n);
		n = (string)Gobuf.ReadString(out msg.String, b, n);
		n = (byte[])Gobuf.ReadBytes(out msg.Bytes, b, n);
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
	public Nullable<string> StringPtr;
	public long[] IntArray;
	public ulong[] UintArray;
	public sbyte[] Int8Array;
	public byte[] Uint8Array;
	public short[] Int16Array;
	public ushort[] Uint16Array;
	public int[] Int32Array;
	public uint[] Uint32Array;
	public long[] Int64Array;
	public ulong[] Uint64Array;
	public float[] Float32Array;
	public double[] Float64Array;
	public string[] StringArray;
	public ScalarTypes Message;
	public ScalarTypes MessagePtr;
	public ScalarTypes[] MessageArray;
	public ScalarTypes[] MessagePtrArray;
	public ScalarTypes[][] MessageArrayArray;
	public Dictionary<long, Dictionary<long, ScalarTypes[]>> IntMap;
	public int Size() {
		int size;
		size += 1;
		if (this.IntPtr != null) {
			size += Gobuf.VarintSize((long)this.IntPtr);
		}
		size += 1;
		if (this.UintPtr != null) {
			size += Gobuf.UvarintSize((ulong)this.UintPtr);
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
			size += Gobuf.UvarintSize((ulong)this.StringPtr.Length) + this.StringPtr.Length;
		}
		size += Gobuf.UvarintSize((ulong)this.IntArray.Length);
		for (var i1 = 0; i1 < this.IntArray.Length; i1 ++) {
			size += Gobuf.VarintSize((long)this.IntArray[i1]);
		}
		size += Gobuf.UvarintSize((ulong)this.UintArray.Length);
		for (var i1 = 0; i1 < this.UintArray.Length; i1 ++) {
			size += Gobuf.UvarintSize((ulong)this.UintArray[i1]);
		}
		size += Gobuf.UvarintSize((ulong)this.Int8Array.Length) + this.Int8Array.Length * 1;
		size += Gobuf.UvarintSize((ulong)this.Uint8Array.Length) + this.Uint8Array.Length;
		size += Gobuf.UvarintSize((ulong)this.Int16Array.Length) + this.Int16Array.Length * 2;
		size += Gobuf.UvarintSize((ulong)this.Uint16Array.Length) + this.Uint16Array.Length * 2;
		size += Gobuf.UvarintSize((ulong)this.Int32Array.Length) + this.Int32Array.Length * 4;
		size += Gobuf.UvarintSize((ulong)this.Uint32Array.Length) + this.Uint32Array.Length * 4;
		size += Gobuf.UvarintSize((ulong)this.Int64Array.Length) + this.Int64Array.Length * 8;
		size += Gobuf.UvarintSize((ulong)this.Uint64Array.Length) + this.Uint64Array.Length * 8;
		size += Gobuf.UvarintSize((ulong)this.Float32Array.Length) + this.Float32Array.Length * 4;
		size += Gobuf.UvarintSize((ulong)this.Float64Array.Length) + this.Float64Array.Length * 8;
		size += Gobuf.UvarintSize((ulong)this.StringArray.Length);
		for (var i1 = 0; i1 < this.StringArray.Length; i1 ++) {
			size += Gobuf.UvarintSize((ulong)this.StringArray[i1].Length) + this.StringArray[i1].Length;
		}
		size += this.Message.Size();
		size += 1;
		if (this.MessagePtr != null) {
			size += this.MessagePtr.Size();
		}
		size += Gobuf.UvarintSize((ulong)this.MessageArray.Length);
		for (var i1 = 0; i1 < this.MessageArray.Length; i1 ++) {
			size += this.MessageArray[i1].Size();
		}
		size += Gobuf.UvarintSize((ulong)this.MessagePtrArray.Length);
		for (var i1 = 0; i1 < this.MessagePtrArray.Length; i1 ++) {
			size += 1;
			if (this.MessagePtrArray[i1] != null) {
				size += this.MessagePtrArray[i1].Size();
			}
		}
		size += Gobuf.UvarintSize((ulong)this.MessageArrayArray.Length);
		for (var i1 = 0; i1 < this.MessageArrayArray.Length; i1 ++) {
			size += Gobuf.UvarintSize((ulong)this.MessageArrayArray[i1].Length);
			for (var i2 = 0; i2 < this.MessageArrayArray[i1].Length; i2 ++) {
				size += this.MessageArrayArray[i1][i2].Size();
			}
		}
		size += Gobuf.UvarintSize((ulong)this.IntMap.Length);
		foreach (var item1 in this.IntMap) {
			size += Gobuf.VarintSize((long)item1.Key);
			size += Gobuf.UvarintSize((ulong)item1.Value.Length);
			foreach (var item2 in item1.Value) {
				size += Gobuf.VarintSize((long)item2.Key);
				size += Gobuf.UvarintSize((ulong)item2.Value.Length);
				for (var i3 = 0; i3 < item2.Value.Length; i3 ++) {
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
		if (msg.IntPtr != null) {
			b[n] = 1;
			n ++;
			n = Gobuf.WriteVarint((long)msg.IntPtr, b, n);
		} else {
			b[n] = 0;
			n ++;
		}
		if (msg.UintPtr != null) {
			b[n] = 1;
			n ++;
			n = Gobuf.WriteUvarint((ulong)msg.UintPtr, b, n);
		} else {
			b[n] = 0;
			n ++;
		}
		if (msg.Int8Ptr != null) {
			b[n] = 1;
			n ++;
			b[n] = (byte)msg.Int8Ptr;
			n += 1;
		} else {
			b[n] = 0;
			n ++;
		}
		if (msg.Uint8Ptr != null) {
			b[n] = 1;
			n ++;
			b[n] = (byte)msg.Uint8Ptr;
			n += 1;
		} else {
			b[n] = 0;
			n ++;
		}
		if (msg.Int16Ptr != null) {
			b[n] = 1;
			n ++;
			n = Gobuf.WriteUint16((ushort)msg.Int16Ptr, b, n);
		} else {
			b[n] = 0;
			n ++;
		}
		if (msg.Uint16Ptr != null) {
			b[n] = 1;
			n ++;
			n = Gobuf.WriteUint16((ushort)msg.Uint16Ptr, b, n);
		} else {
			b[n] = 0;
			n ++;
		}
		if (msg.Int32Ptr != null) {
			b[n] = 1;
			n ++;
			n = Gobuf.WriteUint32((long)msg.Int32Ptr, b, n);
		} else {
			b[n] = 0;
			n ++;
		}
		if (msg.Uint32Ptr != null) {
			b[n] = 1;
			n ++;
			n = Gobuf.WriteUint32((long)msg.Uint32Ptr, b, n);
		} else {
			b[n] = 0;
			n ++;
		}
		if (msg.Int64Ptr != null) {
			b[n] = 1;
			n ++;
			n = Gobuf.WriteUint64((ulong)msg.Int64Ptr, b, n);
		} else {
			b[n] = 0;
			n ++;
		}
		if (msg.Uint64Ptr != null) {
			b[n] = 1;
			n ++;
			n = Gobuf.WriteUint64((ulong)msg.Uint64Ptr, b, n);
		} else {
			b[n] = 0;
			n ++;
		}
		if (msg.Float32Ptr != null) {
			b[n] = 1;
			n ++;
			n = Gobuf.WriteFloat32((float)msg.Float32Ptr, b, n);
		} else {
			b[n] = 0;
			n ++;
		}
		if (msg.Float64Ptr != null) {
			b[n] = 1;
			n ++;
			n = Gobuf.WriteFloat64((double)msg.Float64Ptr, b, n);
		} else {
			b[n] = 0;
			n ++;
		}
		if (msg.StringPtr != null) {
			b[n] = 1;
			n ++;
			n = Gobuf.WriteString(msg.StringPtr, b, n);
		} else {
			b[n] = 0;
			n ++;
		}
		n = Gobuf.WriteUvarint((ulong)msg.IntArray.Length, b, n);
		for (var i1 = 0; i1 < msg.IntArray.Length; i1 ++) {
			n = Gobuf.WriteVarint((long)msg.IntArray[i1], b, n);
		}
		n = Gobuf.WriteUvarint((ulong)msg.UintArray.Length, b, n);
		for (var i1 = 0; i1 < msg.UintArray.Length; i1 ++) {
			n = Gobuf.WriteUvarint((ulong)msg.UintArray[i1], b, n);
		}
		n = Gobuf.WriteUvarint((ulong)msg.Int8Array.Length, b, n);
		for (var i1 = 0; i1 < msg.Int8Array.Length; i1 ++) {
			b[n] = (byte)msg.Int8Array[i1];
			n += 1;
		}
		n = Gobuf.WriteBytes(msg.Uint8Array, b, n);
		n = Gobuf.WriteUvarint((ulong)msg.Int16Array.Length, b, n);
		for (var i1 = 0; i1 < msg.Int16Array.Length; i1 ++) {
			n = Gobuf.WriteUint16((ushort)msg.Int16Array[i1], b, n);
		}
		n = Gobuf.WriteUvarint((ulong)msg.Uint16Array.Length, b, n);
		for (var i1 = 0; i1 < msg.Uint16Array.Length; i1 ++) {
			n = Gobuf.WriteUint16((ushort)msg.Uint16Array[i1], b, n);
		}
		n = Gobuf.WriteUvarint((ulong)msg.Int32Array.Length, b, n);
		for (var i1 = 0; i1 < msg.Int32Array.Length; i1 ++) {
			n = Gobuf.WriteUint32((long)msg.Int32Array[i1], b, n);
		}
		n = Gobuf.WriteUvarint((ulong)msg.Uint32Array.Length, b, n);
		for (var i1 = 0; i1 < msg.Uint32Array.Length; i1 ++) {
			n = Gobuf.WriteUint32((long)msg.Uint32Array[i1], b, n);
		}
		n = Gobuf.WriteUvarint((ulong)msg.Int64Array.Length, b, n);
		for (var i1 = 0; i1 < msg.Int64Array.Length; i1 ++) {
			n = Gobuf.WriteUint64((ulong)msg.Int64Array[i1], b, n);
		}
		n = Gobuf.WriteUvarint((ulong)msg.Uint64Array.Length, b, n);
		for (var i1 = 0; i1 < msg.Uint64Array.Length; i1 ++) {
			n = Gobuf.WriteUint64((ulong)msg.Uint64Array[i1], b, n);
		}
		n = Gobuf.WriteUvarint((ulong)msg.Float32Array.Length, b, n);
		for (var i1 = 0; i1 < msg.Float32Array.Length; i1 ++) {
			n = Gobuf.WriteFloat32((float)msg.Float32Array[i1], b, n);
		}
		n = Gobuf.WriteUvarint((ulong)msg.Float64Array.Length, b, n);
		for (var i1 = 0; i1 < msg.Float64Array.Length; i1 ++) {
			n = Gobuf.WriteFloat64((double)msg.Float64Array[i1], b, n);
		}
		n = Gobuf.WriteUvarint((ulong)msg.StringArray.Length, b, n);
		for (var i1 = 0; i1 < msg.StringArray.Length; i1 ++) {
			n = Gobuf.WriteString(msg.StringArray[i1], b, n);
		}
		n = msg.Message.Marshal(b, n);
		if (msg.MessagePtr != null) {
			b[n] = 1;
			n ++;
			n = msg.MessagePtr.Marshal(b, n);
		} else {
			b[n] = 0;
			n ++;
		}
		n = Gobuf.WriteUvarint((ulong)msg.MessageArray.Length, b, n);
		for (var i1 = 0; i1 < msg.MessageArray.Length; i1 ++) {
			n = msg.MessageArray[i1].Marshal(b, n);
		}
		n = Gobuf.WriteUvarint((ulong)msg.MessagePtrArray.Length, b, n);
		for (var i1 = 0; i1 < msg.MessagePtrArray.Length; i1 ++) {
			if (msg.MessagePtrArray[i1] != null) {
				b[n] = 1;
				n ++;
				n = msg.MessagePtrArray[i1].Marshal(b, n);
			} else {
				b[n] = 0;
				n ++;
			}
		}
		n = Gobuf.WriteUvarint((ulong)msg.MessageArrayArray.Length, b, n);
		for (var i1 = 0; i1 < msg.MessageArrayArray.Length; i1 ++) {
			n = Gobuf.WriteUvarint((ulong)msg.MessageArrayArray[i1].Length, b, n);
			for (var i2 = 0; i2 < msg.MessageArrayArray[i1].Length; i2 ++) {
				n = msg.MessageArrayArray[i1][i2].Marshal(b, n);
			}
		}
		n = Gobuf.WriteUvarint((ulong)msg.IntMap.Length, b, n);
		foreach (var item1 in msg.IntMap) {
			n = Gobuf.WriteVarint((long)item1.Key, b, n);
			n = Gobuf.WriteUvarint((ulong)item1.Value.Length, b, n);
			foreach (var item2 in item1.Value) {
				n = Gobuf.WriteVarint((long)item2.Key, b, n);
				n = Gobuf.WriteUvarint((ulong)item2.Value.Length, b, n);
				for (var i3 = 0; i3 < item2.Value.Length; i3 ++) {
					if (item2.Value[i3] != null) {
						b[n] = 1;
						n ++;
						n = item2.Value[i3].Marshal(b, n);
					} else {
						b[n] = 0;
						n ++;
					}
				}
			}
		}
		return n;
	}
	public int Unmarshal(byte[] b, int n) {
		if (b[n] != 0) {
			n += 1;
			var val1 = new long();
			n = (long)Gobuf.ReadVarint(out val1, b, n);
			msg.IntPtr = val1;
		} else {
			n += 1;
		}
		if (b[n] != 0) {
			n += 1;
			var val1 = new ulong();
			n = (ulong)Gobuf.ReadUvarint(out val1, b, n);
			msg.UintPtr = val1;
		} else {
			n += 1;
		}
		if (b[n] != 0) {
			n += 1;
			var val1 = new sbyte();
			val1 = (sbyte)b[n];
			n += 1;
			msg.Int8Ptr = val1;
		} else {
			n += 1;
		}
		if (b[n] != 0) {
			n += 1;
			var val1 = new byte();
			val1 = (byte)b[n];
			n += 1;
			msg.Uint8Ptr = val1;
		} else {
			n += 1;
		}
		if (b[n] != 0) {
			n += 1;
			var val1 = new short();
			n = (short)Gobuf.ReadUint16(out val1, b, n);
			msg.Int16Ptr = val1;
		} else {
			n += 1;
		}
		if (b[n] != 0) {
			n += 1;
			var val1 = new ushort();
			n = (ushort)Gobuf.ReadUint16(out val1, b, n);
			msg.Uint16Ptr = val1;
		} else {
			n += 1;
		}
		if (b[n] != 0) {
			n += 1;
			var val1 = new int();
			n = (int)Gobuf.ReadUint32(out val1, b, n);
			msg.Int32Ptr = val1;
		} else {
			n += 1;
		}
		if (b[n] != 0) {
			n += 1;
			var val1 = new uint();
			n = (uint)Gobuf.ReadUint32(out val1, b, n);
			msg.Uint32Ptr = val1;
		} else {
			n += 1;
		}
		if (b[n] != 0) {
			n += 1;
			var val1 = new long();
			n = (long)Gobuf.ReadUint64(out val1, b, n);
			msg.Int64Ptr = val1;
		} else {
			n += 1;
		}
		if (b[n] != 0) {
			n += 1;
			var val1 = new ulong();
			n = (ulong)Gobuf.ReadUint64(out val1, b, n);
			msg.Uint64Ptr = val1;
		} else {
			n += 1;
		}
		if (b[n] != 0) {
			n += 1;
			var val1 = new float();
			n = (float)Gobuf.ReadFloat32(out val1, b, n);
			msg.Float32Ptr = val1;
		} else {
			n += 1;
		}
		if (b[n] != 0) {
			n += 1;
			var val1 = new double();
			n = (double)Gobuf.ReadFloat64(out val1, b, n);
			msg.Float64Ptr = val1;
		} else {
			n += 1;
		}
		if (b[n] != 0) {
			n += 1;
			var val1 = new string();
			n = (string)Gobuf.ReadString(out val1, b, n);
			msg.StringPtr = val1;
		} else {
			n += 1;
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.IntArray = new System.Collections.Generic.List<long>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = (long)Gobuf.ReadVarint(out msg.IntArray[i1], b, n);
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.UintArray = new System.Collections.Generic.List<ulong>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = (ulong)Gobuf.ReadUvarint(out msg.UintArray[i1], b, n);
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.Int8Array = new System.Collections.Generic.List<sbyte>();
			for (var i1 = 0; i1 < l; i1 ++) {
				msg.Int8Array[i1] = (sbyte)b[n];
				n += 1;
			}
		}
		n = (byte[])Gobuf.ReadBytes(out msg.Uint8Array, b, n);
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.Int16Array = new System.Collections.Generic.List<short>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = (short)Gobuf.ReadUint16(out msg.Int16Array[i1], b, n);
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.Uint16Array = new System.Collections.Generic.List<ushort>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = (ushort)Gobuf.ReadUint16(out msg.Uint16Array[i1], b, n);
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.Int32Array = new System.Collections.Generic.List<int>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = (int)Gobuf.ReadUint32(out msg.Int32Array[i1], b, n);
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.Uint32Array = new System.Collections.Generic.List<uint>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = (uint)Gobuf.ReadUint32(out msg.Uint32Array[i1], b, n);
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.Int64Array = new System.Collections.Generic.List<long>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = (long)Gobuf.ReadUint64(out msg.Int64Array[i1], b, n);
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.Uint64Array = new System.Collections.Generic.List<ulong>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = (ulong)Gobuf.ReadUint64(out msg.Uint64Array[i1], b, n);
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.Float32Array = new System.Collections.Generic.List<float>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = (float)Gobuf.ReadFloat32(out msg.Float32Array[i1], b, n);
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.Float64Array = new System.Collections.Generic.List<double>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = (double)Gobuf.ReadFloat64(out msg.Float64Array[i1], b, n);
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.StringArray = new System.Collections.Generic.List<string>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = (string)Gobuf.ReadString(out msg.StringArray[i1], b, n);
			}
		}
		n = msg.Message.Unmarshal(b, n);
		if (b[n] != 0) {
			n += 1;
			var val1 = new ScalarTypes();
			n = val1.Unmarshal(b, n);
			msg.MessagePtr = val1;
		} else {
			n += 1;
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.MessageArray = new System.Collections.Generic.List<ScalarTypes>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = msg.MessageArray[i1].Unmarshal(b, n);
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.MessagePtrArray = new System.Collections.Generic.List<ScalarTypes>();
			for (var i1 = 0; i1 < l; i1 ++) {
				if (b[n] != 0) {
					n += 1;
					var val2 = new ScalarTypes();
					n = val2.Unmarshal(b, n);
					msg.MessagePtrArray[i1] = val2;
				} else {
					n += 1;
				}
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.MessageArrayArray = new System.Collections.Generic.List<ScalarTypes[]>();
			for (var i1 = 0; i1 < l; i1 ++) {
				{
					ulong l;
					n = Gobuf.ReadUvarint(out l, b, n);
					msg.MessageArrayArray[i1] = new System.Collections.Generic.List<ScalarTypes>();
					for (var i2 = 0; i2 < l; i2 ++) {
						n = msg.MessageArrayArray[i1][i2].Unmarshal(b, n);
					}
				}
			}
		}
		{
			ulong l;
			n = Gobuf.ReadUvarint(out l, b, n);
			msg.IntMap = System.Collections.Generic.Dictionary<long, Dictionary<long, ScalarTypes[]>>();
			for (var i1 = 0; i1 < l; i1 ++) {
				long key1;
				Dictionary<long, ScalarTypes[]> val1;
				n = (long)Gobuf.ReadVarint(out key1, b, n);
				{
					ulong l;
					n = Gobuf.ReadUvarint(out l, b, n);
					val1 = System.Collections.Generic.Dictionary<long, ScalarTypes[]>();
					for (var i2 = 0; i2 < l; i2 ++) {
						long key2;
						ScalarTypes[] val2;
						n = (long)Gobuf.ReadVarint(out key2, b, n);
						{
							ulong l;
							n = Gobuf.ReadUvarint(out l, b, n);
							val2 = new System.Collections.Generic.List<ScalarTypes>();
							for (var i3 = 0; i3 < l; i3 ++) {
								if (b[n] != 0) {
									n += 1;
									var val4 = new ScalarTypes();
									n = val4.Unmarshal(b, n);
									val2[i3] = val4;
								} else {
									n += 1;
								}
							}
						}
						val1[key2] = val2;
					}
				}
				msg.IntMap[key1] = val1;
			}
		}
		return n;
	}
}
