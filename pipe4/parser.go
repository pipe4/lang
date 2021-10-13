package pipe4

import (
	"fmt"
	"github.com/alecthomas/participle/v2"
	"github.com/kr/pretty"
	"os"
)


func ParseFile(sourcePath string) (*File, error) {
	fileAst := &File{}

	parser, err := participle.Build(
		&File{},
		participle.Lexer(pipe4Lexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)

	if err != nil {
		return nil, fmt.Errorf("failed build pipe4 parser: %w", err)
	}

	fmt.Printf("EBNF:\n\n%v\n\n\n", parser.String())

	sourceReader, err := os.Open(sourcePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %v: %w", sourcePath, err)
	}
	err = parser.Parse(sourcePath, sourceReader, fileAst)
	_ = sourceReader.Close()

	fmt.Printf("%v AST:\n\n%# v\n\n\n", sourcePath, pretty.Formatter(fileAst))

	if err != nil {
		return nil, fmt.Errorf("failed parse %v: %w", sourcePath, err)
	}
	return fileAst, nil
}
