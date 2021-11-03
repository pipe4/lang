package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

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

func ParseString(path string, source string) (*File, error) {
	file := &File{}
	source = strings.ReplaceAll(source, "\r", "")
	if err := parser.ParseString(path, source, file); err != nil {
		return nil, fmt.Errorf("failed parse pipe4 file: %v: %w", path, err)
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
	// file, err := os.Open(path)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed read file %v: %w", path, err)
	}
	ast, err := ParseString(path, string(file))
	if err != nil {
		return nil, err
	}
	return ast, nil
}

// func (f *File) PostParse() {
// 	f.Walk(func(s StatementWithContext) {
// 		switch s.Context {
// 		case DefaultContext:
// 			s.PostParseNameToType()
// 		case ArgsContext:
// 			s.PostParseNameToType()
// 		}
// 	})
// }

// type StatementWithContext struct {
// 	*Statement
//
// 	Parent *Statement
// 	Prev   *Statement
// 	Next   *Statement
// }
//
// func (s *Statement) Walk(walker func(s *Statement, parent *Statement)) {
// 	if s == nil {
// 		return
// 	}
// 	que := []StatementWithContext{{FileContext, s}}
// 	var next StatementWithContext
//
// 	for len(que) > 0 {
// 		// que.pop()
// 		next, que = que[len(que)-1], que[:len(que)-1]
//
// 		props := next.GetArgs()
// 		for i := range props {
// 			que = append(que, StatementWithContext{ArgsContext, &props[i]})
// 		}
//
// 		structure := next.GetStruct()
// 		for i := range structure {
// 			que = append(que, StatementWithContext{StructContext, &structure[i]})
// 		}
//
// 		if next.Default != nil {
// 			que = append(que, StatementWithContext{DefaultContext, next.Default})
// 		}
//
// 		walker(next)
// 	}
// }
//
// func (f *File) Walk(walker func(s StatementWithContext)) {
// 	for i := range f.Statements {
// 		f.Statements[i].Walk(walker)
// 	}
// }
//
// type StatementContext string
//
// const (
// 	FileContext    = "FileContext"
// 	ArgsContext    = "ArgsContext"
// 	StructContext  = "StructContext"
// 	DefaultContext = "DefaultContext"
// )

// func (s *Statement) PostParseNameToType() {
// 	if s == nil {
// 		return
// 	}
// 	if s.Type == "" {
// 		s.Type = s.Name
// 		s.Name = ""
// 	}
// }

func GetBnf() string {
	return parser.String()
}

func Parser() *participle.Parser {
	return parser
}
