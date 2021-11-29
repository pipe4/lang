package parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/alecthomas/participle/v2/lexer"
	"github.com/google/go-cmp/cmp"
	"github.com/kr/pretty"
	"github.com/pipe4/lang/pipe4/ast"
	"github.com/stretchr/testify/assert"
)

func testLexer(t *testing.T, path string) {
	if runtime.GOOS == "windows" {
		t.Skipf("\\r not works on windows")
	}
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
}

func testParseFile(t *testing.T, path string) (ast *File) {
	t.Run("parse", func(t *testing.T) {
		got, err := ParseFile(path)
		if err != nil {
			assert.NoError(t, err, "error while parsing file")
		}
		// ctx := WalkCtx{Down: func(ctx WalkCtx) {
		// 	ctx.
		// }}
		// got.Statements.Walk(ctx)
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
		f := &File{}
		if err := f.FromYaml(yamlAst); err != nil {
			t.Errorf("failed parse ast from yaml: %+v", err)
			return
		}
		out = f
		t.Logf("file:\n%# v", pretty.Formatter(f))
	})
	return
}

func testEqualAst(t *testing.T, got *File, want *File) {
	t.Run("parser.File/eq/struct", func(t *testing.T) {
		if !cmp.Equal(want, got) {
			t.Errorf("AST not match\n%v", cmp.Diff(want, got))
			return
		}
	})
}
func testEqualNodeListAst(t *testing.T, got ast.NodeList, want ast.NodeList) {
	t.Run("ast.NodeList/eq/struct", func(t *testing.T) {
		if !cmp.Equal(want, got) {
			t.Errorf("AST not match\n%v", cmp.Diff(want, got))
			return
		}
	})
}
func testEqualYamlAst(t *testing.T, name string, got string, want string) {
	t.Run(name, func(t *testing.T) {
		assert.YAMLEqf(t, got, want, "ast not match")
	})
}
func testToAstNodeList(t *testing.T, name string, ast *File) (nodes ast.NodeList) {
	t.Run(name, func(t *testing.T) {
		got, err := ast.Statements.AstNode()
		if err != nil {
			t.Errorf("failed File to ast.NodeList: %v", err)
			return
		}
		nodes = got
		t.Logf("ast.NodeList:\n%# v", pretty.Formatter(got))
	})
	return
}
func testAstNodeListToYaml(t *testing.T, name string, nodes ast.NodeList) (yamlAst string) {
	t.Run(name, func(t *testing.T) {
		gotYaml, err := nodes.ToYaml()
		if err != nil {
			t.Errorf("failed to yaml ast: %v", err)
			return
		}
		fmt.Printf("======= AST: ========\n%v\n===========\n\n", gotYaml)
		yamlAst = gotYaml
	})
	return
}
func testFromYamlNodeList(t *testing.T, name string, yamlAst string) (out ast.NodeList) {
	t.Run(name, func(t *testing.T) {
		fmt.Printf("======= AST: ========\n%v\n===========\n\n", yamlAst)
		l := ast.NodeList{}
		if err := l.FromYaml(yamlAst); err != nil {
			t.Errorf("failed parse ast from yaml: %+v", err)
			return
		}
		out = l
		t.Logf("file:\n%# v", pretty.Formatter(l))
	})
	return
}
func testAst(t *testing.T, path string) {
	path = strings.ReplaceAll(path, "\\", "/")
	t.Run(path, func(t *testing.T) {
		t.Run("lexer", func(t *testing.T) {
			testLexer(t, path)
		})
		got := testParseFile(t, path)
		if got == nil {
			return
		}

		wantYaml := ""
		wantNodeListYaml := ""

		for i := 0; i < len(got.Statements); i++ {
			s := got.Statements[i]
			if s.Comment == nil {
				continue
			}
			if s.Comment.Tags.GetString("lang") != "yaml" {
				continue
			}
			if s.Comment.Tags.GetString("test") == "ast" {
				wantYaml = got.Statements[i].Comment.Text
				got.Statements = append(got.Statements[:i], got.Statements[i+1:]...)
				i--
			} else if s.Comment.Tags.GetString("test") == "ast.NodeList" {
				wantNodeListYaml = got.Statements[i].Comment.Text
				got.Statements = append(got.Statements[:i], got.Statements[i+1:]...)
				i--
			}
		}
		gotYaml := testToYaml(t, "parser.File[S]/to/yaml", got)
		got = testFromYaml(t, "parser.File[S]/from/yaml", gotYaml)
		if wantYaml != "" {
			want := testFromYaml(t, "parser.File[T]/from/yaml", wantYaml)
			testEqualYamlAst(t, "parser.File/eq/yaml", gotYaml, wantYaml)
			testEqualAst(t, got, want)
		}
		nodeList := testToAstNodeList(t, "parser.File/to/ast.NodeList", got)

		gotYaml = testAstNodeListToYaml(t, "ast.NodeList[S]/to/yaml", nodeList)
		gotNodeList := testFromYamlNodeList(t, "ast.NodeList[S]/from/yaml", gotYaml)

		if wantNodeListYaml != "" {
			want := testFromYamlNodeList(t, "ast.NodeList[T]/from/yaml", wantNodeListYaml)
			testEqualYamlAst(t, "ast.NodeList[T]/eq/yaml", gotYaml, wantNodeListYaml)
			testEqualNodeListAst(t, gotNodeList, want)
		}
	})
}

func TestExamplesAst(t *testing.T) {
	err := filepath.Walk("examples",
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
