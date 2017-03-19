using System;
using System.Collections.Generic;
using System.Net.Sockets;
using Xunit;

namespace example
{
    public class UnitTest1
    {
        delegate void CheckFunc(Scalar msg1, Scalar msg2);

        CheckFunc check = (Scalar msg11, Scalar msg22) => {
            Assert.Equal(msg11.Byte, msg22.Byte);
            Assert.Equal(msg11.Int, msg22.Int);
            Assert.Equal(msg11.Uint, msg22.Uint);
            Assert.Equal(msg11.Int8, msg22.Int8);
            Assert.Equal(msg11.Uint8, msg22.Uint8);
            Assert.Equal(msg11.Int16, msg22.Int16);
            Assert.Equal(msg11.Uint16, msg22.Uint16);
            Assert.Equal(msg11.Int32, msg22.Int32);
            Assert.Equal(msg11.Uint32, msg22.Uint32);
            Assert.Equal(msg11.Int64, msg22.Int64);
            Assert.Equal(msg11.Uint64, msg22.Uint64);
            Assert.Equal(msg11.Float32, msg22.Float32);
            Assert.Equal(msg11.Float64, msg22.Float64);
            Assert.Equal(msg11.String, msg22.String);
            Assert.Equal(msg11.Bytes, msg22.Bytes);
            Assert.Equal(msg11.Bool, msg22.Bool);
        };

        [Fact]
        public void TestScalar()
        {
            var msg1 = new Scalar();

            msg1.Byte = System.Byte.MaxValue;
            msg1.Int = System.Int64.MaxValue;
            msg1.Uint = System.UInt64.MaxValue;
            msg1.Int8 = System.SByte.MaxValue;
            msg1.Uint8 = System.Byte.MaxValue;
            msg1.Int16 = System.Int16.MaxValue;
            msg1.Uint16 = System.UInt16.MaxValue;
            msg1.Int32 = System.Int32.MaxValue;
            msg1.Uint32 = System.UInt32.MaxValue;
            msg1.Int64 = System.Int64.MaxValue;
            msg1.Uint64 = System.UInt64.MaxValue;
            msg1.Float32 = System.Single.MaxValue;
            msg1.Float64 = System.Double.MaxValue;
            msg1.String =  "test string content";
            msg1.Bytes = new byte[]{1,2,3,4};
            msg1.Bool = true;

            var data = new byte[msg1.Size()];

            var size1 = msg1.Marshal(data, 0);

            Assert.Equal(size1, data.Length);

            var msg2 = new Scalar();

            var size2 = msg2.Unmarshal(data, 0);

            Assert.Equal(size2, data.Length);

            check(msg1, msg2);
        }

        [Fact]
        public void TestPointer()
        {
            var msg1 = new Pointer();

            msg1.IntPtr = System.Int64.MaxValue;
            msg1.UintPtr = System.UInt64.MaxValue;
            msg1.Int8Ptr = System.SByte.MaxValue;
            msg1.Uint8Ptr = System.Byte.MaxValue;
            msg1.Int16Ptr = System.Int16.MaxValue;
            msg1.Uint16Ptr = System.UInt16.MaxValue;
            msg1.Int32Ptr = System.Int32.MaxValue;
            msg1.Uint32Ptr = System.UInt32.MaxValue;
            msg1.Int64Ptr = System.Int64.MaxValue;
            msg1.Uint64Ptr = System.UInt64.MaxValue;
            msg1.Float32Ptr = System.Single.MaxValue;
            msg1.Float64Ptr = System.Double.MaxValue;
            msg1.StringPtr =  "test string content";
            msg1.BoolPtr = true;

            var data = new byte[msg1.Size()];

            var size1 = msg1.Marshal(data, 0);

            Assert.Equal(size1, data.Length);

            var msg2 = new Pointer();

            var size2 = msg2.Unmarshal(data, 0);

            Assert.Equal(size2, data.Length);

            Assert.Equal(msg1.IntPtr, msg2.IntPtr);
            Assert.Equal(msg1.UintPtr, msg2.UintPtr);
            Assert.Equal(msg1.Int8Ptr, msg2.Int8Ptr);
            Assert.Equal(msg1.Uint8Ptr, msg2.Uint8Ptr);
            Assert.Equal(msg1.Int16Ptr, msg2.Int16Ptr);
            Assert.Equal(msg1.Uint16Ptr, msg2.Uint16Ptr);
            Assert.Equal(msg1.Int32Ptr, msg2.Int32Ptr);
            Assert.Equal(msg1.Uint32Ptr, msg2.Uint32Ptr);
            Assert.Equal(msg1.Int64Ptr, msg2.Int64Ptr);
            Assert.Equal(msg1.Uint64Ptr, msg2.Uint64Ptr);
            Assert.Equal(msg1.Float32Ptr, msg2.Float32Ptr);
            Assert.Equal(msg1.Float64Ptr, msg2.Float64Ptr);
            Assert.Equal(msg1.StringPtr, msg2.StringPtr);
            Assert.Equal(msg1.BoolPtr, msg2.BoolPtr);
        }

