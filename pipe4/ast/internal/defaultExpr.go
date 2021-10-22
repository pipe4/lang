package internal

import "github.com/pipe4/lang/pipe4/parser"

type defaultExpr struct {
	constExpr constExpr
	typeExpr  typeExpr
}

func (d defaultExpr) exists() bool {
	return d.typeExpr.exists() || d.constExpr.exists()
}

func defaultExprFromParser(s *parser.Statement) defaultExpr {
	return defaultExpr{
		constExpr: constExprFromParser(s.Default),
		typeExpr:  typeExprFromParser(s.Default),
	}
}
