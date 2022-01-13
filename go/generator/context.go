package generator

import "github.com/pipe4/lang/pipe4/ast"

type Context struct {
	graph ast.Node
	root  string

	leftPad int
	// scope   io.Writer
	body    []string
	imports []string
	// prefix string
}
