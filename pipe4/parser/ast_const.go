package parser

import (
	"fmt"
	"math/big"

	"github.com/pipe4/lang/pipe4/ast"
)

type Const struct {
	String   *string   `parser:"@String" json:"String,omitempty"`
	Bool     *Bool     `parser:"| @Bool" json:"Bool,omitempty"`
	Rational *Rational `parser:"| @Rational" json:"Rational,omitempty"`

	Meta `json:"-"`
}

func (c Const) AstType() (*ast.Type, error) {
	t := &ast.Type{}
	switch {
	case c.String != nil:
		t.SetString(*c.String)
	case c.Bool != nil:
		t.SetBool(bool(*c.Bool))
	case c.Rational != nil:
		t.SetRational(c.Rational.Rat)
	default:
		return nil, fmt.Errorf("%v: unimplemented const type: %+v", c.Meta.Pos, c)
	}

	return t, nil
}

type Bool bool

func (b *Bool) Capture(values []string) error {
	if len(values) != 1 {
		return fmt.Errorf("to parse bool i need exactly one string but got: '%+v'", values)
	}
	switch values[0] {
	case `true`:
		*b = true
	case `false`:
		*b = false
	default:
		return fmt.Errorf("failed parse bool from: '%+v'", values[0])
	}
	return nil
}

type Rational struct {
	big.Rat
}

func (r *Rational) Equal(other *Rational) bool {
	if r == nil || other == nil {
		return r == other
	}
	return r.Rat.Cmp(&other.Rat) == 0
}

func (r *Rational) Capture(values []string) error {
	if len(values) != 1 {
		return fmt.Errorf("to parse rational number i need exactly one string but got: '%+v'", values)
	}
	rat, ok := new(big.Rat).SetString(values[0])
	if !ok {
		return fmt.Errorf("failed parse rational number from string: '%v'", values[0])
	}
	*r = Rational{*rat}
	return nil
}
