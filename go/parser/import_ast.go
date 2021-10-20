package lang_go

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func ImportAst(url string) (map[string]*ast.Package, error) {
	// positions are relative to fileSet
	fileSet := token.NewFileSet()
	mode := parser.ParseComments | parser.AllErrors
	pkgAst, err := parser.ParseDir(fileSet, url, nil, mode)

	if err != nil {
		return nil, fmt.Errorf("failed to import go ast from %v: %w", url, err)
	}
	return pkgAst, nil
}
