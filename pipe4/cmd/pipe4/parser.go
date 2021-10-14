package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/pipe4/lang/pipe4"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var ParserCommand = &cli.Command{
	Name: "parser",
	Subcommands: []*cli.Command{
		ParserBnfCommand,
		ParserAstCommand,
	},
}

var ParserBnfCommand = &cli.Command{
	Name:        "bnf",
	Description: "print pipe4 lang bnf",
	Action: func(context *cli.Context) error {
		if _, err := fmt.Fprintf(os.Stdout, "%v\n", pipe4.GetBnf()); err != nil {
			return fmt.Errorf("failed write bnf syntax: %w", err)
		}
		return nil
	},
}

var ParserAstCommand = &cli.Command{
	Name:        "ast",
	ArgsUsage:   "./some_file.pipe4",
	Description: "print ast of file",
	Action: func(context *cli.Context) error {
		fileName := context.Args().First()
		if fileName == "" {
			return errors.New("file path required")
		}
		ast, err := pipe4.ParseFile(fileName)
		if err != nil {
			return fmt.Errorf("failed parse file %v: %w", fileName, err)
		}

		// pretty.Fprintf(os.Stdout, "%# v", ast)
		encoder := yaml.NewEncoder(os.Stdout)
		encoder.SetIndent(2)
		if err := encoder.Encode(ast); err != nil {
			return fmt.Errorf("failed print ast tree: %w", err)
		}
		return nil
	},
}
