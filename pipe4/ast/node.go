package ast

import "math/big"

type Node struct {
	Ident   Ident
	Comment string

	Type    Type
	Default Type
}
type Type struct {
	Ident Ident
	Args  []Node

	BodyOneOf BodyOneOf
	String    string
	Rational  big.Rat
	Bool      bool
	Struct    []Node
	Type      *Type
}

type BodyOneOf uint8

const (
	BodyVoid BodyOneOf = iota
	BodyString
	BodyRational
	BodyBool
	BodyStruct
	BodyType
)