        [Fact]
        public void TestArray()
        {
            var msg1 = new Array();

            msg1.IntArray.Add(System.Int64.MaxValue);
            msg1.UintArray.Add(System.UInt64.MaxValue);
            msg1.Int8Array.Add(System.SByte.MaxValue);
            msg1.Uint8Array = new byte[]{System.Byte.MaxValue};
            msg1.Int16Array.Add(System.Int16.MaxValue);
            msg1.Uint16Array.Add(System.UInt16.MaxValue);
            msg1.Int32Array.Add(System.Int32.MaxValue);
            msg1.Uint32Array.Add(System.UInt32.MaxValue);
            msg1.Int64Array.Add(System.Int64.MaxValue);
            msg1.Uint64Array.Add(System.UInt64.MaxValue);
            msg1.Float32Array.Add(System.Single.MaxValue);
            msg1.Float64Array.Add(System.Double.MaxValue);
            msg1.StringArray.Add("test string content");
            msg1.BoolArray.Add(true);

            var data = new byte[msg1.Size()];

            var size1 = msg1.Marshal(data, 0);

            Assert.Equal(size1, data.Length);

            var msg2 = new Array();

            var size2 = msg2.Unmarshal(data, 0);

            Assert.Equal(size2, data.Length);

            Assert.Equal(msg1.IntArray, msg2.IntArray);
            Assert.Equal(msg1.UintArray, msg2.UintArray);
            Assert.Equal(msg1.Int8Array, msg2.Int8Array);
            Assert.Equal(msg1.Uint8Array, msg2.Uint8Array);
            Assert.Equal(msg1.Int16Array, msg2.Int16Array);
            Assert.Equal(msg1.Uint16Array, msg2.Uint16Array);
            Assert.Equal(msg1.Int32Array, msg2.Int32Array);
            Assert.Equal(msg1.Uint32Array, msg2.Uint32Array);
            Assert.Equal(msg1.Int64Array, msg2.Int64Array);
            Assert.Equal(msg1.Uint64Array, msg2.Uint64Array);
            Assert.Equal(msg1.Float32Array, msg2.Float32Array);
            Assert.Equal(msg1.Float64Array, msg2.Float64Array);
            Assert.Equal(msg1.StringArray, msg2.StringArray);
            Assert.Equal(msg1.BoolArray, msg2.BoolArray);
        }

