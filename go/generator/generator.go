package generator

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"

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
			return errors.Wrapf(err, "failed to generate %v", c.graph.Ident.GetURI())
		}
		if typeNode.Ident.Match(_go.Func) {
			if err := c.generateFuncInvoke(typeNode.Type); err != nil {
				return errors.Wrapf(err, "failed generate func invokation")
			}
		}
		return nil
	}

	return errors.Wrapf(resolver.Unimplemented, "failed to generate: %+v", c.graph)
}

func (c *Context) generateFuncInvoke(typeNode ast.Type) error {
	c.writeLeftPad()
	c.addImport(c.graph.Ident)
	c.write(c.graph.Ident.Name, "(")
	for i, arg := range c.graph.Type.Args {
		// if len(typeNode.Args)
		if err := c.generateGoExpression(arg.Type, typeNode); err != nil {
			return errors.Wrapf(err, "failed write %v invoke argument", i)
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
		return errors.Wrapf(resolver.Unimplemented, "%+v to %+v", value, goType)
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
			return "", errors.Wrapf(err, "failed write to file %v", filePath)
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
		return errors.Wrapf(err, "failed to generate body for %v %v", c.graph.Ident.GetURI(), funcName)
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
		return nil, "", errors.Wrapf(err, "failed create dir %v", dirName)
	}

	name := stringy.New(ident.Name)
	// name := strings.Split(ctx.graph.Ident.Name, ".")

	fileName := path.Join(dirName, name.SnakeCase("?", "").ToLower()+".go")
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(0600))
	if err != nil {
		return nil, "", errors.Wrapf(err, "failed create file %v", fileName)
	}
	return file, fileName, nil
}
