package ast

type Ident interface {
	Path() string
	Package() Package
}

type Package interface {
	Path() string
	Module() Module
}

type Module interface {
	URL() string
	Version() Version
	Hash() Hash
}

type Version interface {
	String()
}

type Hash interface {
	String()
}
