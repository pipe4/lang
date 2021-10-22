package parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/alecthomas/participle/v2/lexer"
	"github.com/google/go-cmp/cmp"
	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func testLexer(t *testing.T, path string) {
	t.Run("lexer", func(t *testing.T) {
		tokens, err := LexFile(path)
		if err != nil {
			assert.NoError(t, err, "error while lexing file")
			return
		}

		names := make(map[lexer.TokenType]string)
		for k, v := range Lexer().Symbols() {
			names[v] = k
		}
		t.Logf("tokens:\n")
		for _, token := range tokens {
			fmt.Printf("====== %v ======\n%v\n", names[token.Type], token.String())
		}
		t.Logf("\n")
	})
}

func testParseFile(t *testing.T, path string) (ast *File) {
	t.Run("parse", func(t *testing.T) {
		got, err := ParseFile(path)
		if err != nil {
			assert.NoError(t, err, "error while parsing file")
		}
		got.Walk(func(s StatementWithContext) {
			s.Pos = lexer.Position{}
			s.EndPos = lexer.Position{}
			// s.Tokens = nil
		})
		ast = got
		t.Logf("file:\n%# v", pretty.Formatter(got))
	})
	return
}
func testToYaml(t *testing.T, name string, ast *File) (yamlAst string) {
	t.Run(name, func(t *testing.T) {
		gotYaml, err := ast.ToYaml()
		if err != nil {
			t.Errorf("failed to yaml ast: %v", err)
			return
		}
		fmt.Printf("======= AST: ========\n%v\n===========\n\n", gotYaml)
		yamlAst = gotYaml
	})
	return
}
func testFromYaml(t *testing.T, name string, yamlAst string) (out *File) {
	t.Run(name, func(t *testing.T) {
		fmt.Printf("======= AST: ========\n%v\n===========\n\n", yamlAst)
		ast := &File{}
		if err := ast.FromYaml(yamlAst); err != nil {
			t.Errorf("failed parse ast from yaml: %+v", err)
			return
		}
		out = ast
		t.Logf("file:\n%# v", pretty.Formatter(ast))
	})
	return
}

func testEqualAst(t *testing.T, got *File, want *File) {
	t.Run("parsedAstEqualToTest", func(t *testing.T) {
		if !cmp.Equal(want, got) {
			t.Errorf("AST not match\n%v", cmp.Diff(want, got))
			return
		}
	})
}
func testEqualYamlAst(t *testing.T, got string, want string) {
	t.Run("parsedAstEqualToTestInYaml", func(t *testing.T) {
		assert.YAMLEqf(t, got, want, "ast not match")
	})
}
func testAst(t *testing.T, path string) {
	t.Run(path, func(t *testing.T) {
		testLexer(t, path)
		got := testParseFile(t, path)
		if got == nil {
			return
		}

		wantYaml := ""

		for i, s := range got.Statements {
			if s.Comment.GetTag("lang") != "yaml" {
				continue
			}
			if s.Comment.GetTag("test") != "ast" {
				continue
			}
			wantYaml = got.Statements[i].Comment.Text
			got.Statements[i].Comment = nil
			if reflect.DeepEqual(got.Statements[i], Statement{}) {
				got.Statements = append(got.Statements[:i], got.Statements[i+1:]...)
			}
		}
		gotYaml := testToYaml(t, "toYamlParsedAst", got)
		got = testFromYaml(t, "fromYamlParsedAst", gotYaml)
		if wantYaml == "" {
			return
		}
		want := testFromYaml(t, "fromYamlTestAst", wantYaml)
		testEqualYamlAst(t, gotYaml, wantYaml)
		testEqualAst(t, got, want)
	})
}

func TestExamplesAst(t *testing.T) {
	err := filepath.Walk("./examples",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			if !strings.HasSuffix(path, ".pipe4") {
				return nil
			}
			testAst(t, path)
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
