package ast

type Meta struct {
	Filename string
	Position MetaPosition
}

type MetaPosition struct {
	FromRow int
	FromCol int
	ToRow   int
	ToCol   int
}
