package ast

import "crypto"

type Ident struct {
	// Name of node relative to scope, for example: Name
	Name string
	// Scope of node relative to package, for example Ident
	Scope string
	// Filename relative to package directory, for example: ident.go
	Filename string
	// Package path relative to module, for example: pipe4/ast
	Package string

	Module *Module
}

type Module struct {
	// URI that exactly identify module, for example: github.com/pipe4/lang
	URI     string
	Version Version
}

type Version struct {
	Major int
	Minor int
	Patch int

	Ref  string
	Tags []string

	Hash crypto.Hash
	Sum  string
}
