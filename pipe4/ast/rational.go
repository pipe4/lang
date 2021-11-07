package ast

import (
	"math/big"
)

type Rational struct {
	big.Rat
}

func (r Rational) Equal(other Rational) bool {
	return r.Cmp(&other.Rat) == 0
}

func (r Rational) IsZero() bool {
	return r.Rat.Cmp(&big.Rat{}) == 0
}
