package ast

import "github.com/pipe4/lang/pipe4/parser"

type importStatement struct {
	rename declarationName // local-rename
	url    importUrl
}

type importUrl struct {
	url string
}

func importFromParserAst(parserImport parser.Import) importStatement {
	return importStatement{
		rename: declarationName{name: parserImport.Name},
		url:    importUrl{url: parserImport.Url},
	}
}
