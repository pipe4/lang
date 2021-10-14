package pipe4

type File struct {
	ImportBlock ImportBlock `parser:"@@" yaml:",omitempty"`
	Statements  []Statement `parser:"@@*" yaml:",omitempty"`
}

type ImportBlock struct {
	Imports []Import `parser:"('import' (@@ | '(' @@* ')'))*" yaml:",omitempty"`
}

type Import struct {
	Name string `parser:"@Ident?" yaml:",omitempty"`
	Path string `parser:"@String" yaml:",omitempty"`
}

type Statement struct {
	ShortDeclaration *ShortDeclaration `parser:"@@" yaml:",omitempty"`
	Declaration      *Declaration      `parser:"| @@" yaml:",omitempty"`
}

type Declaration struct {
	TypeFamily string `parser:"@Ident" yaml:",omitempty"`
	Name       string `parser:"@Ident" yaml:",omitempty"`
	Type       Type   `parser:"@@" yaml:",omitempty"`
}

type ShortDeclaration struct {
	Name string `parser:"@Ident" yaml:",omitempty"`
	Type Type   `parser:"':=' @@" yaml:",omitempty"`
}

type Type struct {
	Instantiation *Instantiation `parser:"@@" yaml:",omitempty"`
	Path          string         `parser:"| @Ident" yaml:",omitempty"`
	Constant      *Constant      `parser:"| @@" yaml:",omitempty"`
}

type Instantiation struct {
	Path string  `parser:"@Ident?" yaml:",omitempty"`
	Body Pattern `parser:"@@" yaml:",omitempty"`
}

type Pattern struct {
	Struct   *Struct   `parser:"@@" yaml:",omitempty"`
	Function *Function `parser:"| @@" yaml:",omitempty"`
}

type Struct struct {
	Fields []StructField `parser:"'{' @@* '}'" yaml:",omitempty"`
}

type Function struct {
	Arguments Struct      `parser:"'(' @@ ')'" yaml:",omitempty"`
	Body      []Statement `parser:"'{' @@* '}'" yaml:",omitempty"`
}

type Constant struct {
	String string `parser:"@String" yaml:",omitempty"`
	Float  string `parser:"| @Float" yaml:",omitempty"`
	Int    string `parser:"| @Int" yaml:",omitempty"`
}

type StructField struct {
	Name string `parser:"@Ident" yaml:",omitempty"`
	Type Type   `parser:"@@" yaml:",omitempty"`
}
