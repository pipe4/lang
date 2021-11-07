package parser

import "github.com/alecthomas/participle/v2/lexer"

type Statements []Statement

type Meta struct {
	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

type File struct {
	Name       string     `parser:"" yaml:"-"`
	Statements Statements `parser:"EOS* (@@ EOS+)*" yaml:"Statements,omitempty"`
}

type Statement struct {
	Comment      *Comment      `parser:"( @(LineComment | BlockComment+)"  yaml:"Comment,omitempty"`
	SingleImport *SingleImport `parser:"| @@" yaml:"SingleImport,omitempty"`
	BlockImport  *BlockImport  `parser:"| @@" yaml:"BlockImport,omitempty"`
	Type         *Type         `parser:"| @@ )" yaml:"Type,omitempty"`

	Default *Type `parser:"('=' @@)?" yaml:"Default,omitempty"`

	Meta `yaml:"-"`
}

type WalkCtx struct {
	Statement
	Parent Statement

	Down, Up func(ctx WalkCtx)
}

func (s *Statement) Walk(ctx WalkCtx) {
	if s == nil {
		return
	}
	ctx.Parent = *s
	s.Type.Walk(ctx)
}

func (s Statements) Walk(ctx WalkCtx) {
	for i := 0; i < len(s); i++ {
		ctx.Statement = s[i]
		if ctx.Down != nil {
			ctx.Down(ctx)
		}
		s[i].Walk(ctx)
		if ctx.Up != nil {
			ctx.Up(ctx)
		}
	}
}
