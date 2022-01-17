package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pipe4/lang/go/generator"
	"github.com/pipe4/lang/pipe4/ast"
	"github.com/pipe4/lang/pipe4/parser"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func run(context *cli.Context) error {
	fileName := context.Args().First()
	if fileName == "" {
		return errors.New("file path required")
	}
	if !strings.HasSuffix(fileName, ".pipe4") {
		fileName += ".pipe4"
	}

	module := strings.TrimSuffix(fileName, ".pipe4")

	parsedAst, err := parser.ParseFile(fileName)
	if err != nil {
		return fmt.Errorf("failed parse file %v: %w", fileName, err)
	}
	nodeList, err := parsedAst.Statements.AstNode()
	if err != nil {
		return fmt.Errorf("failed convert parser ast to pipe4 ast: %w", err)
	}

	root, err := os.MkdirTemp(os.TempDir(), "pipe4_")
	if err != nil {
		return errors.Wrapf(err, "failed create codegen dir")
	}
	path, err := generator.GenerateMain(ast.NewIdent(module, "", "main"), nodeList, root)
	if err != nil {
		return errors.Wrapf(err, "failed generate main package")
	}
	log.Println(path)

	return nil
}
