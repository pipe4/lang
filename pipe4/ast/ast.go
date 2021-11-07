package ast

type Node struct {
	Ident   Ident
	Comment Comment

	Type    Type
	Default Type
}

type NodeList []Node

type Type struct {
	Ident Ident
	Args  NodeList

	BodyOneOf BodyOneOf
	String    string
	Rational  Rational
	Bool      bool
	Struct    NodeList
	Type      *Type
}

type Comment struct {
	Text string   `json:"Text,omitempty"`
	Tags NodeList `json:"Tags,omitempty"`
}

type BodyOneOf string

const (
	BodyVoid     BodyOneOf = ""
	BodyString   BodyOneOf = "string"
	BodyRational BodyOneOf = "rational"
	BodyBool     BodyOneOf = "bool"
	BodyStruct   BodyOneOf = "struct"
	BodyType     BodyOneOf = "type"
)
