package parser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"gopkg.in/yaml.v3"
)

type Meta struct {
	Pos    lexer.Position `parser:"" json:"-"`
	EndPos lexer.Position `parser:"" json:"-"`
	Tokens []lexer.Token  `parser:"" json:"-"`
}

type File struct {
	Name       string     `parser:"" json:"-"`
	Statements Statements `parser:"EOS* (@@ EOS+)*" json:"Statements,omitempty"`
}

func (f *File) ToYaml() (string, error) {
	jsonOut := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(jsonOut)
	err := jsonEncoder.Encode(f)
	if err != nil {
		return "", fmt.Errorf("failed print ast tree: json.Marshal: %w", err)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(jsonOut.Bytes(), &m)
	if err != nil {
		return "", fmt.Errorf("failed print ast tree: json.Unmarshal: %w", err)
	}

	// pretty.Fprintf(os.Stdout, "%# v", ast)
	out := &strings.Builder{}
	encoder := yaml.NewEncoder(out)
	encoder.SetIndent(2)
	if err := encoder.Encode(m); err != nil {
		return "", fmt.Errorf("failed print ast tree: yaml.Encode: %w", err)
	}
	str := out.String()
	// str = strings.ReplaceAll(str, "  - ", "- ")
	return str, nil
}

func (f *File) FromYaml(source string) error {
	m := make(map[string]interface{})
	if err := yaml.Unmarshal([]byte(source), &m); err != nil {
		return fmt.Errorf("failed to read ast from json: yaml.Unmarshal: %w", err)
	}
	buf, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to read ast from json: json.Marshal: %w", err)
	}

	err = json.Unmarshal(buf, f)
	if err != nil {
		return fmt.Errorf("failed to read ast from json: json.Unmarshal: %w", err)
	}
	return nil
}

var parser = participle.MustBuild(
	&File{},
	participle.Lexer(pipe4Lexer),
	participle.Elide("Whitespace", "BlockCommentStart", "BlockCommentEnd"),
	participle.UseLookahead(10),
	participle.Unquote("String"),
	// participle.Map(func(token lexer.Token) (lexer.Token, error) {
	// 	// token.Value = strings.TrimPrefix(token.Value, "//")
	// 	// token.Value = strings.TrimSuffix(token.Value, "\n")
	// 	return token, nil
	// }, "LineComment"),
)

func ParseString(path string, source string) (*File, error) {
	file := &File{}
	source = strings.ReplaceAll(source, "\r", "")
	if err := parser.ParseString(path, source, file); err != nil {
		return nil, fmt.Errorf("failed parse pipe4 file: %v: %w", path, err)
	}
	// file.PostParse()
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

func GetBnf() string {
	return parser.String()
}

func Parser() *participle.Parser {
	return parser
}
