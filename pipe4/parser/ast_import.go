package parser

import "github.com/alecthomas/participle/v2/lexer"

type Import struct {
	Name string `parser:"@Ident?" yaml:"Name,omitempty"`
	URL  string `parser:"@String" yaml:"URL,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}
type SingleImport struct {
	Import Import `parser:"'import' @@ " yaml:"Import,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}
type BlockImport struct {
	Imports []Import `parser:"'import' '(' @@* ')'" yaml:"Imports,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}
