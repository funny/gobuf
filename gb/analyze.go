package gb

import (
	"fmt"
	"go/types"
)

func analyzeFile(f *file) (*Doc, error) {
	var enums []*Enum
	var curEnum *Enum

	for _, c := range f.Consts {
		if curEnum == nil || c.Type().String() != curEnum.Name {
			t, ok := c.Type().Underlying().(*types.Basic)
			if !ok {
				return nil, fmt.Errorf("gobuf: unsupported const type \"%s\"", c.Type())
			}

			kind := kindOfType(t)
			if kind == "" {
				return nil, fmt.Errorf("gobuf: unsupported const type \"%s\"", c.Type())
			}

			curEnum = &Enum{Kind: kind, Name: c.Type().String()}
			enums = append(enums, curEnum)
		}

		curEnum.Values = append(curEnum.Values, &Value{
			Name:  c.Name(),
			Value: c.Val().ExactString(),
		})
	}

	var msgs []*Message

	for name, s := range f.Messages {
		msg := &Message{
			Name:   name,
			Fields: make([]*Field, s.NumFields()),
		}
		msgs = append(msgs, msg)

		for i := 0; i < s.NumFields(); i++ {
			field := s.Field(i)

			msg.Fields[i] = &Field{
				Name: field.Name(),
				Type: analyzeType(field.Type()),
			}

			if msg.Fields[i].Type == nil {
				return nil, fmt.Errorf("gobuf: unsupported field type \"%s\"", field.Type().String())
			}
		}
	}

	return &Doc{f.Package, enums, msgs}, nil
}

func analyzeType(t types.Type) *Type {
	array := analyzeArray(t)
	if array != nil {
		return array
	}

	mapType := analyzeMap(t)
	if mapType != nil {
		return mapType
	}

	pointer := analyzePointer(t)
	if pointer != nil {
		return pointer
	}

	return analyzeScalar(t)
}

func analyzeNamed(t types.Type) (string, types.Type) {
	if named, ok := t.(*types.Named); ok {
		return named.String(), named.Underlying()
	}
	return "", t
}

func analyzeArray(t types.Type) *Type {
	_, isArray := t.(*types.Array)
	_, isSlice := t.(*types.Slice)

	if isArray || isSlice {
		type Array interface {
			Elem() types.Type
		}

		if array, ok := t.(Array); ok {
			var length int

			if array2, ok := t.(*types.Array); ok {
				length = int(array2.Len())
			}

			elem := analyzeType(array.Elem())
			if elem.Kind == UINT8 {
				return &Type{Kind: BYTES, Len: length}
			}

			return &Type{Kind: ARRAY, Elem: elem, Len: length}
		}
	}

	return nil
}

func analyzeMap(t types.Type) *Type {
	if mapType, ok := t.(*types.Map); ok {
		key := analyzeScalar(mapType.Key())
		elem := analyzeType(mapType.Elem())
		if key != nil && elem != nil {
			return &Type{Kind: MAP, Key: key, Elem: elem}
		}
	}
	return nil
}

func analyzePointer(t types.Type) *Type {
	if pointer, ok := t.(*types.Pointer); ok {
		return &Type{Kind: POINTER, Elem: analyzeScalar(pointer.Elem())}
	}
	return nil
}

func analyzeScalar(t types.Type) *Type {
	name, t2 := analyzeNamed(t)
	if basic, ok := t2.(*types.Basic); ok {
		kind := kindOfType(basic)
		if kind != "" {
			return &Type{Kind: kind, Name: name}
		}
	}
	if _, ok := t2.(*types.Struct); ok {
		return &Type{Kind: MESSAGE, Name: name}
	}
	return nil
}

func kindOfType(t *types.Basic) string {
	switch t.Kind() {
	case types.Int:
		return INT
	case types.Uint:
		return UINT
	case types.Int8:
		return INT8
	case types.Uint8:
		return UINT8
	case types.Int16:
		return INT16
	case types.Uint16:
		return UINT16
	case types.Int32:
		return INT32
	case types.Uint32:
		return UINT32
	case types.Int64:
		return INT64
	case types.Uint64:
		return UINT64
	case types.Float32:
		return FLOAT32
	case types.Float64:
		return FLOAT64
	case types.String:
		return STRING
	case types.Bool:
		return BOOL
	}
	return ""
}
