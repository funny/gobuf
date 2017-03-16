Scalar Types
====

| Type | Encoding |
| --- | --- |
| int | zig-zag + base128 uvarint |
| uint | base128 uvarint |
| byte, int8, uint8 | 1byte |
| int16, uint16 | 2byte little-endian |
| int32, uint32 | 4byte little-endian |
| int64, uint64 | 8byte little-endian |
| float32 | float32 to uint32 |
| float64 | float64 to uint64 |
| string, []byte | uvarint(length) + length |

Composite Types
=====

| Type | Supports | Encoding |
| --- | --- | --- |
| Message | struct | foreach(field) |
| Array | []{Composite \| Scalar} | uvarint(count) + foreach(item) |
| Map | map[Scalar]{Composite \| Scalar} | uvarint(count) + foreach(key, value)
| Pointer | *{Message \| Scalar} | byte(0) or byte(1) + {Scalar \| Message} |
