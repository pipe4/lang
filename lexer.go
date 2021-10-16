package lang

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	pipe4Lexer = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"LineComment", `(?://)[^\n]*(?:\n)?`, nil},
			{"BlockCommentStart", `/\*`, lexer.Push("BlockCommentStart")}, // 3
			{"Ident", `[a-zA-Z][\w-.]*`, nil},                             // 4
			{"String", `"[^"]*"`, nil},                                    // 5
			{"Rat", `\d+([./]\d+)?`, nil},                                 // 6
			{"Punctuation", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},     // 7
			{"Whitespace", `[ \t\n\r]+`, nil},                             // 8
		},
		"BlockCommentStart": {
			{"BlockCommentEnd", `\*/`, lexer.Pop()},
			{"BlockComment", `[\s\S][^*]*`, nil},
		},
		// 2
	})
)

func Lexer() *lexer.StatefulDefinition {
	return pipe4Lexer
}
