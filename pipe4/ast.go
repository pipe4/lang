package pipe4

type File struct {
	ImportBlock ImportBlock `@@`
}

type ImportBlock struct {
	Imports []Import `"import" ("(" @@* ")" | @@)`
}

type Import struct {
	Name string `@Ident?`
	Path string `@String`
}
