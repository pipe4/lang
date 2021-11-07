package parser

import (
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pipe4/lang/pipe4/parser"
	"github.com/stretchr/testify/require"
)

func TestImportAst(t *testing.T) {
	pkg := path.Join("examples", "structs")
	nodes, err := ImportAst(pkg)
	require.NoErrorf(t, err, "failed parse %v from go to pipe4", pkg)

	yamlNodes, err := nodes.ToYaml()
	require.NoErrorf(t, err, "failed print nodes to yaml %v", nodes)
	t.Logf("Nodes:\n%v\n===============================\n\n\n", yamlNodes)

	pipe4File, err := parser.ParseFile("./examples/structs/structs.pipe4")
	require.NoErrorf(t, err, "failed parse structs.pipe4")
	p4nodes, err := pipe4File.Statements.AstNode()
	require.NoErrorf(t, err, "failed get ast.Nodes from pipe4file: %v", pipe4File)

	if !cmp.Equal(p4nodes, nodes) {
		t.Errorf("AST not match\n%v", cmp.Diff(p4nodes, nodes))
		return
	}

	// stdPath := path.Join("..", "..", "std")
	// stdAst, err := ImportAst(stdPath)
	// require.NoError(t, err, "failed parse pipe4/lang/std/*.go files ast")
	// fileSet := token.NewFileSet()
	// fmt.Printf("Go Ast for pipe4/lang/std/*.go:\n=================\n")
	// ifacePath := path.Join(stdPath, "interface.go")
	// interfaceFile := stdAst["std"].Files[ifacePath]
	// fmt.Printf("ast:\n============\n%# v\n================\n", pretty.Formatter(interfaceFile))
	// if err := printer.Fprint(os.Stdout, fileSet, interfaceFile); err != nil {
	// 	t.Errorf("failed print ast to stdout: %+v", err)
	// }
	// fmt.Printf("\n=================\n")
}
