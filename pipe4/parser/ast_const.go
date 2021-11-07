package parser

import (
	"fmt"
	"math/big"
	"strings"

	"gopkg.in/yaml.v3"
)

type Const struct {
	String   *string   `parser:"@String" yaml:"String,omitempty"`
	Bool     *Bool     `parser:"| @Bool" yaml:"Bool,omitempty"`
	Rational *Rational `parser:"| @Rational" yaml:"Rational,omitempty"`

	Meta `yaml:"-"`
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

func (f *File) ToYaml() (string, error) {
	// pretty.Fprintf(os.Stdout, "%# v", ast)
	out := &strings.Builder{}
	encoder := yaml.NewEncoder(out)
	encoder.SetIndent(2)
	if err := encoder.Encode(f); err != nil {
		return "", fmt.Errorf("failed print ast tree: %w", err)
	}
	str := out.String()
	// str = strings.ReplaceAll(str, "  - ", "- ")
	return str, nil
}

func (f *File) FromYaml(source string) error {
	if err := yaml.Unmarshal([]byte(source), f); err != nil {
		return fmt.Errorf("failed to read ast from yaml: %w", err)
	}
	return nil
}
