package ast

type Program struct {
	nodes []Node
}

// type URI string

type Pointer int64

type Node struct {
	pointer Pointer
	name    string
	path    []Pointer

	relations []Relation
}

type Relation struct {
	Link Pointer

	BodyOneOf BodyOneOf
	Node      Pointer
	String    string
	Rational  Rational
	Bool      bool

	// relations []Relation
	//
	// // Struct    NodeList
	// Type Pointer
}

// HashMap.anyGoMap
//	[go-std.Map, HashMap.anyGoMap, ]
