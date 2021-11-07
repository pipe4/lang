package parser

type Type struct {
	Args Args `parser:"@@?" yaml:"Args,omitempty"`

	IdentWithType *IdentWithType `parser:"(@@" yaml:"IdentWithType,omitempty"`
	IdentWithArgs *IdentWithArgs `parser:"| @@" yaml:"IdentWithArgs,omitempty"`
	Ident         *Ident         `parser:"| @@" yaml:"Ident,omitempty"`
	Const         *Const         `parser:"| @@" yaml:"Const,omitempty"`
	Struct        *Struct        `parser:"| @@)" yaml:"Struct,omitempty"`

	Meta `yaml:"-"`
}

func (s *Type) Walk(ctx WalkCtx) {
	// if s != nil {
	// 	// s.Struct.Walk(ctx)
	// 	// s.TypeWithLeftArgs.Walk(ctx)
	// 	// s.TypeWithRightArgs.Walk(ctx)
	// }
}

type IdentWithType struct {
	Ident Ident `parser:"@@" yaml:"Ident,omitempty"`
	Type  Type  `parser:"@@" yaml:"Type,omitempty"`

	Meta `yaml:"-"`
}

type IdentWithArgs struct {
	Ident Ident `parser:"@@" yaml:"Type,omitempty"`
	Args  Args  `parser:"@@" yaml:"Args,omitempty"`

	Meta `yaml:"-"`
}

func (s *IdentWithArgs) Walk(ctx WalkCtx) {
	if s != nil {
		s.Args.Walk(ctx)
	}
}

type Ident struct {
	Path string `parser:"@Ident" yaml:"Path,omitempty"`

	Meta `yaml:"-"`
}

type Struct struct {
	Statements []Statement `parser:"'{' EOS* (@@ ((EOS|',')+ @@)*)? EOS* '}'" yaml:"Statements,omitempty"`

	Meta `yaml:"-"`
}

func (s *Struct) Walk(ctx WalkCtx) {
	if s != nil {
		Statements(s.Statements).Walk(ctx)
	}
}

type Args struct {
	Statements []Statement `parser:"'(' EOS* (@@ ((EOS|',')+ @@)*)? EOS* ')'" yaml:"Statements,omitempty"`

	Meta `yaml:"-"`
}

func (s *Args) Walk(ctx WalkCtx) {
	if s != nil {
		Statements(s.Statements).Walk(ctx)
	}
}
