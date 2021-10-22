package internal

import (
	"github.com/pipe4/lang/pipe4/parser"
)

type File struct {
	name string // example.pipe4

	imports    []importStatement
	statements []statement
}

func FileFromParser(s parser.File) File {
	f := File{
		name: s.Name,
	}
	s.Walk(func(s parser.StatementWithContext) {
		for _, parserImport := range s.Import {
			f.imports = append(f.imports, importFromParserAst(parserImport))
		}
	})

	f.statements = statementsFromParser(s.Statements)
	return f
}
