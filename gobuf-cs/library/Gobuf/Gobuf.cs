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
		int s = 0;
		int i = 0;
		ulong v = 0;
		while (offset < b.Length) {
			var x = b[offset++];
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
		return v;
	}

	public static void WriteUvarint(ulong v, byte[] b, ref int offset) {
		var i = 0;
		while (v >= 0x80) {
			b[i] = (byte)((byte)(v) | (byte)0x80);
			v >>= 7;
			i++;
		}
		b[i] = (byte)(v);
		offset = i + 1;
	}

	public static int UvarintSize(ulong x) {
		var i = 0;
		while (x >= 0x80) {
			x >>= 7;
			i++;
		}
		return i + 1;
	}

	public static long ReadVarint(byte[] b, ref int offset) {
		return Zag(ReadUvarint(b, ref offset));
	}

	public static void WriteVarint(long v, byte[] b, ref int offset) {
		WriteUvarint(Zig(v), b, ref offset);
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

	public static int StringSize(string str) {
		var count = System.Text.Encoding.UTF8.GetByteCount(str);
		return UvarintSize((ulong)count) + count;
	}
}