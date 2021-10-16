package lang

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"

	"github.com/alecthomas/participle/v2/lexer"
	"gopkg.in/yaml.v3"
)

type File struct {
	Import     []Import        `parser:"('import' (@@ | '(' @@* ')'))*" yaml:"Import,omitempty"`
	Statements []RootStatement `parser:"@@*" yaml:"Statements,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

type Import struct {
	Name string `parser:"@Ident?" yaml:"Name,omitempty"`
	Url  string `parser:"@String" yaml:"Url,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

type RootStatement struct {
	Comment *Comment `parser:"( @(LineComment | BlockComment)?"  yaml:"Comment,omitempty"`

	Name string `parser:"( @Ident" yaml:"Name,omitempty"`

	String *string `parser:"( (@String" yaml:"String,omitempty"`
	Number *Rat    `parser:"| @Rat)" yaml:"Number,omitempty"`

	Type   string            `parser:"| ( @Ident?" yaml:"Type,omitempty"`
	Props  []PropsStatement  `parser:"( '(' (@@ (',' @@)*)? ')' )?" yaml:"Props,omitempty"`
	Struct []StructStatement `parser:"( '{' @@* '}' )?  )  ) )?  )!" yaml:"Struct,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

type StructStatement struct {
	Comment *Comment `parser:"( @(LineComment | BlockComment)?" yaml:"Comment,omitempty"`
	Name    string   `parser:"( @Ident" yaml:"Name,omitempty"`

	String *string `parser:"( ( (@String" yaml:"String,omitempty"`
	Number *Rat    `parser:"| @Rat)" yaml:"Number,omitempty"`

	Type   string            `parser:"| ( @Ident" yaml:"Type,omitempty"`
	Props  []PropsStatement  `parser:"( '(' (@@ (',' @@)*)? ')' )?" yaml:"Props,omitempty"`
	Struct []StructStatement `parser:"( '{' @@* '}' )? ) )?" yaml:"Struct,omitempty"`

	Default *DefaultStatement `parser:"( '=' @@ )?  )! )?  )!" yaml:"Default,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

type PropsStatement struct {
	Type   string            `parser:"( @Ident" yaml:"Type,omitempty"`
	String *string           `parser:"| @String" yaml:"String,omitempty"`
	Number *Rat              `parser:"| @Rat" yaml:"Number,omitempty"`
	Struct []StructStatement `parser:"| '{' @@* '}' )" yaml:"Struct,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}
type DefaultStatement struct {
	String *string `parser:"( (@String" yaml:"String,omitempty"`
	Number *Rat    `parser:"| @Rat)" yaml:"Number,omitempty"`

	Type   string            `parser:"| ( @Ident?" yaml:"Type,omitempty"`
	Props  []PropsStatement  `parser:"( '(' (@@ (',' @@)*)? ')' )?" yaml:"Props,omitempty"`
	Struct []StructStatement `parser:"( '{' @@* '}' )? ) )" yaml:"Struct,omitempty"`

	Pos    lexer.Position `parser:"" yaml:"-"`
	EndPos lexer.Position `parser:"" yaml:"-"`
	Tokens []lexer.Token  `parser:"" yaml:"-"`
}

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
		if len(kv) == 2 {
			tags[kv[0]] = kv[1]
		} else if len(kv) == 0 {
			tags[kv[0]] = ""
		} else {
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

type Rat struct {
	big.Rat
}

func (r *Rat) Equal(other *Rat) bool {
	if r == nil || other == nil {
		return r == other
	}
	return r.Rat.Cmp(&other.Rat) == 0
}

func (r *Rat) Capture(values []string) error {
	if len(values) != 1 {
		return fmt.Errorf("to parse rational number i need exactly one string but got: '%+v'", values)
	}
	rat, ok := new(big.Rat).SetString(values[0])
	if !ok {
		return fmt.Errorf("failed parse rational number from string: '%v'", values[0])
	}
	*r = Rat{*rat}
	return nil
}

func (f *File) ToYaml() (string, error) {
	// pretty.Fprintf(os.Stdout, "%# v", ast)
	out := &strings.Builder{}
	encoder := yaml.NewEncoder(out)
	encoder.SetIndent(2)
	if err := encoder.Encode(f); err != nil {
		return "", fmt.Errorf("failed print ast tree: %w", err)
	}
	return out.String(), nil
}

func (f *File) FromYaml(source string) error {
	if err := yaml.Unmarshal([]byte(source), f); err != nil {
		return fmt.Errorf("failed to read ast from yaml: %w", err)
	}
	return nil
}
