package parser

import (
	"fmt"
	"go/ast"
	"log"

	pipe4ast "github.com/pipe4/lang/pipe4/ast"
)

type Ctx struct {
	Package StructCtx
	pkg     *ast.Package
}

type StructCtx struct {
	*pipe4ast.NodeList

	err string
}

type NodeCtx struct {
	*pipe4ast.Node

	err string
}
type TypeCtx struct {
	*pipe4ast.Type

	err string
}

func (c *StructCtx) NewNode() *NodeCtx {
	*c.NodeList = append(*c.NodeList, pipe4ast.Node{})
	return &NodeCtx{
		Node: &(*c.NodeList)[len(*c.NodeList)-1],
	}
}

func (c *StructCtx) Visit(n ast.Node) (w ast.Visitor) {
	switch x := n.(type) {
	case *ast.Package:
		return c
	case *ast.File:
		return c
	case *ast.GenDecl:
		return c
	case *ast.TypeSpec:
		if !x.Name.IsExported() {
			return nil
		}
		o := c.NewNode()
		o.Ident.Name = x.Name.String()
		return o
	default:
		log.Printf("%v: %s unimplemented", n, x)
		return nil
	}
}
func (c *NodeCtx) StructType() *StructCtx {
	c.Type.BodyOneOf = pipe4ast.BodyStruct
	return &StructCtx{
		NodeList: &c.Type.Struct,
	}
}
func (c *NodeCtx) TypeCtx() *TypeCtx {
	return &TypeCtx{
		Type: &c.Type,
	}
}
func (c *NodeCtx) Visit(n ast.Node) (w ast.Visitor) {
	switch x := n.(type) {
	case *ast.StructType:
		s := c.StructType()
		for _, f := range x.Fields.List {
			o := s.NewNode()
			o.Visit(f)
		}
		return nil
	case *ast.Field:
		if len(x.Names) != 1 {
			c.err = fmt.Sprintf("unimplemented multiple names for field: %+v", x.Names)
			return nil
		}
		c.Ident.Name = x.Names[0].Name
		t := c.TypeCtx()
		t.Visit(x.Type)
		return nil
	default:
		log.Printf("%v: %s unimplemented", n, x)
		return nil
	}
}
func (c *TypeCtx) Visit(n ast.Node) (w ast.Visitor) {
	switch x := n.(type) {
	case *ast.Ident:
		// c.BodyOneOf = pipe4ast.BodyVoid
		c.Ident.Name = x.Name
		return nil
	default:
		log.Printf("%v: %s unimplemented", n, x)
		return nil
	}
}

// type PackageVisitor struct {
// 	*Ctx
//
// 	c *pipe4ast.Node
// 	t *pipe4ast.Type
// }

// func (c *PackageVisitor) Visit(n ast.Node) (w ast.Visitor) {
// 	switch x := n.(type) {
// 	case *ast.TypeSpec:
//
// 	case *ast.StructType:
//
// 	default:
// 		log.Printf("%s unimplemented", x)
// 	}
//
// 	return nil
// }
