package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type Comment struct {
	Tags Tags   `parser:"(  '/' @(Ident:Ident) (' ' @(Ident:Ident))* '/' )?"  yaml:"Tags,omitempty"`
	Text string `parser:"" yaml:"Text,omitempty"`

	Meta `yaml:"-"`
}

var CommentRegexp = regexp.MustCompile(`(?m)^/(\w+:\w+)(?:\s+(\w+:\w+))*/(?:\s*\n)?([\s\S]*)$`)

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
