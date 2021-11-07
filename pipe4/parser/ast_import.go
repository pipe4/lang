package parser

import "github.com/pipe4/lang/pipe4/ast"

type Import struct {
	Name string `parser:"@Ident?" json:"Name,omitempty"`
	URL  string `parser:"@String" json:"URL,omitempty"`

	Meta `json:"-"`
}
type SingleImport struct {
	Import Import `parser:"'import' @@ " json:"Import,omitempty"`

	Meta `json:"-"`
}
type BlockImport struct {
	Imports []Import `parser:"'import' '(' EOS* (@@ EOS)*  @@? EOS* ')'" json:"Imports,omitempty"`

	Meta `json:"-"`
}

func (c Import) AstNode() *ast.Node {
	if c.URL == "" {
		return nil
	}
	ident := ast.Ident{Name: c.Name, ImportURI: c.URL}
	ident.Normalize()
	return &ast.Node{Ident: ident}
}

func (c BlockImport) AstNode() []ast.Node {
	var imports []ast.Node
	for _, i := range c.Imports {
		node := i.AstNode()
		if node != nil {
			imports = append(imports, *node)
		}
	}
	return imports
}

func (c SingleImport) AstNode() *ast.Node {
	return c.Import.AstNode()
}