        [Fact]
        public void TestFixlenArray()
        {
            var msg1 = new FixlenArray();

            msg1.IntArray[0] = System.Int64.MaxValue;
            msg1.UintArray[0] = System.UInt64.MaxValue;
            msg1.Int8Array[0] = System.SByte.MaxValue;
            msg1.Uint8Array[0] = System.Byte.MaxValue;
            msg1.Int16Array[0] = System.Int16.MaxValue;
            msg1.Uint16Array[0] = System.UInt16.MaxValue;
            msg1.Int32Array[0] = System.Int32.MaxValue;
            msg1.Uint32Array[0] = System.UInt32.MaxValue;
            msg1.Int64Array[0] = System.Int64.MaxValue;
            msg1.Uint64Array[0] = System.UInt64.MaxValue;
            msg1.Float32Array[0] = System.Single.MaxValue;
            msg1.Float64Array[0] = System.Double.MaxValue;
            msg1.StringArray[0] = "test string content";
            msg1.BoolArray[0] = true;

            var data = new byte[msg1.Size()];

            var size1 = msg1.Marshal(data, 0);

            Assert.Equal(size1, data.Length);

            var msg2 = new FixlenArray();

            var size2 = msg2.Unmarshal(data, 0);

            Assert.Equal(size2, data.Length);

            Assert.Equal(msg1.IntArray, msg2.IntArray);
            Assert.Equal(msg1.UintArray, msg2.UintArray);
            Assert.Equal(msg1.Int8Array, msg2.Int8Array);
            Assert.Equal(msg1.Uint8Array, msg2.Uint8Array);
            Assert.Equal(msg1.Int16Array, msg2.Int16Array);
            Assert.Equal(msg1.Uint16Array, msg2.Uint16Array);
            Assert.Equal(msg1.Int32Array, msg2.Int32Array);
            Assert.Equal(msg1.Uint32Array, msg2.Uint32Array);
            Assert.Equal(msg1.Int64Array, msg2.Int64Array);
            Assert.Equal(msg1.Uint64Array, msg2.Uint64Array);
            Assert.Equal(msg1.Float32Array, msg2.Float32Array);
            Assert.Equal(msg1.Float64Array, msg2.Float64Array);
            Assert.Equal(msg1.StringArray, msg2.StringArray);
            Assert.Equal(msg1.BoolArray, msg2.BoolArray);
        }

        [Fact]
        public void TestMap()
        {
            var msg1 = new Map();

            msg1.IntMap.Add(1, System.Int64.MaxValue); 
            msg1.UintMap.Add(1, System.UInt64.MaxValue); 
            msg1.Int8Map.Add(1, System.SByte.MaxValue); 
            msg1.Uint8Map.Add(1, System.Byte.MaxValue); 
            msg1.Int16Map.Add(1, System.Int16.MaxValue); 
            msg1.Uint16Map.Add(1, System.UInt16.MaxValue); 
            msg1.Int32Map.Add(1, System.Int32.MaxValue); 
            msg1.Uint32Map.Add(1, System.UInt32.MaxValue); 
            msg1.Int64Map.Add(1, System.Int64.MaxValue); 
            msg1.Uint64Map.Add(1, System.UInt64.MaxValue); 
            msg1.Float32Map.Add(1, System.Single.MaxValue); 
            msg1.Float64Map.Add(1, System.Double.MaxValue); 
            msg1.StringMap.Add(1, "test string content"); 
            msg1.BoolMap.Add(1, true); 

            var data = new byte[msg1.Size()];

            var size1 = msg1.Marshal(data, 0);

            Assert.Equal(size1, data.Length);

            var msg2 = new Map();

            var size2 = msg2.Unmarshal(data, 0);

            Assert.Equal(size2, data.Length);

            Assert.Equal(msg1.IntMap, msg2.IntMap);
            Assert.Equal(msg1.UintMap, msg2.UintMap);
            Assert.Equal(msg1.Int8Map, msg2.Int8Map);
            Assert.Equal(msg1.Uint8Map, msg2.Uint8Map);
            Assert.Equal(msg1.Int16Map, msg2.Int16Map);
            Assert.Equal(msg1.Uint16Map, msg2.Uint16Map);
            Assert.Equal(msg1.Int32Map, msg2.Int32Map);
            Assert.Equal(msg1.Uint32Map, msg2.Uint32Map);
            Assert.Equal(msg1.Int64Map, msg2.Int64Map);
            Assert.Equal(msg1.Uint64Map, msg2.Uint64Map);
            Assert.Equal(msg1.Float32Map, msg2.Float32Map);
            Assert.Equal(msg1.Float64Map, msg2.Float64Map);
            Assert.Equal(msg1.StringMap, msg2.StringMap);
            Assert.Equal(msg1.BoolMap, msg2.BoolMap);
        }

