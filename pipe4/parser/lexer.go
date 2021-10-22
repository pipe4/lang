package parser

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	pipe4Lexer = lexer.MustStateful(lexer.Rules{
		"Root": {
			{Name: "LineComment", Pattern: `(?://)[^\n]*(?:\n)?`},
			{Name: "BlockCommentStart", Pattern: `/\*`, Action: lexer.Push("BlockCommentStart")},
			{Name: "Bool", Pattern: `true|false`},
			{Name: "Nil", Pattern: `nil`},
			{Name: "Void", Pattern: `void`},
			{Name: "Ident", Pattern: `[a-zA-Z][\w-.]*`},
			{Name: "String", Pattern: `"[^"]*"`},
			{Name: "Rat", Pattern: `\d+([./]\d+)?`},
			{Name: "Punctuation", Pattern: `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
			{Name: "Whitespace", Pattern: `[ \t\n\r]+`},
		},
		"BlockCommentStart": {
			{Name: "BlockCommentEnd", Pattern: `\*/`, Action: lexer.Pop()},
			{Name: "BlockComment", Pattern: `[\s\S][^*]*`},
		},
		// 2
	})
)

func Lexer() *lexer.StatefulDefinition {
	return pipe4Lexer
}
