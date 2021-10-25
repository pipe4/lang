package ast

type Node struct {
	Comment string

	Ident   Ident
	Type    Type
	Default Type
}

type Type struct {
	Const Const

	Ident     Ident
	Arguments []Node
	Body      []Node
}

type Comment struct {
	Text string
}