        [Fact]
        public void TestMessage()
        {
            var scalar = new Scalar();

            scalar.Byte = System.Byte.MaxValue;
            scalar.Int = System.Int64.MaxValue;
            scalar.Uint = System.UInt64.MaxValue;
            scalar.Int8 = System.SByte.MaxValue;
            scalar.Uint8 = System.Byte.MaxValue;
            scalar.Int16 = System.Int16.MaxValue;
            scalar.Uint16 = System.UInt16.MaxValue;
            scalar.Int32 = System.Int32.MaxValue;
            scalar.Uint32 = System.UInt32.MaxValue;
            scalar.Int64 = System.Int64.MaxValue;
            scalar.Uint64 = System.UInt64.MaxValue;
            scalar.Float32 = System.Single.MaxValue;
            scalar.Float64 = System.Double.MaxValue;
            scalar.String =  "test string content";
            scalar.Bytes = new byte[]{1,2,3,4};
            scalar.Bool = true;

            var msg1 = new Message();
            msg1.Scalar = scalar;
            msg1.ScalarPtr = scalar;
            msg1.ScalarArray.Add(scalar);
            msg1.ScalarMap.Add(1, scalar);

            var data = new byte[msg1.Size()];

            var size1 = msg1.Marshal(data, 0);

            Assert.Equal(size1, data.Length);

            var msg2 = new Message();

            var size2 = msg2.Unmarshal(data, 0);

            Assert.Equal(size2, data.Length);

            check(msg1.Scalar, msg2.Scalar);
            check(msg1.ScalarPtr, msg2.ScalarPtr);
            check(msg1.ScalarArray[0], msg2.ScalarArray[0]);
            check(msg1.ScalarMap[1], msg2.ScalarMap[1]);
        }

        [Fact]
        public void TestCommunication()
        {
            var scalar = new Scalar();

            scalar.Byte = System.Byte.MaxValue;
            scalar.Int = System.Int64.MaxValue;
            scalar.Uint = System.UInt64.MaxValue;
            scalar.Int8 = System.SByte.MaxValue;
            scalar.Uint8 = System.Byte.MaxValue;
            scalar.Int16 = System.Int16.MaxValue;
            scalar.Uint16 = System.UInt16.MaxValue;
            scalar.Int32 = System.Int32.MaxValue;
            scalar.Uint32 = System.UInt32.MaxValue;
            scalar.Int64 = System.Int64.MaxValue;
            scalar.Uint64 = System.UInt64.MaxValue;
            scalar.Float32 = System.Single.MaxValue;
            scalar.Float64 = System.Double.MaxValue;
            scalar.String =  "test string content";
            scalar.Bytes = new byte[]{1,2,3,4};
            scalar.Bool = true;

            var msg1 = new Message();
            msg1.Scalar = scalar;
            msg1.ScalarPtr = scalar;
            msg1.ScalarArray.Add(scalar);
            msg1.ScalarMap.Add(1, scalar);

            var body = new byte[msg1.Size()];
            var size = msg1.Marshal(body, 0);
            Assert.Equal(size, body.Length);

            var head = new byte[4];
            int offset = 0;
            Gobuf.WriteUint32((uint)size, head, ref offset);

            var conn = new TcpClient();
            var task = conn.ConnectAsync("127.0.0.1", 12345);
            task.Wait();
            Assert.True(conn.Connected);
            var stream = conn.GetStream();

            stream.Write(head, 0, 4);
            stream.Write(body, 0, body.Length);

            offset = 0;
            while (offset != 4) {
                offset += stream.Read(head, offset, 4 - offset);
            }

            offset = 0;
            size = (int)Gobuf.ReadUint32(head, ref offset);
            Assert.Equal(size, body.Length);

            offset = 0;
            while (offset != body.Length) {
                offset += stream.Read(body, offset, body.Length - offset);
            }

            var msg2 = new Message();
            size = msg2.Unmarshal(body, 0);
            Assert.Equal(size, body.Length);
            
            check(msg1.Scalar, msg2.Scalar);
            check(msg1.ScalarPtr, msg2.ScalarPtr);
            check(msg1.ScalarArray[0], msg2.ScalarArray[0]);
            check(msg1.ScalarMap[1], msg2.ScalarMap[1]);
        }
    }
}
