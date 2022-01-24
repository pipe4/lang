package ast

import "strings"

type Scope struct {
	ident  Ident
	parent *Scope

	declarations []Declaration

	filename string
}

type Declaration struct {
	Name  string
	Ident Ident
}

func (s *Scope) Resolve(ident Ident) Ident {
	if ident.IsComplete() {
		return ident
	}
	baseIdent := s.FindByBaseName(ident.BaseName())
	if baseIdent != "" {
		return ident.Rebase(baseIdent)
	}
	if strings.ContainsRune(ident.FullName(), '.') {
		return ident
	}
	return Ident(string(s.ident) + "." + ident.FullName())
}

func (s *Scope) FindByBaseName(basename string) Ident {
	for _, decl := range s.declarations {
		if decl.Name == basename {
			return decl.Ident
		}
	}
	if s.parent != nil {
		return s.parent.FindByBaseName(basename)
	}
	return ""
}
