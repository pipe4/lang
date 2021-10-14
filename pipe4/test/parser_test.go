package test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pipe4/lang/pipe4"
	"github.com/stretchr/testify/assert"
)

func testAst(t *testing.T, path string, want *pipe4.File) {
	t.Run(path, func(t *testing.T) {
		if got, err := pipe4.ParseFile(path); err != nil {
			assert.NoError(t, err, "error while parsing file")
		} else {
			if !cmp.Equal(want, got) {
				t.Errorf("AST not match\n%v", cmp.Diff(want, got))
			}

			// assert.Equalf(t, string(wantYaml), string(gotYaml), "wrong file ast")
			// assert.Equalf(t, want, got, "wrong file ast")
		}
	})
}
