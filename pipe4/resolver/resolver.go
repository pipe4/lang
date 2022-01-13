package resolver

import (
	"errors"
	"fmt"
	"strings"

	"github.com/pipe4/lang/go/loader"
	"github.com/pipe4/lang/pipe4/ast"
)

var (
	Unimplemented = errors.New("unimplemented")
)

type Resolver struct {
	cache map[string]ast.Node
}

func NewResolver() *Resolver {
	return &Resolver{
		cache: make(map[string]ast.Node),
	}
}

func (r *Resolver) Resolve(ident ast.Ident) (*ast.Node, error) {
	if node, ok := r.cache[ident.GetURI()]; ok {
		return &node, nil
	}
	if strings.HasPrefix(ident.GetImportURI(), "github.com/pipe4/lang/go") {
		node, err := loader.Resolve(ident)
		if err != nil {
			return nil, err
		}
		r.cache[ident.GetURI()] = *node
	}
	return nil, fmt.Errorf("%w: %v", Unimplemented, ident.GetURI())
}
