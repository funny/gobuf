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
		int size;
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
		Gobuf.WriteUint16((short)this.Int16, b, ref n);
		Gobuf.WriteUint16((ushort)this.Uint16, b, ref n);
		Gobuf.WriteUint32((int)this.Int32, b, ref n);
		Gobuf.WriteUint32((uint)this.Uint32, b, ref n);
		Gobuf.WriteUint64((long)this.Int64, b, ref n);
		Gobuf.WriteUint64((ulong)this.Uint64, b, ref n);
		Gobuf.WriteFloat32(this.Float32, b, ref n);
		Gobuf.WriteFloat64(this.Float64, b, ref n);
		Gobuf.WriteString(this.String, b, ref n);
		Gobuf.WriteBytes(this.Bytes, b, ref n);
		b[n++] = this.Bool ? 1 : 0;
		return n;
	}
	public int Unmarshal(byte[] b, int n) {
		this.Byte = (byte)b[n++];
		this.Byte = Gobuf.ReadUvarint(b, ref n);
		this.Int = Gobuf.ReadVarint(b, ref n);
		this.Int8 = (sbyte)b[n++];
		this.Int8 = Gobuf.ReadUvarint(b, ref n);
		this.Uint8 = (byte)b[n++];
		this.Uint8 = Gobuf.ReadUvarint(b, ref n);
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
			size += Gobuf.VarintSize(this.IntPtr);
		}
		size += 1;
		if (this.UintPtr != null) {
			size += Gobuf.UvarintSize(this.UintPtr);
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
		size += Gobuf.UvarintSize((ulong)this.IntArray.Length);
		for (var i1 = 0; i1 < this.IntArray.Length; i1 ++) {
			size += Gobuf.VarintSize(this.IntArray[i1]);
		}
		size += Gobuf.UvarintSize((ulong)this.UintArray.Length);
		for (var i1 = 0; i1 < this.UintArray.Length; i1 ++) {
			size += Gobuf.UvarintSize(this.UintArray[i1]);
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
			size += Gobuf.StringSize(this.StringArray[i1]);
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
			size += Gobuf.VarintSize(item1.Key);
			size += Gobuf.UvarintSize((ulong)item1.Value.Length);
			foreach (var item2 in item1.Value) {
				size += Gobuf.VarintSize(item2.Key);
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
		if (this.IntPtr != null) {
			b[n++] = 1;
			Gobuf.WriteVarint(this.IntPtr, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.UintPtr != null) {
			b[n++] = 1;
			Gobuf.WriteUvarint(this.UintPtr, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Int8Ptr != null) {
			b[n++] = 1;
			b[n++] = (byte)this.Int8Ptr;
		} else {
			b[n++] = 0;
		}
		if (this.Uint8Ptr != null) {
			b[n++] = 1;
			b[n++] = (byte)this.Uint8Ptr;
		} else {
			b[n++] = 0;
		}
		if (this.Int16Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteUint16((short)this.Int16Ptr, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Uint16Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteUint16((ushort)this.Uint16Ptr, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Int32Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteUint32((int)this.Int32Ptr, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Uint32Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteUint32((uint)this.Uint32Ptr, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Int64Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteUint64((long)this.Int64Ptr, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Uint64Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteUint64((ulong)this.Uint64Ptr, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Float32Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteFloat32(this.Float32Ptr, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.Float64Ptr != null) {
			b[n++] = 1;
			Gobuf.WriteFloat64(this.Float64Ptr, b, ref n);
		} else {
			b[n++] = 0;
		}
		if (this.StringPtr != null) {
			b[n++] = 1;
			Gobuf.WriteString(this.StringPtr, b, ref n);
		} else {
			b[n++] = 0;
		}
		Gobuf.WriteUvarint((ulong)this.IntArray.Length, b, ref n);
		for (var i1 = 0; i1 < this.IntArray.Length; i1 ++) {
			Gobuf.WriteVarint(this.IntArray[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.UintArray.Length, b, ref n);
		for (var i1 = 0; i1 < this.UintArray.Length; i1 ++) {
			Gobuf.WriteUvarint(this.UintArray[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Int8Array.Length, b, ref n);
		for (var i1 = 0; i1 < this.Int8Array.Length; i1 ++) {
			b[n++] = (byte)this.Int8Array[i1];
		}
		Gobuf.WriteBytes(this.Uint8Array, b, ref n);
		Gobuf.WriteUvarint((ulong)this.Int16Array.Length, b, ref n);
		for (var i1 = 0; i1 < this.Int16Array.Length; i1 ++) {
			Gobuf.WriteUint16((short)this.Int16Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Uint16Array.Length, b, ref n);
		for (var i1 = 0; i1 < this.Uint16Array.Length; i1 ++) {
			Gobuf.WriteUint16((ushort)this.Uint16Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Int32Array.Length, b, ref n);
		for (var i1 = 0; i1 < this.Int32Array.Length; i1 ++) {
			Gobuf.WriteUint32((int)this.Int32Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Uint32Array.Length, b, ref n);
		for (var i1 = 0; i1 < this.Uint32Array.Length; i1 ++) {
			Gobuf.WriteUint32((uint)this.Uint32Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Int64Array.Length, b, ref n);
		for (var i1 = 0; i1 < this.Int64Array.Length; i1 ++) {
			Gobuf.WriteUint64((long)this.Int64Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Uint64Array.Length, b, ref n);
		for (var i1 = 0; i1 < this.Uint64Array.Length; i1 ++) {
			Gobuf.WriteUint64((ulong)this.Uint64Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Float32Array.Length, b, ref n);
		for (var i1 = 0; i1 < this.Float32Array.Length; i1 ++) {
			Gobuf.WriteFloat32(this.Float32Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.Float64Array.Length, b, ref n);
		for (var i1 = 0; i1 < this.Float64Array.Length; i1 ++) {
			Gobuf.WriteFloat64(this.Float64Array[i1], b, ref n);
		}
		Gobuf.WriteUvarint((ulong)this.StringArray.Length, b, ref n);
		for (var i1 = 0; i1 < this.StringArray.Length; i1 ++) {
			Gobuf.WriteString(this.StringArray[i1], b, ref n);
		}
		n = this.Message.Marshal(b, n);
		if (this.MessagePtr != null) {
			b[n++] = 1;
			n = this.MessagePtr.Marshal(b, n);
		} else {
			b[n++] = 0;
		}
		Gobuf.WriteUvarint((ulong)this.MessageArray.Length, b, ref n);
		for (var i1 = 0; i1 < this.MessageArray.Length; i1 ++) {
			n = this.MessageArray[i1].Marshal(b, n);
		}
		Gobuf.WriteUvarint((ulong)this.MessagePtrArray.Length, b, ref n);
		for (var i1 = 0; i1 < this.MessagePtrArray.Length; i1 ++) {
			if (this.MessagePtrArray[i1] != null) {
				b[n++] = 1;
				n = this.MessagePtrArray[i1].Marshal(b, n);
			} else {
				b[n++] = 0;
			}
		}
		Gobuf.WriteUvarint((ulong)this.MessageArrayArray.Length, b, ref n);
		for (var i1 = 0; i1 < this.MessageArrayArray.Length; i1 ++) {
			Gobuf.WriteUvarint((ulong)this.MessageArrayArray[i1].Length, b, ref n);
			for (var i2 = 0; i2 < this.MessageArrayArray[i1].Length; i2 ++) {
				n = this.MessageArrayArray[i1][i2].Marshal(b, n);
			}
		}
		Gobuf.WriteUvarint((ulong)this.IntMap.Length, b, ref n);
		foreach (var item1 in this.IntMap) {
			Gobuf.WriteVarint(item1.Key, b, ref n);
			Gobuf.WriteUvarint((ulong)item1.Value.Length, b, ref n);
			foreach (var item2 in item1.Value) {
				Gobuf.WriteVarint(item2.Key, b, ref n);
				Gobuf.WriteUvarint((ulong)item2.Value.Length, b, ref n);
				for (var i3 = 0; i3 < item2.Value.Length; i3 ++) {
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
			var val1 = new long();
			val1 = Gobuf.ReadVarint(b, ref n);
			this.IntPtr = val1;
		}
		if (b[n++] != 0) {
			var val1 = new ulong();
			this.UintPtr = val1;
		}
		if (b[n++] != 0) {
			var val1 = new sbyte();
			val1 = (sbyte)b[n++];
			val1 = Gobuf.ReadUvarint(b, ref n);
			this.Int8Ptr = val1;
		}
		if (b[n++] != 0) {
			var val1 = new byte();
			val1 = (byte)b[n++];
			val1 = Gobuf.ReadUvarint(b, ref n);
			this.Uint8Ptr = val1;
		}
		if (b[n++] != 0) {
			var val1 = new short();
			val1 = (short)Gobuf.ReadUint16(b, ref n);
			this.Int16Ptr = val1;
		}
		if (b[n++] != 0) {
			var val1 = new ushort();
			val1 = (ushort)Gobuf.ReadUint16(b, ref n);
			this.Uint16Ptr = val1;
		}
		if (b[n++] != 0) {
			var val1 = new int();
			val1 = (int)Gobuf.ReadUint32(b, ref n);
			this.Int32Ptr = val1;
		}
		if (b[n++] != 0) {
			var val1 = new uint();
			val1 = (uint)Gobuf.ReadUint32(b, ref n);
			this.Uint32Ptr = val1;
		}
		if (b[n++] != 0) {
			var val1 = new long();
			val1 = (long)Gobuf.ReadUint64(b, ref n);
			this.Int64Ptr = val1;
		}
		if (b[n++] != 0) {
			var val1 = new ulong();
			val1 = (ulong)Gobuf.ReadUint64(b, ref n);
			this.Uint64Ptr = val1;
		}
		if (b[n++] != 0) {
			var val1 = new float();
			val1 = Gobuf.ReadFloat32(b, ref n);
			this.Float32Ptr = val1;
		}
		if (b[n++] != 0) {
			var val1 = new double();
			val1 = Gobuf.ReadFloat64(b, ref n);
			this.Float64Ptr = val1;
		}
		if (b[n++] != 0) {
			var val1 = new string();
			val1 = Gobuf.ReadString(b, ref n);
			this.StringPtr = val1;
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.IntArray = new System.Collections.Generic.List<long>();
			for (var i1 = 0; i1 < l; i1 ++) {
				this.IntArray[i1] = Gobuf.ReadVarint(b, ref n);
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.UintArray = new System.Collections.Generic.List<ulong>();
			for (var i1 = 0; i1 < l; i1 ++) {
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.Int8Array = new System.Collections.Generic.List<sbyte>();
			for (var i1 = 0; i1 < l; i1 ++) {
				this.Int8Array[i1] = (sbyte)b[n++];
				this.Int8Array[i1] = Gobuf.ReadUvarint(b, ref n);
			}
		}
		this.Uint8Array = Gobuf.ReadBytes(b, ref n);
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.Int16Array = new System.Collections.Generic.List<short>();
			for (var i1 = 0; i1 < l; i1 ++) {
				this.Int16Array[i1] = (short)Gobuf.ReadUint16(b, ref n);
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.Uint16Array = new System.Collections.Generic.List<ushort>();
			for (var i1 = 0; i1 < l; i1 ++) {
				this.Uint16Array[i1] = (ushort)Gobuf.ReadUint16(b, ref n);
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.Int32Array = new System.Collections.Generic.List<int>();
			for (var i1 = 0; i1 < l; i1 ++) {
				this.Int32Array[i1] = (int)Gobuf.ReadUint32(b, ref n);
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.Uint32Array = new System.Collections.Generic.List<uint>();
			for (var i1 = 0; i1 < l; i1 ++) {
				this.Uint32Array[i1] = (uint)Gobuf.ReadUint32(b, ref n);
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.Int64Array = new System.Collections.Generic.List<long>();
			for (var i1 = 0; i1 < l; i1 ++) {
				this.Int64Array[i1] = (long)Gobuf.ReadUint64(b, ref n);
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.Uint64Array = new System.Collections.Generic.List<ulong>();
			for (var i1 = 0; i1 < l; i1 ++) {
				this.Uint64Array[i1] = (ulong)Gobuf.ReadUint64(b, ref n);
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.Float32Array = new System.Collections.Generic.List<float>();
			for (var i1 = 0; i1 < l; i1 ++) {
				this.Float32Array[i1] = Gobuf.ReadFloat32(b, ref n);
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.Float64Array = new System.Collections.Generic.List<double>();
			for (var i1 = 0; i1 < l; i1 ++) {
				this.Float64Array[i1] = Gobuf.ReadFloat64(b, ref n);
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.StringArray = new System.Collections.Generic.List<string>();
			for (var i1 = 0; i1 < l; i1 ++) {
				this.StringArray[i1] = Gobuf.ReadString(b, ref n);
			}
		}
		n = this.Message.Unmarshal(b, n);
		if (b[n++] != 0) {
			var val1 = new ScalarTypes();
			n = val1.Unmarshal(b, n);
			this.MessagePtr = val1;
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.MessageArray = new System.Collections.Generic.List<ScalarTypes>();
			for (var i1 = 0; i1 < l; i1 ++) {
				n = this.MessageArray[i1].Unmarshal(b, n);
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.MessagePtrArray = new System.Collections.Generic.List<ScalarTypes>();
			for (var i1 = 0; i1 < l; i1 ++) {
				if (b[n++] != 0) {
					var val2 = new ScalarTypes();
					n = val2.Unmarshal(b, n);
					this.MessagePtrArray[i1] = val2;
				}
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.MessageArrayArray = new System.Collections.Generic.List<ScalarTypes[]>();
			for (var i1 = 0; i1 < l; i1 ++) {
				{
					ulong l;
					Gobuf.ReadUvarint(out l, b, ref n);
					this.MessageArrayArray[i1] = new System.Collections.Generic.List<ScalarTypes>();
					for (var i2 = 0; i2 < l; i2 ++) {
						n = this.MessageArrayArray[i1][i2].Unmarshal(b, n);
					}
				}
			}
		}
		{
			ulong l;
			Gobuf.ReadUvarint(out l, b, ref n);
			this.IntMap = System.Collections.Generic.Dictionary<long, Dictionary<long, ScalarTypes[]>>();
			for (var i1 = 0; i1 < l; i1 ++) {
				long key1;
				Dictionary<long, ScalarTypes[]> val1;
				key1 = Gobuf.ReadVarint(b, ref n);
				{
					ulong l;
					Gobuf.ReadUvarint(out l, b, ref n);
					val1 = System.Collections.Generic.Dictionary<long, ScalarTypes[]>();
					for (var i2 = 0; i2 < l; i2 ++) {
						long key2;
						ScalarTypes[] val2;
						key2 = Gobuf.ReadVarint(b, ref n);
						{
							ulong l;
							Gobuf.ReadUvarint(out l, b, ref n);
							val2 = new System.Collections.Generic.List<ScalarTypes>();
							for (var i3 = 0; i3 < l; i3 ++) {
								if (b[n++] != 0) {
									var val4 = new ScalarTypes();
									n = val4.Unmarshal(b, n);
									val2[i3] = val4;
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
