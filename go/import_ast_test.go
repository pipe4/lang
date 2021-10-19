package lang_go

import (
	"fmt"
	"go/printer"
	"go/token"
	"os"
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/require"
)

func TestImportAst(t *testing.T) {
	stdAst, err := ImportAst("../std")
	require.NoError(t, err, "failed parse pipe4/lang/std/*.go files ast")
	fileSet := token.NewFileSet()
	fmt.Printf("Go Ast for pipe4/lang/std/*.go:\n=================\n")
	interfaceFile := stdAst["std"].Files["../std/interface.go"]
	fmt.Printf("ast:\n============\n%# v\n================\n", pretty.Formatter(interfaceFile))
	if err := printer.Fprint(os.Stdout, fileSet, interfaceFile); err != nil {
		t.Errorf("failed print ast to stdout: %+v", err)
	}
	fmt.Printf("\n=================\n")
}
