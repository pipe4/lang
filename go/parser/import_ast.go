package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	pipe4ast "github.com/pipe4/lang/pipe4/ast"
)

func ImportAst(url string) (pipe4ast.NodeList, error) {
	// positions are relative to fileSet
	fileSet := token.NewFileSet()
	mode := parser.ParseComments | parser.AllErrors
	pkgAst, err := parser.ParseDir(fileSet, url, nil, mode)

	if err != nil {
		return nil, fmt.Errorf("failed to import go ast from %v: %w", url, err)
	}

	list := &pipe4ast.NodeList{}

	for _, pkg := range pkgAst {
		ctx := &Ctx{
			pkg: pkg,
			Package: StructCtx{
				NodeList: list,
			},
		}
		ast.Walk(&ctx.Package, ctx.pkg)
	}

	return *list, nil
}
