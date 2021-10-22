package internal

import "github.com/pipe4/lang/pipe4/parser"

type statement struct {
	commentExpr commentExpr

	declarationName declarationName // local-name

	constExpr constExpr
	typeExpr  typeExpr

	defaultExpr defaultExpr
}

type declarationName struct {
	name string
}

func (d statement) HasName() bool {
	return d.declarationName.name != ""
}

func (d statement) HasConst() bool {
	return d.constExpr.exists()
}

func (d statement) HasType() bool {
	return d.typeExpr.exists()
}

func (d statement) HasDefault() bool {
	return d.defaultExpr.exists()
}

func (d statement) HasComment() bool {
	return d.commentExpr.exists
}

func (d statement) exists() bool {
	return d.HasComment() || d.HasName() || d.HasConst() || d.HasType()
}

func statementFromParser(s *parser.Statement) statement {
	if s == nil {
		return statement{}
	}
	return statement{
		commentExpr:     commentExprFromParser(s),
		declarationName: declarationName{name: s.Name},
		constExpr:       constExprFromParser(s),
		typeExpr:        typeExprFromParser(s),
		defaultExpr:     defaultExprFromParser(s),
	}
}

func statementsFromParser(s []parser.Statement) []statement {
	var statements []statement
	for i := range s {
		field := statementFromParser(&s[i])
		if !field.exists() {
			continue
		}
		statements = append(statements, field)
	}
	return statements
}
