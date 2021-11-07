package ast

import (
	"math/big"
)

type Rational struct {
	big.Rat
}

func (r Rational) Equal(other Rational) bool {
	// if r == nil || other == nil {
	// 	return r == other
	// }
	return r.Cmp(&other.Rat) == 0
}

func (r Rational) IsZero() bool {
	return r.Rat.Cmp(&big.Rat{}) == 0
}

// func (r Rational) MarshalJSON() (interface{}, error) {
// 	return r.Rat.RatString(), nil
// }

// func (r Rational) MarshalJSON() ([]byte, error) {
// 	if r.IsZero() {
// 		return []byte(`"0"`), nil
// 	}
// 	return []byte(r.Rat.RatString()), nil
// }
