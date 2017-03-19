What's is Gobuf?
================

Gobuf is a binary serialization format like protobuf.

But it didn't need any DSL file to define messages.

It just needs you write down data format by Go structs.

It can analize Go struct and generate marshaller and unmarshaller code automaticlly.

And it not only supports generate Go code but also C#.

It has a good plugin mechanism, let you can add different programming language generation easily.

Pipeline
============

```
 +-----------+           +-----------------+        +----------+          +-------------------+
 |  Go code  <-----------+  Gobuf Command  +-------->  Plugin  +---------->  Target Language  |
 +-----------+  analyze  +-----------------+  JSON  +----------+  output  +-------------------+
```

1. Gobuf command line tool analyze a Go source file.
2. Gobuf output type informations to command line standard output.
3. Plugin use type informations to generate target language source code.

Example:

```
gobuf protocol.go | gobuf-cs > protocol.gb.cs
```

Types
=====

Gobuf supports almost all kinds of Go data type, but it has some limits.

This is all the data types that Gobuf supports:

* Type = Scalar | Composite
* Scalar = Numberic | "string" | "bool"
* Numeric = "int" | "uint" | "int8" | "uint8" | "int16" | "uint16" | "int32" | "uint32" | "int64" | "uint64" | "float32" | "float64"
* Composite = Pointer | Array | Map | Bytes | Struct
* Pointer = "\*" ( Scalar | Struct )
* Array = "[" "]" Type
* Bytes = "[" "]" "byte"
* Map = "map" "[" Scalar "]" Type
* Struct = "type" Name "struct" "{" Name Type [";" Name Type] "}"

Type Maping
===========

Gobuf maping Go type to different programming languages.

This is the type maping rule:

| Go | C# |
| -- | -- |
| int | long |
| uint | ulong |
| int8 | sbyte |
| uint8 | byte |
| int16 | short |
| uint16 | ushort |
| int32 | int |
| uint32 | uint |
| int64 | long |
| uint64 | ulong |
| float32 | float |
| float64 | double |
| bool | bool |
| string | utf8 string |
| []byte | byte[] |
| []Type | List\<Type\> |
| \*Scalar | Nullable\<Scalar\> |
| \*Struct | Class |
| map[Scalar]Type | Dictionary\<Scalar, Type\> |
| Struct | Class |

Wired Format
============

| Type | Format |
| -- | -- |
| int | base128 varint + zig-zag |
| uint | base128 varint |
| int8, uint8, bool | 1byte |
| int16, uint16 | 2byte little-endian |
| int32, uint32 | 4byte little-endian |
| int64, uint64 | 8byte little-endian |
| float32 | convert to uint32 bits |
| float64 | convert to uint64 bits |
| string | uint(utf8\_length) + utf8\_length |
| []byte | uint(length) + length |
| []Type | uint(count) + items |
| \*Scalar | is_null ? uint8(0) : uint8(1) + element |
| \*Struct | is_null ? uint8(0) : uint8(1) + fields |
| map[Scalar]Type | uint(count) + items(key + value) |
| Struct | fields |
