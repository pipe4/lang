package internal

import "github.com/pipe4/lang/pipe4/parser"

type typeExpr struct {
	baseType   baseType
	propsExpr  propsExpr
	structExpr structExpr
}

type baseType struct {
	path string
}

func (b baseType) exists() bool {
	return b.path != ""
}

func (t typeExpr) exists() bool {
	return t.baseType.exists() || t.propsExpr.exists || t.structExpr.exists
}

func typeExprFromParser(s *parser.Statement) typeExpr {
	if s == nil {
		return typeExpr{}
	}
	t := typeExpr{
		baseType:   baseType{path: s.Name},
		propsExpr:  propsExprFromParser(s),
		structExpr: structExprFromParser(s),
	}

	return t
}
