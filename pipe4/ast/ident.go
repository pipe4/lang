package ast

import (
	"crypto"
	"path"
	"strings"
)

type Ident struct {
	// Name of node relative to scope, for example: Name
	Name string
	// Scope of node relative to package, for example Ident
	Scope string
	// Filename relative to package directory, for example: ident.go
	Filename string
	// Package path relative to module, for example: pipe4/ast
	Package string
	// Module
	Module *Module
	// ImportURI
	ImportURI string
}

func (i Ident) MarshalJSON() ([]byte, error) {
	if i.Name == "" {
		return []byte(`""`), nil
	}
	return []byte(`"` + i.Name + `"`), nil
}

func (i *Ident) UnmarshalJSON(value []byte) error {
	i.Name = strings.Trim(string(value), `"`)
	return nil
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

func (i *Ident) Normalize() {
	if i.Name == "" {
		namePath := i.Package
		if namePath == "" {
			namePath = i.ImportURI
		}
		if namePath != "" {
			_, i.Name = path.Split(namePath)
		}
	}
}
