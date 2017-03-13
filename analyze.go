package gobuf

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
	type Array interface {
		Elem() types.Type
	}

	switch tt := t.(type) {
	case *types.Pointer:
		if st := analyzeScalar(tt.Elem()); st != nil {
			st.IsPointer = true
			return st
		}
	case Array:
		if st := analyzeScalar(tt.Elem()); st != nil {
			if st.Kind == UINT8 {
				st.Kind = BYTES
			} else {
				st.IsArray = true
			}
			return st
		}
	}

	return analyzeScalar(t)
}

func analyzeScalar(t types.Type) *Type {
	switch tt := t.(type) {
	case *types.Basic:
		kind := kindOfType(tt)
		if kind != "" {
			return &Type{Kind: kind}
		}
	case *types.Named:
		if tt2, ok := tt.Underlying().(*types.Basic); ok {
			kind := kindOfType(tt2)
			if kind != "" {
				return &Type{Kind: kind, Name: t.String()}
			}
		}
		if _, ok := tt.Underlying().(*types.Struct); ok {
			return &Type{MESSAGE, tt.String(), false, false}
		}
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
	}
	return ""
}
