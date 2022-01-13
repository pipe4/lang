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
		Action:               run,
		ArgsUsage:            "./some_file.[someNode]",
		Description:          "run pipeline",
		BashComplete: func(c *cli.Context) {
			if c.NArg() > 0 {
				fmt.Println(c.Args().Slice())
				completePipe4(c.Args().Get(c.NArg() - 1))
				return
			}
			completePipe4("")
		},
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
