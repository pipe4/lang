package ast

import (
	"crypto"
	"path"
	"strings"
)

type Ident struct {
	// Name of node relative to scope, for example: Name
	Name string `json:"Name,omitempty"`
	// Scope of node relative to package, for example Ident
	Scope string `json:"Scope,omitempty"`
	// Filename relative to package directory, for example: ident.go
	Filename string `json:"Filename,omitempty"`
	// Package path relative to module, for example: pipe4/ast
	Package string `json:"Package,omitempty"`
	// Module
	Module *Module `json:"Module,omitempty"`
	// ImportURI
	ImportURI string `json:"ImportURI,omitempty"`
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
	URI     string  `json:"URI,omitempty"`
	Version Version `json:"Version,omitempty"`
}

type Version struct {
	Major int `json:"Major,omitempty"`
	Minor int `json:"Minor,omitempty"`
	Patch int `json:"Patch,omitempty"`

	Ref  string   `json:"Ref,omitempty"`
	Tags []string `json:"Tags,omitempty"`

	Hash crypto.Hash `json:"Hash,omitempty"`
	Sum  string      `json:"Sum,omitempty"`
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
