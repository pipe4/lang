package ast

import "crypto"

type Ident struct {
	// Name - ident path relative to package, for example: Ident/Path
	Name string
	// Filename relative to package directory, for example: ident.go
	Filename string
	// Package path relative to module, for example: pipe4/ast
	Package string

	Module
}

type Module struct {
	// URI that exactly identify module, for example: github.com/pipe4/lang
	URI string

	Version Version
	Hash    Hash
	Tags    []string
}

type Version struct {
	Major int
	Minor int
	Patch int
}

type Hash struct {
	String string
	Alg    crypto.Hash
}
