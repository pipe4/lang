package generator

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gobeam/stringy"
	_go "github.com/pipe4/lang/go"
	"github.com/pipe4/lang/go/loader"
	"github.com/pipe4/lang/pipe4/ast"
	"github.com/pipe4/lang/pipe4/resolver"
)

func (c *Context) write(strings ...string) {
	c.body = append(c.body, strings...)
}

func (c *Context) addImport(ident ast.Ident) {
	c.imports = append(c.imports, ident.GetImportURI())
}

func (c *Context) writeLeftPad() {
	c.write(strings.Repeat("\t", c.leftPad))
}

func (c *Context) Generate() error {
	// if ctx.graph.Ident.Match(_go.Func) {
	// 	return generateFunc(ctx)
	// }
	if c.graph.Type.Ident != (ast.Ident{}) {
		typeNode, err := loader.Resolve(c.graph.Type.Ident)
		if err != nil {
			return fmt.Errorf("failed to generate %v: %w", c.graph.Ident.GetURI(), err)
		}
		if typeNode.Ident.Match(_go.Func) {
			if err := c.generateFuncInvoke(typeNode.Type); err != nil {
				return fmt.Errorf("failed generate func invokation: %w", err)
			}
		}
	}

	return fmt.Errorf("failed to generate: %w", resolver.Unimplemented)
}

func (c *Context) generateFuncInvoke(typeNode ast.Type) error {
	c.writeLeftPad()
	c.addImport(c.graph.Ident)
	c.write(c.graph.Ident.Name, "(")
	for i, arg := range c.graph.Type.Args {
		// if len(typeNode.Args)
		if err := c.generateGoExpression(arg.Type, typeNode); err != nil {
			return fmt.Errorf("failed write %v invoke argument: %w", i, err)
		}
	}
	c.write(")", "\n")
	return nil
}

func (c *Context) generateGoExpression(value ast.Type, goType ast.Type) error {
	switch value.BodyOneOf {
	case ast.BodyString:
		c.write(`"`, value.String, `"`)
	default:
		return fmt.Errorf("%w: %+v to %+v", resolver.Unimplemented, value, goType)
	}
	return nil
}

func (c *Context) generateFile(typeNode ast.Node, isMain bool, bodyGenerator func() error) (string, error) {
	if err := c.generateFunc(typeNode, bodyGenerator); err != nil {
		return "", err
	}
	body := c.body
	c.body = []string{}

	file, filePath, err := c.getFile(c.graph.Ident, isMain)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if isMain {
		c.write("package main\n\n")
	} else {
		c.write("package ", path.Base(path.Dir(filePath)), "\n\n")
	}

	if len(c.imports) > 0 {
		c.write("import (\n")
		for _, item := range c.imports {
			c.write("\t", `"`, item, `"`, "\n")
		}
		c.write(")\n\n")
	}

	c.write(body...)

	for _, item := range c.body {
		if _, err := file.WriteString(item); err != nil {
			return "", fmt.Errorf("failed write to file %v: %w", filePath, err)
		}
	}
	return filePath, nil
}

func (c *Context) generateFunc(typeNode ast.Node, bodyGenerator func() error) error {
	name := stringy.New(c.graph.Ident.Name)
	funcName := name.CamelCase("?", "")

	args := ""
	retTypes := ""
	// if
	// fmt.Sprintf("")

	c.write(fmt.Sprintf("func %s(%s) %s {\n", funcName, args, retTypes))
	if err := bodyGenerator(); err != nil {
		return fmt.Errorf("failed to generate body for %v %v: %w", c.graph.Ident.GetURI(), funcName, err)
	}
	c.write("}\n\n")
	return nil
}

func (c *Context) getFile(ident ast.Ident, isMain bool) (*os.File, string, error) {
	cmd := ""
	if isMain {
		cmd = "cmd"
	}
	dirName := path.Join(c.root, cmd, ident.GetImportURI())
	if err := os.MkdirAll(dirName, os.FileMode(0700)); err != nil {
		return nil, "", fmt.Errorf("failed create dir %v: %w", dirName, err)
	}

	name := stringy.New(ident.Name)
	// name := strings.Split(ctx.graph.Ident.Name, ".")

	fileName := path.Join(dirName, name.SnakeCase("?", "").ToLower()+".go")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(0600))
	if err != nil {
		return nil, "", fmt.Errorf("failed create file %v: %w", fileName, err)
	}
	return file, fileName, nil
}
