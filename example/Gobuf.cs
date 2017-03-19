using System;
using System.Runtime.InteropServices;

class Gobuf 
{
	public static ushort ReadUint16(byte[] b, ref int offset) {
		return (ushort)(
			(ushort)(b[offset++]) | 
			(ushort)(b[offset++]) << 8
		);
	}

	public static uint ReadUint32(byte[] b, ref int offset) {
		return (uint)(b[offset++]) | 
			(uint)(b[offset++]) << 8  | 
			(uint)(b[offset++]) << 16 | 
			(uint)(b[offset++]) << 24;
	}

	public static ulong ReadUint64(byte[] b, ref int offset) {
		return (ulong)(b[offset++]) | 
			(ulong)(b[offset++]) << 8  | 
			(ulong)(b[offset++]) << 16 | 
			(ulong)(b[offset++]) << 24 | 
			(ulong)(b[offset++]) << 32 | 
			(ulong)(b[offset++]) << 40 | 
			(ulong)(b[offset++]) << 48 | 
			(ulong)(b[offset++]) << 56;
	}

	public static float ReadFloat32(byte[] b, ref int offset) {
		return Int32BitsToFloat(ReadUint32(b, ref offset));
	}

	public static double ReadFloat64(byte[] b, ref int offset) {
		return BitConverter.Int64BitsToDouble((long)ReadUint64(b, ref offset));
	}

	public static byte[] ReadBytes(byte[] b, ref int offset) {
		var length = (int)ReadUvarint(b, ref offset);
		var data = new byte[length];
		Buffer.BlockCopy(b, offset, data, 0, length);
		offset += length;
		return data;
	}

	public static string ReadString(byte[] b, ref int offset) {
		return System.Text.Encoding.UTF8.GetString(ReadBytes(b, ref offset));
	}



	public static void WriteUint16(ushort v, byte[] b, ref int offset) {
		b[offset++] = (byte)(v);
		b[offset++] = (byte)(v >> 8);
	}

	public static void WriteUint32(uint v, byte[] b, ref int offset) {
		b[offset++] = (byte)(v);
		b[offset++] = (byte)(v >> 8);
		b[offset++] = (byte)(v >> 16);
		b[offset++] = (byte)(v >> 24);
	}

	public static void WriteUint64(ulong v, byte[] b, ref int offset) {
		b[offset++] = (byte)(v);
		b[offset++] = (byte)(v >> 8);
		b[offset++] = (byte)(v >> 16);
		b[offset++] = (byte)(v >> 24);
		b[offset++] = (byte)(v >> 32);
		b[offset++] = (byte)(v >> 40);
		b[offset++] = (byte)(v >> 48);
		b[offset++] = (byte)(v >> 56);
	}

	public static void WriteFloat32(float v, byte[] b, ref int offset) {
		WriteUint32((uint)FloatToInt32Bits(v), b, ref offset);
	}

	public static void WriteFloat64(double v, byte[] b, ref int offset) {
		WriteUint64((ulong)BitConverter.DoubleToInt64Bits(v), b, ref offset);
	}

	public static void WriteBytes(byte[] data, byte[] b, ref int offset) {
		WriteUvarint((ulong)data.Length, b, ref offset);
		Buffer.BlockCopy(data, 0, b, offset, data.Length);
		offset += data.Length;
	}

	public static void WriteString(string str, byte[] b, ref int offset) {
		WriteBytes(System.Text.Encoding.UTF8.GetBytes(str), b, ref offset);
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

	public static ulong ReadUvarint(byte[] b, ref int offset) {
		ulong value = 0;
		for (int i = 0, s = 0; ;i++, s += 7) {
			var x = b[offset++];
			if (x < 0x80) {
				if (i > 9 || i == 9 && x > 1) {
					value = 0;
					throw new Exception("uvarint overflow");
				}
				value |= (ulong)x << s;
				break;
			}
			value |= (ulong)(x & 0x7f) << s;
		}
		return value;
	}

	public static void WriteUvarint(ulong v, byte[] b, ref int offset) {
		while (v >= 0x80) {
			b[offset ++] = (byte)(v | 0x80);
			v >>= 7;
		}
		b[offset ++] = (byte)v;
	}

	public static int UvarintSize(ulong v) {
		int i = 0;
		while (v >= 0x80) {
			i ++;
			v >>= 7;
		}
		return i + 1;
	}

	public static long ReadVarint(byte[] b, ref int offset) {
		ulong ux = ReadUvarint(b, ref offset);
		long value = (long)(ux >> 1);
		if ((ux & 1) != 0)
			value ^= value;
		return value;
	}

	public static void WriteVarint(long v, byte[] b, ref int offset) {
		var ux = (ulong)v << 1;
		if (v < 0)
			ux ^= ux;
		WriteUvarint(ux, b, ref offset);
	}

	public static int VarintSize(long v) {
		var ux = (ulong)v << 1;
		if (v < 0)
			ux ^= ux;
		return UvarintSize(ux);
	}
	
	public static int StringSize(string str) {
		var count = System.Text.Encoding.UTF8.GetByteCount(str);
		return UvarintSize((ulong)count) + count;
	}
}