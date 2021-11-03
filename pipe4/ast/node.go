package ast

import (
	"math/big"
)

type Node struct {
	Ident
	Comment string

	Type    Type
	Default Type
}

type Type struct {
	Ident
	Args []Node

	BodyType // one of
	String   string
	Rational big.Rat
	Bool     bool
	Struct   []Node
}

type BodyType uint8

const (
	Void BodyType = iota
	String
	Rational
	Bool
	Struct
)
