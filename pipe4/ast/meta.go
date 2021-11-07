package ast

type Meta struct {
	Filename string
	Position PositionRange
}

type PositionRange struct {
	From Position
	To   Position
}
type Position struct {
	Row int
	Col int
}
