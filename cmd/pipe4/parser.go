package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/pipe4/lang/pipe4/parser"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var ParserCommand = &cli.Command{
	Name: "parser",
	Subcommands: []*cli.Command{
		ParserBnfCommand,
		ParserAstCommand,
		ParserRailroadCommand,
	},
}

var ParserBnfCommand = &cli.Command{
	Name:        "ebnf",
	Description: "print pipe4 lang ebnf",
	Action: func(context *cli.Context) error {
		if _, err := fmt.Fprintf(os.Stdout, "%v\n", parser.GetBnf()); err != nil {
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
		ast, err := parser.ParseFile(fileName)
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

var ParserRailroadCommand = &cli.Command{
	Name:        "railroad",
	Description: "open railroad diagram",
	Action: func(context *cli.Context) error {
		tempDir, err := os.MkdirTemp("", "pipe4-railroad")
		if err != nil {
			return fmt.Errorf("failed create temp dir: %w", err)
		}

		railroadHtml := path.Join(tempDir, "./railroad.html")
		cmd := exec.Command("go", "run", "-v", "github.com/alecthomas/participle/v2/cmd/railroad@latest", "-w", "-o", railroadHtml)
		cmd.Dir = tempDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmdInput, err := cmd.StdinPipe()
		if err != nil {
			return fmt.Errorf("failed open railroad generator stdin writer")
		}
		if err := cmd.Start(); err != nil {
			return fmt.Errorf("failed start railroad generator")
		}
		ebnf := parser.GetBnf()
		if _, err := fmt.Fprintf(cmdInput, "%v", ebnf); err != nil {
			return fmt.Errorf("failed write ebnf to railroad generator stdin: %w", err)
		}
		if err := cmdInput.Close(); err != nil {
			log.Printf("failed close reailroad generator stdinput: %+v", err)
		}
		if err := cmd.Wait(); err != nil {
			return fmt.Errorf("failed exec railroad generator:\n===EBNF===\n%v\n=====\n\n%+v", ebnf, err)
		}
		htmlUrl := url.URL{Scheme: "file", Path: railroadHtml}
		if err := openBrowser(htmlUrl.String()); err != nil {
			return fmt.Errorf("failed open generated railroad in default browser: %w", err)
		}
		return nil
	},
}

func openBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		return fmt.Errorf("failed open url in default browser %v: %w", url, err)
	}
	return nil
}
