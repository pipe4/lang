package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                 "pipe4",
		Usage:                "pipe4 lang cli tool",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			ParserCommand,
			CompletionCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
