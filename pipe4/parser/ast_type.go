package parser

import (
	"fmt"

	"github.com/pipe4/lang/pipe4/ast"
)

type Type struct {
	Args *Args `parser:"@@?" json:"Args,omitempty"`

	IdentWithType *IdentWithType `parser:"(@@" json:"IdentWithType,omitempty"`
	IdentWithArgs *IdentWithArgs `parser:"| @@" json:"IdentWithArgs,omitempty"`
	Ident         *Ident         `parser:"| @@" json:"Ident,omitempty"`
	Const         *Const         `parser:"| @@" json:"Const,omitempty"`
	Struct        *Struct        `parser:"| @@)" json:"Struct,omitempty"`

	Meta `json:"-"`
}

type IdentWithType struct {
	Ident Ident `parser:"@@" json:"Ident,omitempty"`
	Type  Type  `parser:"@@" json:"Type,omitempty"`

	Meta `json:"-"`
}

type IdentWithArgs struct {
	Ident Ident `parser:"@@" json:"Type,omitempty"`
	Args  Args  `parser:"@@" json:"Args,omitempty"`

	Meta `json:"-"`
}

type Ident struct {
	Path string `parser:"@Ident" json:"Path,omitempty"`

	Meta `json:"-"`
}

type Struct struct {
	Statements Statements `parser:"'{' EOS* (@@ ((EOS|',')+ @@)*)? EOS* '}'" json:"Statements,omitempty"`

	Meta `json:"-"`
}
type Args struct {
	Statements Statements `parser:"'(' EOS* (@@ ((EOS|',')+ @@)*)? EOS* ')'" json:"Statements,omitempty"`

	Meta `json:"-"`
}

func (c Struct) AstNode() ([]ast.Node, error) {
	return c.Statements.AstNode()
}
func (c Args) AstNode() ([]ast.Node, error) {
	return c.Statements.AstNode()
}

func (c Type) AstType() (*ast.Type, error) {
	t := &ast.Type{}
	if c.Args != nil {
		if args, err := c.Args.AstNode(); err != nil {
			return nil, err
		} else if len(args) > 0 {
			t.Args = args
		}
	}
	switch {
	case c.IdentWithArgs != nil:
		identWithArgs := &ast.Type{
			Ident: ast.Ident{
				Name: c.IdentWithArgs.Ident.Path,
			},
		}
		if args, err := c.IdentWithArgs.Args.AstNode(); err != nil {
			return nil, err
		} else if len(args) > 0 {
			identWithArgs.Args = args
		}
		if t.Args == nil {
			return identWithArgs, nil
		}
		if identWithArgs.Args == nil {
			t.Ident = identWithArgs.Ident
			return t, nil
		}
		t.Type = identWithArgs
		t.BodyOneOf = ast.BodyType
		return t, nil
	case c.IdentWithType != nil:
		t.Ident.Name = c.IdentWithType.Ident.Path
		subType, err := c.IdentWithType.Type.AstType()
		if err != nil {
			return nil, err
		}
		if subType.Ident.Name != "" || len(t.Args) > 0 {
			t.Type = subType
			t.BodyOneOf = ast.BodyType
			return t, nil
		}
		subType.Ident = t.Ident
		return subType, nil
	case c.Ident != nil:
		t.Ident.Name = c.Ident.Path
		return t, nil
	case c.Const != nil:
		subType, err := c.Const.AstType()
		if err != nil {
			return nil, err
		}
		subType.Ident = t.Ident
		subType.Args = t.Args
		return subType, nil
	case c.Struct != nil:
		subNodes, err := c.Struct.AstNode()
		if err != nil {
			return nil, err
		}
		t.BodyOneOf = ast.BodyStruct
		t.Struct = subNodes
		return t, nil
	default:
		return nil, fmt.Errorf("failed to get Type.AstType(), expected one of fields to set, but no one found in %v: %+v", c.Pos, c)
	}
}
