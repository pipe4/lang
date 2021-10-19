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
	file.PostParse()
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
	ast.PostParse()
	return ast, nil
}

func (f *File) PostParse() {
	f.Walk(func(s StatementWithContext) {
		switch s.Context {
		case DefaultContext:
			s.PostParseNameToType()
		case PropsContext:
			s.PostParseNameToType()
		}
	})
}

type StatementWithContext struct {
	Context StatementContext
	*Statement
}

func (s *Statement) Walk(walker func(s StatementWithContext)) {
	if s == nil {
		return
	}
	que := []StatementWithContext{{FileContext, s}}
	var next StatementWithContext

	for len(que) > 0 {
		next, que = que[len(que)-1], que[:len(que)-1]

		props := next.GetProps()
		for i := range props {
			que = append(que, StatementWithContext{PropsContext, &props[i]})
		}

		structure := next.GetStruct()
		for i := range structure {
			que = append(que, StatementWithContext{StructContext, &structure[i]})
		}

		if next.Default != nil {
			que = append(que, StatementWithContext{DefaultContext, next.Default})
		}

		walker(next)
	}
}

func (f *File) Walk(walker func(s StatementWithContext)) {
	for i := range f.Statements {
		f.Statements[i].Walk(walker)
	}
}

type StatementContext string

const (
	FileContext    = "FileContext"
	PropsContext   = "PropsContext"
	StructContext  = "StructContext"
	DefaultContext = "DefaultContext"
)

func (s *Statement) PostParseNameToType() {
	if s == nil {
		return
	}
	if s.Type == "" {
		s.Type = s.Name
		s.Name = ""
	}
}

func GetBnf() string {
	return parser.String()
}

func Parser() *participle.Parser {
	return parser
}
