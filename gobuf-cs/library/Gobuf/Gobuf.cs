using System;
using System.Runtime.InteropServices;

class Gobuf 
{
	public static int ReadUint16(out ushort v, byte[] b, int offset) {
		v = (ushort)(
			(ushort)(b[offset++]) | 
			(ushort)(b[offset++]) << 8
		);
		return offset;
	}

	public static int ReadUint32(out uint v, byte[] b, int offset) {
		v = (uint)(b[offset++]) | 
			(uint)(b[offset++]) << 8  | 
			(uint)(b[offset++]) << 16 | 
			(uint)(b[offset++]) << 24;
		return offset;
	}

	public static int ReadUint64(out ulong v, byte[] b, int offset) {
		v = (ulong)(b[offset++]) | 
			(ulong)(b[offset++]) << 8  | 
			(ulong)(b[offset++]) << 16 | 
			(ulong)(b[offset++]) << 24 | 
			(ulong)(b[offset++]) << 32 | 
			(ulong)(b[offset++]) << 40 | 
			(ulong)(b[offset++]) << 48 | 
			(ulong)(b[offset++]) << 56;
		return offset;
	}


	public static int WriteUint16(ushort v, byte[] b, int offset) {
		b[offset++] = (byte)(v);
		b[offset++] = (byte)(v >> 8);
		return offset;
	}

	public static int WriteUint32(uint v, byte[] b, int offset) {
		b[offset++] = (byte)(v);
		b[offset++] = (byte)(v >> 8);
		b[offset++] = (byte)(v >> 16);
		b[offset++] = (byte)(v >> 24);
		return offset;
	}

	public static int WriteUint64(ulong v, byte[] b, int offset) {
		b[offset++] = (byte)(v);
		b[offset++] = (byte)(v >> 8);
		b[offset++] = (byte)(v >> 16);
		b[offset++] = (byte)(v >> 24);
		b[offset++] = (byte)(v >> 32);
		b[offset++] = (byte)(v >> 40);
		b[offset++] = (byte)(v >> 48);
		b[offset++] = (byte)(v >> 56);
		return offset;
	}

	#region Float32 and Float64

	public static int ReadFloat32(out float v, byte[] b, int offset) {
		uint v2;
		offset = ReadUint32(out v2, b, offset);
		v = Int32BitsToFloat(v2);
		return offset;
	}

	public static int WriteFloat32(float v, byte[] b, int offset) {
		return WriteUint32(FloatToInt32Bits(v), b, offset);
	}

	public static int ReadFloat64(out double v, byte[] b, int offset) {
		ulong v2;
		offset = ReadUint64(out v2, b, offset);
		v = BitConverter.Int64BitsToDouble((long)v2);
		return offset;
	}

	public static int WriteFloat64(double v, byte[] b, int offset) {
		return WriteUint64((ulong)BitConverter.DoubleToInt64Bits(v), b, offset);
	}

	private static uint FloatToInt32Bits(float f)
	{
		var bits = default(FloatUnion);
		bits.FloatData = f;
		return bits.IntData;
	}

	private static float Int32BitsToFloat(uint i)
	{
		var bits = default(FloatUnion);
		bits.IntData = i;
		return bits.FloatData;
	}

	[StructLayout(LayoutKind.Explicit)]
	private struct FloatUnion
	{
		[FieldOffset(0)]
		public uint IntData;
		[FieldOffset(0)]
		public float FloatData;
	}

	#endregion

	#region Varint and Uvarint

	public static int ReadUvarint(out ulong v, byte[] b, int offset) {
		int s = 0;
        int i = 0;
		v = 0;
		while (offset < b.Length) {
			var x = b[offset + i];
			if (x < 0x80) {
				if (i > 9 || i == 9 && x > 1) {
					throw new Exception("uvarint overflow");
				}
				v |= (ulong)(x) << s;
				break;
			}
			v |= (ulong)(x & 0x7f) << s;
			s += 7;
            i ++;
		}
		return offset;
	}

	public static int WriteUvarint(ulong v, byte[] b, int offset) {
		var i = 0;
		while (v >= 0x80) {
			b[i] = (byte)((byte)(v) | (byte)0x80);
			v >>= 7;
			i++;
		}
		b[i] = (byte)(v);
		return i + 1;
	}

	public static int UvarintSize(ulong x) {
		var i = 0;
		while (x >= 0x80) {
			x >>= 7;
			i++;
		}
		return i + 1;
	}

	public static int ReadVarint(out long v, byte[] b, int offset) {
		ulong uv;
		offset = ReadUvarint(out uv, b, offset);
		v = Zag(uv);
		return offset;
	}

	public static int WriteVarint(long v, byte[] b, int offset) {
		return WriteUvarint(Zig(v), b, offset);
	}

	public static int VarintSize(long v) {
		return UvarintSize(Zig(v));
	}

	private static long Zag(ulong uv) {
		if ((uv & 0x1) == 0x1) {
                return - ((long)(uv >> 1) + 1);
        }
		return (long)(uv >> 1);
	}

	private static ulong Zig(long v) {
		return (ulong)((v << 1) ^ (v >> 63));
	}

	#endregion

	#region Bytes and String

	public static int ReadBytes(out byte[] data, byte[] b, int offset) {
		ulong length;
		offset = ReadUvarint(out length, b, offset);
		data = new byte[length];
		Buffer.BlockCopy(b, offset, data, 0, (int)length);
		offset += (int)length;
		return offset;
	}

	public static int WriteBytes(byte[] data, byte[] b, int offset) {
		offset = WriteUvarint((ulong)data.Length, b, offset);
		Buffer.BlockCopy(data, 0, b, offset, data.Length);
		return offset + data.Length;
	}

	public static int ReadStrng(out string str, byte[] b, int offset) {
		byte[] data;
		offset = ReadBytes(out data, b, offset);
		str = System.Text.Encoding.UTF8.GetString(data);
		return offset;
	}

	public static int WriteString(string str, byte[] b, int offset) {
		return WriteBytes(System.Text.Encoding.UTF8.GetBytes(str), b, offset);
	}

	public static int StringSize(string str) {
		var count = System.Text.Encoding.UTF8.GetByteCount(str);
		return UvarintSize((ulong)count) + count;
	}

	#endregion
}