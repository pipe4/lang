package loader

import (
	"fmt"
	"go/types"
	"log"
	"strings"

	"github.com/pipe4/lang/pipe4/ast"
	"golang.org/x/tools/go/loader"
)

func resolve(ident ast.Ident) (*ast.Node, error) {
	var conf loader.Config
	pkg := ident.GoImport()
	pkg = strings.TrimPrefix(pkg, "lang/go/")

	conf.Import(pkg)
	prog, err := conf.Load()
	log.Println(prog)
	if err != nil {
		return nil, fmt.Errorf("failed to load go import for %v: %w", ident, err)
	}
	scope := prog.Package(pkg).Pkg.Scope()
	nodeObj := scope.Lookup(ident.Name)

	node := ast.Node{
		Ident: ident,
	}
	nodeType, err := GoTypeToPipe4(nodeObj.Type())
	if err != nil {
		return nil, fmt.Errorf("failed to convert type for %v: %w", ident, err)
	}
	node.Type = nodeType
	return &node, nil
}

func GoTypeToPipe4(in types.Type) (ast.Type, error) {
	out := ast.Type{}
	switch def := in.(type) {
	case *types.Signature:
		out.Ident = ast.Ident{Name: "go.func"}
		params := def.Params()
		for i := 0; i < params.Len(); i++ {
			param := params.At(i)

			paramType, err := GoTypeToPipe4(param.Type())
			if err != nil {
				return out, fmt.Errorf("failed to convert %v func parameter: %w", i, err)
			}
			out.Args = append(out.Args, ast.Node{
				Ident: ast.Ident{Name: param.Name()},
				Type:  paramType,
			})
		}

		log.Println("found", def.String())
	case *types.Basic:
		out.Ident = ast.Ident{Name: "go." + def.String()}
		return out, nil
	case *types.Interface:
		out.Ident = ast.Ident{Name: "go." + def.String()}
		return out, nil
	case *types.Slice:
		out.Ident = ast.Ident{Name: "array"}
		elemType, err := GoTypeToPipe4(def.Elem())
		if err != nil {
			return out, fmt.Errorf("failed to convert slice type: %w", err)
		}
		out.Args = []ast.Node{{Type: elemType}}
		return out, nil
	default:
		return out, fmt.Errorf("unimplemented type: %v", def.String())
	}
	return out, nil
}
