package pipe4

import "github.com/alecthomas/participle/v2/lexer"

var (
	pipe4Lexer = lexer.MustSimple([]lexer.Rule{
		{Name: "Comment", Pattern: `(?:#|//)[^\n]*\n?`},
		{Name: "Ident", Pattern: `[a-zA-Z][\w-.]*`},
		{Name: "String", Pattern: `"[^"]*"`},
		{Name: "Float", Pattern: `\d+\.\d+`},
		{Name: "Int", Pattern: `\d+`},
		{Name: "DeclareAssign", Pattern: `:=`},
		{Name: "Punct", Pattern: `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
		{Name: "Whitespace", Pattern: `[ \t\n\r]+`},
	})
)
