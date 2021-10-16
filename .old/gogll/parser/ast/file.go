package ast

import (
	"fmt"
	"github.com/pipe4/lang/ast"
	"github.com/pipe4/lang/gogll/lexer"
	"github.com/pipe4/lang/gogll/parser"
	"go.uber.org/multierr"
	"os"
)

func ParseFile(path string) (*ast.File, error) {
	sourceBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed read file %v: %w", path, err)
	}
	source := []rune(string(sourceBytes))
	pipe4Lexer := lexer.New(source)
	set, errs := parser.Parse(pipe4Lexer)
	if errs != nil {
		return nil, fmt.Errorf("failed to parse file %v: %w", path, ToError(errs))
	}
	children := set.GetRoot()
	fmt.Println(children)
	for _, lvl1 := range children {
		for _, lvl2 := range lvl1 {
			lvl2.
		}
	}
}

type Error parser.Error

func (e Error) Error() string {
	return (*parser.Error)(&e).String()
}

func ToError(errors []*parser.Error) error {
	var errs []error
	for _, err := range errors {
		errs = append(errs, Error(*err))
	}
	return multierr.Combine(errs...)
}

