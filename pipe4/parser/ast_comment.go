package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/alecthomas/participle/v2/lexer"
)

type Comment struct {
	Tags Tags   `parser:"(  '/' @(Ident:Ident) (' ' @(Ident:Ident))* '/' )?"  yaml:"Tags,omitempty"`
	Text string `parser:"" yaml:"Text,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

var CommentRegexp = regexp.MustCompile(`(?m)^/(\w+:\w+)(?:\s+(\w+:\w+))*/(?:\s*\n)?([\s\S]*)$`)

func (c *Comment) Capture(values []string) error {
	if len(values) != 1 {
		return fmt.Errorf("multiple comment merge into one not implemented: %+v", values)
	}
	*c = Comment{}

	matches := CommentRegexp.FindStringSubmatch(values[0])
	if matches == nil {
		c.Text = values[0]
		return nil
	}
	c.Text = matches[len(matches)-1]
	pairs := matches[1 : len(matches)-1]

	tags := make(map[string]string)
	for _, pair := range pairs {
		kv := strings.Split(pair, ":")
		switch {
		case len(kv) == 2:
			tags[kv[0]] = kv[1]
		case len(kv) == 0:
			tags[kv[0]] = ""
		default:
			return fmt.Errorf("failed to parse tag in format 'key:val' from '%v'", pair)
		}
	}
	if len(tags) > 0 {
		c.Tags = tags
	}
	return nil
}

type Tags map[string]string

func (c *Comment) GetTag(name string) string {
	if c == nil || c.Tags == nil {
		return ""
	}
	return c.Tags[name]
}
