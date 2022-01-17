package generator

import (
	_go "github.com/pipe4/lang/go"
	"github.com/pipe4/lang/pipe4/ast"
)

func GenerateMain(ident ast.Ident, items ast.NodeList, root string) (string, error) {
	ctx := Context{
		root: root,
	}

	filePath, err := ctx.generateFile(
		ast.Node{
			Ident: ident,
			Type: ast.Type{
				Ident: _go.Func,
			},
		},
		true,
		func() error {
			for i, item := range items {
				if i == 0 {
					continue
				}
				subCtx := Context{
					root:    ctx.root,
					graph:   item,
					leftPad: ctx.leftPad + 1,
				}
				if err := subCtx.Generate(); err != nil {
					return err
				}
			}
			return nil
		},
	)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
