package ast

import (
	"crypto"
	"path"
	"strconv"
	"strings"
)

type Ident struct {
	// Name of node relative to scope, for example: Ident.Name
	Name string
	// // Scope of node relative to package, for example Ident
	// Scope string
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

func (i *Ident) String() string {
	return i.GoImport()
}

func (i *Ident) GoImport() string {
	var uri []string
	if i.Module != nil && i.Module.URI != "" {
		uri = append(uri, i.Module.URI)
	}
	if i.Package != "" {
		uri = append(uri, i.Package)
	}

	if i.Module != nil && i.Module.Version.Major > 1 {
		uri = append(uri, strconv.Itoa(i.Module.Version.Major))
	}
	return path.Join(uri...)
}

func (i Ident) WithModule(uri string) Ident {
	i.Module = &Module{URI: uri}
	return i
}
func (i Ident) WithPackage(Package string) Ident {
	i.Package = Package
	return i
}

func (i Ident) Match(with Ident) bool {
	return i.GetURI() == with.GetURI()
}

func (i Ident) GetImportURI() string {
	if i.ImportURI != "" {
		return i.ImportURI
	}
	uri := ""
	if i.Module != nil {
		uri = i.Module.URI
	}
	if i.Module != nil && i.Module.Version.Major > 1 {
		uri += "/" + strconv.Itoa(i.Module.Version.Major)
	}
	if i.Package != "" {
		uri += "/" + i.Package
	}
	return uri
}

func (i Ident) GetURI() string {
	uri := i.GetImportURI()
	if i.Name != "" {
		uri += "." + i.Name
	}
	return uri
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

func NewIdent(module string, Package string, name string) Ident {
	return Ident{Name: name}.WithModule(module).WithPackage(Package)
}
