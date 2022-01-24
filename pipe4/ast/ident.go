package ast

import (
	"path"
	"strconv"
	"strings"
)

// Ident like "github.com/pipe4/lang/v1/pipe4/ast.Ident.URI"
type Ident string

func (i Ident) Name() string {
	names := strings.Split(i.FullName(), ".")
	return names[len(names)-1]
}

func (i Ident) BaseName() string {
	return strings.Split(i.FullName(), ".")[0]
}

func (i Ident) Rebase(base Ident) Ident {
	subName := strings.Join(strings.Split(i.FullName(), ".")[1:], ".")
	return Ident(string(base) + "." + subName)
}

func (i Ident) IsComplete() bool {
	return strings.ContainsRune(string(i), '/')
}

func (i Ident) FullName() string {
	paths := strings.Split(string(i), "/")
	return paths[len(paths)-1]
}

func (i Ident) MarshalJSON() ([]byte, error) {
	return []byte(`"` + i.URI + `"`), nil
}
func (i *Ident) UnmarshalJSON(value []byte) error {
	i.URI = strings.Trim(string(value), `"`)
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

// type Module struct {
// 	// URI that exactly identify module, for example: github.com/pipe4/lang
// 	// URI     string
// 	// Version Version
// }

// type Version struct {
// 	Major int
// 	Minor int
// 	Patch int
//
// 	Ref  string
// 	Tags []string
//
// 	Hash crypto.Hash
// 	Sum  string
// }

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
