package loader

import (
	"testing"

	"github.com/pipe4/lang/pipe4/ast"
	"github.com/stretchr/testify/require"
)

func TestImport(t *testing.T) {
	ret, err := Resolve(ast.Ident{Name: "Printf", Package: "lang/go/log"})
	require.NoError(t, err)
	yamlString, err := ast.NodeList{*ret}.ToYaml()
	require.NoError(t, err)
	t.Log("\n", yamlString)
}
