package parser

import (
	"github.com/alecthomas/participle/v2/lexer"
)

//
// type Statement2 struct {
// 	// Imports               1  2         3                -3 -2  -1'
// 	Import []Import `parser:"(  ('import' (@@ | '(' @@* ')' )  )+  )" yaml:"Import,omitempty"`
//
// 	// Comment                  1  2'                        -2'
// 	Comment *Comment `parser:"| ( @(LineComment | BlockComment)?"  yaml:"Comment,omitempty"`
//
// 	// Definition        2' 2
// 	Name string `parser:"(  ( @Ident?" yaml:"Name,omitempty"`
//
// 	// Constant             3 4
// 	String *string `parser:"( (@String" yaml:"String,omitempty"`
// 	//                     5               -5
// 	Bool *Bool `parser:"| @Bool" yaml:"Bool,omitempty"`
// 	//                        -4
// 	Number *Rat `parser:"| @Rat)" yaml:"Number,omitempty"`
//
// 	// Type                4
// 	Type string `parser:"| ( @Ident?" yaml:"Type,omitempty"`
// 	//                          5     6   7     -7 -6     -5
// 	Props *[]Statement `parser:"( '(' (@@ (',' @@) *)? ')' )?" yaml:"Props,omitempty"`
// 	// //                                 5            -5  -4 -3 -2  -1
// 	// Struct *[]StructStatement `parser:"( '{' @@* '}' )?  )  )  )?  )!" yaml:"Struct,omitempty"`
//
// 	//                           5            -5  -4 -3
// 	Struct *[]Statement `parser:"( '{' @@* '}' )?  )! )?" yaml:"Struct,omitempty"`
//
// 	// Defaults                 3       -3  -2 -2' -1
// 	Default *Statement `parser:"( '=' @@ )?  )! )?  )!" yaml:"Default,omitempty"`
//
// 	Pos    lexer.Position `parser:"" yaml:"-"`
// 	EndPos lexer.Position `parser:"" yaml:"-"`
// 	// Tokens []lexer.Token `parser:"" yaml:"-"`
// }

type Definition struct {
	Ident Ident `parser:"@@" yaml:"Ident,omitempty"`
	Type  Type  `parser:"@@" yaml:"Type,omitempty"`

	Default *Type `parser:"('=' @@)?" yaml:"Default,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

func (s *Definition) Walk(ctx WalkCtx) {
	if s != nil {
		s.Type.Walk(ctx)
		s.Default.Walk(ctx)
	}
}

type DefaultDefinitionWithoutType struct {
	Ident   Ident `parser:"@@" yaml:"Ident,omitempty"`
	Default Type  `parser:"'=' @@" yaml:"Default,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

func (s *DefaultDefinitionWithoutType) Walk(ctx WalkCtx) {
	if s != nil {
		s.Default.Walk(ctx)
	}
}

type Type struct {
	TypeInstantiation *TypeInstantiation `parser:"@@" yaml:"TypeInstantiation,omitempty"`
	JustType          *JustType          `parser:"| @@" yaml:"JustType,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

func (s *Type) Walk(ctx WalkCtx) {
	if s != nil {
		s.TypeInstantiation.Walk(ctx)
		s.JustType.Walk(ctx)
	}
}

type JustType struct {
	Ident Ident `parser:"@@" yaml:"Ident,omitempty"`
	Args  Args  `parser:"@@?" yaml:"Args,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

func (s *JustType) Walk(ctx WalkCtx) {
	if s != nil {
		s.Args.Walk(ctx)
	}
}

type Ident struct {
	Path string `parser:"@Ident" yaml:"Path,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

type TypeInstantiation struct {
	Ident Ident `parser:"@@?" yaml:"Ident,omitempty"`

	BodyWithArgs *BodyWithArgs `parser:"@@" yaml:"BodyWithArgs,omitempty"`
	Body         *Body         `parser:"| @@" yaml:"Body,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

func (s *TypeInstantiation) Walk(ctx WalkCtx) {
	if s != nil {
		s.BodyWithArgs.Walk(ctx)
		s.Body.Walk(ctx)
	}
}

type BodyWithArgs struct {
	Args Args `parser:"@@" yaml:"Args,omitempty"`
	Body Body `parser:"@@" yaml:"Body,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

func (s *BodyWithArgs) Walk(ctx WalkCtx) {
	if s != nil {
		s.Args.Walk(ctx)
		s.Body.Walk(ctx)
	}
}

type Body struct {
	Struct *Struct `parser:"@@" yaml:"Struct,omitempty"`
	Const  *Const  `parser:"| @@" yaml:"Const,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

func (s *Body) Walk(ctx WalkCtx) {
	if s != nil {
		s.Struct.Walk(ctx)
	}
}

type Struct struct {
	Statements Statements `parser:"'{' @@* '}'" yaml:"Statements,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

func (s *Struct) Walk(ctx WalkCtx) {
	if s != nil {
		s.Statements.Walk(ctx)
	}
}

type Args struct {
	Statements Statements `parser:"'(' @@* ')'" yaml:"Statements,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

func (s *Args) Walk(ctx WalkCtx) {
	if s != nil {
		s.Statements.Walk(ctx)
	}
}
