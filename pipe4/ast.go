package pipe4

type File struct {
	ImportBlock ImportBlock `parser:"@@"`
	Statements  []Statement `parser:"@@*"`
}

type ImportBlock struct {
	Imports []Import `parser:"('import' (@@ | '(' @@* ')'))*"`
}

type Import struct {
	Name string `parser:"@Ident?"`
	Path string `parser:"@String"`
}

type Statement struct {
	ShortDeclaration *ShortDeclaration `parser:"@@"`
	Declaration      *Declaration      `parser:"| @@"`
}

type Declaration struct {
	TypeFamily string `parser:"@Ident"`
	Name       string `parser:"@Ident"`
	Type       Type   `parser:"@@"`
}

type ShortDeclaration struct {
	Name string `parser:"@Ident"`
	Type Type   `parser:"':=' @@"`
}

type Type struct {
	Instantiation *Instantiation `parser:"@@"`
	Path          string         `parser:"| @Ident"`
	Constant      *Constant      `parser:"| @@"`
}

type Instantiation struct {
	Path string  `parser:"@Ident?"`
	Body Pattern `parser:"@@"`
}

type Pattern struct {
	Struct   *Struct   `parser:"@@"`
	Function *Function `parser:"| @@"`
}

type Struct struct {
	Fields []StructField `parser:"'{' @@* '}'"`
}

type Function struct {
	Arguments Struct      `parser:"'(' @@ ')'"`
	Body      []Statement `parser:"'{' @@* '}'"`
}

type Constant struct {
	String string `parser:"@String"`
	Float  string `parser:"| @Float"`
	Int    string `parser:"| @Int"`
}

type StructField struct {
	Name string `parser:"@Ident"`
	Type Type   `parser:"@@"`
}
