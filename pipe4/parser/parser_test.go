package parser

import (
	"testing"
)

func TestGetBnf(t *testing.T) {
	t.Run("bnf able to generate", func(t *testing.T) {
		if bnf := GetBnf(); len(bnf) == 0 {
			t.Errorf("bnf generation return empty string")
		} else {
			t.Logf("BNF:\n%v\n\n", bnf)
		}
	})
}
