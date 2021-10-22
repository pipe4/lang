package internal

import (
	"math/big"

	"github.com/pipe4/lang/pipe4/parser"
)

type constType int

const (
	constNotExists = iota
	constString
	constRat
	constBool
)

type constExpr struct {
	string stringExpr
	rat    ratExpr
	bool   boolExpr

	constType constType
}

func (c constExpr) exists() bool {
	return c.constType != constNotExists
}

type stringExpr struct {
	value string
}

type ratExpr struct {
	value big.Rat
}

type boolExpr struct {
	value bool
}

func constExprFromParser(s *parser.Statement) constExpr {
	if s == nil {
		return constExpr{}
	}
	c := constExpr{}
	switch {
	case s.String != nil:
		c.string = stringExpr{value: *s.String}
		c.constType = constString
	case s.Number != nil:
		c.rat = ratExpr{value: s.Number.Rat}
		c.constType = constRat
	case s.Bool != nil:
		c.bool = boolExpr{value: bool(*s.Bool)}
		c.constType = constBool
	}
	return c
}
