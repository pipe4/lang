package parser

type Import struct {
	Name string `parser:"@Ident?" yaml:"Name,omitempty"`
	URL  string `parser:"@String" yaml:"URL,omitempty"`

	Meta `yaml:"-"`
}
type SingleImport struct {
	Import Import `parser:"'import' @@ " yaml:"Import,omitempty"`

	Meta `yaml:"-"`
}
type BlockImport struct {
	Imports []Import `parser:"'import' '(' EOS* (@@ EOS)*  @@? EOS* ')'" yaml:"Imports,omitempty"`

	Meta `yaml:"-"`
}
