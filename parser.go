package lang

import (
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var parser = participle.MustBuild(
	&File{},
	participle.Lexer(pipe4Lexer),
	participle.Elide("Whitespace", "BlockCommentStart", "BlockCommentEnd"),
	participle.UseLookahead(2),
	participle.Unquote("String"),
	participle.Map(func(token lexer.Token) (lexer.Token, error) {
		// token.Value = strings.TrimPrefix(token.Value, "//")
		// token.Value = strings.TrimSuffix(token.Value, "\n")
		return token, nil
	}, "LineComment"),
)

func ParseString(source string) (*File, error) {
	file := &File{}
	if err := parser.ParseString("", source, file); err != nil {
		return nil, fmt.Errorf("failed parse pipe4 source: %w", err)
	}
	return file, nil
}

func LexFile(path string) ([]lexer.Token, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed open file %v: %w", path, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("failed to close file %v: %+v", path, err)
		}
	}()
	tokens, err := parser.Lex(path, file)
	if err != nil {
		return nil, fmt.Errorf("failed lex file %v: %w", path, err)
	}
	return tokens, err
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

func Parser() *participle.Parser {
	return parser
}
