package internal

import "github.com/pipe4/lang/pipe4/parser"

type structExpr struct {
	fields []statement

	exists bool
}

func structExprFromParser(s *parser.Statement) structExpr {
	return structExpr{
		fields: statementsFromParser(s.GetStruct()),
		exists: s.HasStruct(),
	}
}
