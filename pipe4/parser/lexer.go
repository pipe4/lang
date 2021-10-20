package parser

import (
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	pipe4Lexer = lexer.MustStateful(lexer.Rules{
		"Root": {
			{"LineComment", `(?://)[^\n]*(?:\n)?`, nil},
			{"BlockCommentStart", `/\*`, lexer.Push("BlockCommentStart")},
			{"Bool", `true|false`, nil},
			{"Nil", `nil`, nil},
			{"Void", `void`, nil},
			{"Ident", `[a-zA-Z][\w-.]*`, nil},
			{"String", `"[^"]*"`, nil},
			{"Rat", `\d+([./]\d+)?`, nil},
			{"Punctuation", `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`, nil},
			{"Whitespace", `[ \t\n\r]+`, nil},
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
