package internal

import (
	"strings"

	"github.com/pipe4/lang/pipe4/parser"
)

type commentExpr struct {
	tags []commentTag
	text commentText

	block  bool
	exists bool
}

type commentText struct {
	text string
}

type commentTag struct {
	key commentTagKey
	val commentTagVal
}

type commentTagKey struct {
	text string
}

type commentTagVal struct {
	text string
}

func commentExprFromParser(s *parser.Statement) commentExpr {
	if s == nil || s.Comment == nil {
		return commentExpr{}
	}

	c := commentExpr{
		text:   commentText{text: s.Comment.Text},
		block:  strings.Contains(s.Comment.Text, "\n"),
		exists: true,
	}

	for k, v := range s.Comment.Tags {
		if k == "" {
			continue
		}
		c.tags = append(c.tags, commentTag{key: commentTagKey{text: k}, val: commentTagVal{text: v}})
	}

	return c
}
