package parser

import (
	"fmt"

	"github.com/pipe4/lang/pipe4/ast"
)

type Statements []Statement

type Statement struct {
	Comment      *Comment      `parser:"( @(LineComment | BlockComment+)"  json:"Comment,omitempty"`
	SingleImport *SingleImport `parser:"| @@" json:"SingleImport,omitempty"`
	BlockImport  *BlockImport  `parser:"| @@" json:"BlockImport,omitempty"`
	Type         *Type         `parser:") | (@@ " json:"Type,omitempty"`

	Default *Type `parser:"('=' @@)? )" json:"Default,omitempty"`

	Meta `json:"-"`
}

func (c Statement) AstNode() ([]ast.Node, error) {
	switch {
	case c.Comment != nil:
		node := c.Comment.AstNode()
		if node == nil {
			return nil, nil
		}
		return []ast.Node{*node}, nil
	case c.SingleImport != nil:
		node := c.SingleImport.AstNode()
		if node == nil {
			return nil, nil
		}
		return []ast.Node{*node}, nil
	case c.BlockImport != nil:
		return c.BlockImport.AstNode(), nil
	}
	if c.Type == nil {
		return nil, fmt.Errorf("%v: empty statement, expected some fields present: %+v", c.Meta.Pos, c)
	}

	nodeType, err := c.Type.AstType()
	if err != nil {
		return nil, err
	}
	if nodeType == nil {
		return nil, nil
	}

	node := ast.Node{}

	var nodeDefault *ast.Type
	if c.Default != nil {
		if nodeDefault, err = c.Default.AstType(); err != nil {
			return nil, fmt.Errorf("failed to get default value for %v: %w", c.Meta.Pos, err)
		}
		if nodeDefault != nil {
			node.Default = *nodeDefault
		}
	}
	if nodeType.BodyOneOf == ast.BodyVoid {
		if nodeDefault != nil && len(nodeType.Args) == 0 {
			node.Ident = nodeType.Ident
			return []ast.Node{node}, nil
		}
		node.Type = *nodeType
		return []ast.Node{node}, nil
	}
	node.Ident = nodeType.Ident
	nodeType.Ident = ast.Ident{}
	if nodeType.BodyOneOf == ast.BodyType && len(nodeType.Args) == 0 {
		node.Type = *nodeType.Type
	} else {
		node.Type = *nodeType
	}
	return []ast.Node{node}, nil
}
func (c Statements) AstNode() ([]ast.Node, error) {
	var nodes []ast.Node
	for _, s := range c {
		subNodes, err := s.AstNode()
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, subNodes...)
	}
	return nodes, nil
}
