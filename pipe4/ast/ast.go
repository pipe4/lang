package ast

type Node struct {
	Ident   Ident   `json:"Ident,omitempty"`
	Comment Comment `json:"Comment,omitempty"`

	Type    Type `json:"Type,omitempty"`
	Default Type `json:"Default,omitempty"`
}

type NodeList []Node

type Type struct {
	Ident Ident    `json:"Ident,omitempty"`
	Args  NodeList `json:"Args,omitempty"`

	BodyOneOf BodyOneOf `json:"BodyOneOf,omitempty"`
	String    string    `json:"String,omitempty"`
	Rational  Rational  `json:"Rational,omitempty"`
	Bool      bool      `json:"Bool,omitempty"`
	Struct    NodeList  `json:"Struct,omitempty"`
	Type      *Type     `json:"Type,omitempty"`
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
