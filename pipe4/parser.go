package pipe4

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
)

var parser = participle.MustBuild(
	&File{},
	participle.Lexer(pipe4Lexer),
	participle.Elide("Comment", "Whitespace"),
	participle.UseLookahead(2),
)

func ParseString(source string) (*File, error) {
	file := &File{}
	if err := parser.ParseString("", source, file); err != nil {
		return nil, fmt.Errorf("failed parse pipe4 source: %w", err)
	}
	return file, nil
}

func ParseFile(path string) (*File, error) {
	ast := &File{}
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed open file %v: %w", path, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("failed to close file %v: %+v", path, err)
		}
	}()

	if err := parser.Parse(path, file, ast); err != nil {
		return nil, fmt.Errorf("failed parse pipe4 source: %w", err)
	}
	return ast, nil
}

func GetBnf() string {
	return parser.String()
}
