package gb

import (
	"errors"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
)

type file struct {
	Package string
	Consts  []*types.Const
	Structs map[string]*types.Struct
}

func parseFile(name string) (*file, error) {
	stat, err := os.Stat(name)
	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		return nil, errors.New("gobuf: \"'" + name + "'\" not a file")
	}

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, name, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	conf := types.Config{
		Importer:         importer.Default(),
		IgnoreFuncBodies: true,
	}

	pkg, err := conf.Check("", fset, []*ast.File{f}, nil)
	if err != nil {
		return nil, err
	}

	return &file{
		Package: f.Name.Name,
		Consts:  scanConsts(f, pkg),
		Structs: scanStructs(f, pkg),
	}, nil
}

func scanConsts(f *ast.File, pkg *types.Package) []*types.Const {
	var consts []*types.Const

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		if genDecl.Tok != token.CONST {
			continue
		}

		for _, spec := range genDecl.Specs {
			valSpec, _ := spec.(*ast.ValueSpec)

			for _, name := range valSpec.Names {
				consts = append(consts, pkg.Scope().Lookup(name.Name).(*types.Const))
			}
		}
	}

	return consts
}

func scanStructs(f *ast.File, pkg *types.Package) map[string]*types.Struct {
	messages := make(map[string]*types.Struct)

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		if genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, _ := spec.(*ast.TypeSpec)

			obj := pkg.Scope().Lookup(typeSpec.Name.Name)

			typeInfo, ok := obj.Type().Underlying().(*types.Struct)
			if !ok {
				continue
			}

			messages[typeSpec.Name.Name] = typeInfo
		}
	}

	return messages
}
