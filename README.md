
Scalar Types
====

| Type | Encoding |
| --- | --- |
| int | varint |
| uint | uvarint |
| byte, int8, uint8 | 1byte |
| int16, uint16 | 2byte little-endian |
| int32, uint32 | 4byte little-endian |
| int64, uint64 | 8byte little-endian |

Composite Types
=====

| Type | Encoding |
| --- | --- |
| string, []byte | uvarint(length) + length |
| []{Scalar} | uvarint(count) + count * {Scalar} |
| []{Message} | uvarint(count) + foreach(content) |
| *{Scalar} | byte(0) / byte(1) + {Scalar} |
| *{Message} | byte(0) / byte(1) + {Message} |
