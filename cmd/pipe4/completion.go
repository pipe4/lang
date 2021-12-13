package main

import (
	_ "embed"
	"fmt"

	"github.com/urfave/cli/v2"
)

//go:embed bash_autocomplete
var bashAutocomplete string

var CompletionCommand = &cli.Command{
	Name: "completion",
	Subcommands: []*cli.Command{{
		Name: "bash",
		Action: func(c *cli.Context) error {
			fmt.Println(bashAutocomplete)
			return nil
		},
	}},
}

func completePipe4(prefix string) {
	fmt.Println(prefix)
}
