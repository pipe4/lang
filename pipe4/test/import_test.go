package test

import (
	"testing"

	"github.com/pipe4/lang/pipe4"
)

var importAst = &pipe4.File{ImportBlock: pipe4.ImportBlock{
	Imports: []pipe4.Import{
		{Name: "", Path: `"github.com/pipe4/lang/pipe4"`},
		{Name: "local-rename", Path: `"github.com/pipe4/lang/asd"`},
		{Name: "pipe4", Path: `"github.com/pipe4/lang/pipe4"`},
		{Name: "", Path: `"github.com/stretchr/testify/require"`},
	},
}}

func TestImport(t *testing.T) {
	testAst(t, "./import.pipe4", importAst)
}
