package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pipe4/lang/pipe4/ast"
)

type Comment struct {
	Tags ast.NodeList `parser:"(  '/' @(Ident:Ident) (' ' @(Ident:Ident))* '/' )?"  json:"Tags,omitempty"`
	Text string       `parser:"" json:"Text,omitempty"`

	Meta `json:"-"`
}

func (c Comment) AstNode() *ast.Node {
	if c.Text == "" {
		return nil
	}
	return &ast.Node{
		Comment: ast.Comment{
			Text: c.Text,
			Tags: c.Tags,
		},
	}
}

var CommentRegexp = regexp.MustCompile(`(?m)^/([\w.-]+:[\w.-]+)(?:\s+([\w.-]+:[\w.-]+))*/(?:\s*\n)?([\s\S]*)$`)

func (c *Comment) Capture(values []string) error {
	*c = Comment{
		Text: strings.Join(values, ""),
	}
	matches := CommentRegexp.FindStringSubmatch(c.Text)
	if matches == nil {
		return nil
	}
	c.Text = matches[len(matches)-1]
	pairs := matches[1 : len(matches)-1]

	for _, pair := range pairs {
		kv := strings.Split(pair, ":")
		tag := ast.Node{}
		switch {
		case len(kv) == 2:
			tag.Ident.Name = kv[0]
			tag.Type.SetString(kv[1])
		case len(kv) == 1:
			tag.Ident.Name = kv[0]
			tag.Type.SetString("")
		default:
			return fmt.Errorf("failed to parse tag in format 'key:val' from '%v'", pair)
		}
		c.Tags = append(c.Tags, tag)
	}
	return nil
}
