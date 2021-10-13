package pipe4

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseFile(t *testing.T) {
	file, err := ParseFile("./examples/import.pipe4")
	require.Nilf(t, err, "should parse import.pipe4 file")
	assert.ElementsMatchf(t, []Import{
		{Name: "pipe4", Path: `"github.com/pipe4/lang/pipe4"`},
		{Path: `"github.com/stretchr/testify/require"`},
	}, file.ImportBlock.Imports, "imports should match")
}
