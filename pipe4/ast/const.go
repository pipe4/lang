package ast

import (
	"math/big"
)

type ConstType uint8

const (
	ConstVoid = iota
	ConstString
	ConstRational
	ConstBool
)

type Const struct {
	String   string
	Rational big.Rat
	Bool     bool

	ConstType
}

// func constExprFromParser(s *parser.Statement) Const {
// 	if s == nil {
// 		return Const{}
// 	}
// 	c := Const{}
// 	switch {
// 	case s.String != nil:
// 		c.String = String{value: *s.String}
// 		c.ConstVariant = ConstString
// 	case s.Number != nil:
// 		c.Rational = Rational{value: s.Number.Rat}
// 		c.ConstVariant = ConstRational
// 	case s.Bool != nil:
// 		c.Bool = Bool{value: bool(*s.Bool)}
// 		c.ConstVariant = ConstBool
// 	}
// 	return c
// }
