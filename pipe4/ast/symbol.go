package ast

import "strings"

type Symbol struct {
	Name  string
	Scope Scope
	Ident Ident
}

func (i Symbol) MarshalJSON() ([]byte, error) {
	ident := i.Ident.FullName()
	scopeIdent := i.Scope.Ident
	if strings.HasPrefix(i.Ident.URI, scopeIdent.URI) {
		ident = strings.TrimPrefix(ident, scopeIdent.FullName()+".")
	}
	return []byte(`"` + ident + `"`), nil
}

func (i *Symbol) UnmarshalJSON(value []byte) error {
	ident := strings.Trim(string(value), `"`)
	if !strings.ContainsRune(trimmedIdent, '.') {

	}
	return nil
}
