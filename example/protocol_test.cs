using System;
using System.Collections.Generic;
using Xunit;

namespace example
{
    public class UnitTest1
    {
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

            Assert.Equal(msg1.Byte, msg2.Byte);
            Assert.Equal(msg1.Int, msg2.Int);
            Assert.Equal(msg1.Uint, msg2.Uint);
            Assert.Equal(msg1.Int8, msg2.Int8);
            Assert.Equal(msg1.Uint8, msg2.Uint8);
            Assert.Equal(msg1.Int16, msg2.Int16);
            Assert.Equal(msg1.Uint16, msg2.Uint16);
            Assert.Equal(msg1.Int32, msg2.Int32);
            Assert.Equal(msg1.Uint32, msg2.Uint32);
            Assert.Equal(msg1.Int64, msg2.Int64);
            Assert.Equal(msg1.Uint64, msg2.Uint64);
            Assert.Equal(msg1.Float32, msg2.Float32);
            Assert.Equal(msg1.Float64, msg2.Float64);
            Assert.Equal(msg1.String, msg2.String);
            Assert.Equal(msg1.Bytes, msg2.Bytes);
            Assert.Equal(msg1.Bool, msg2.Bool);
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
    }
}
