package parser

// func fixGoExact(f *ast.File) bool {
// 	// This one is harder because the import name changes.
// 	// First find the import spec.
// 	var funcDecl *ast.FuncDecl
// 	walk(f, func(n interface{}) {
// 		if importSpec != nil {
// 			return
// 		}
// 		spec, ok := n.(*ast.ImportSpec)
// 		if !ok {
// 			return
// 		}
// 		path, err := strconv.Unquote(spec.Path.Value)
// 		if err != nil {
// 			return
// 		}
// 		if path == "golang.org/x/tools/go/exact" {
// 			importSpec = spec
// 		}
//
// 	})
// 	if importSpec == nil {
// 		return false
// 	}
//
// 	// We are about to rename exact.* to constant.*, but constant is a common
// 	// name. See if it will conflict. This is a hack but it is effective.
// 	exists := renameTop(f, "constant", "constant")
// 	suffix := ""
// 	if exists {
// 		suffix = "_"
// 	}
// 	// Now we need to rename all the uses of the import. RewriteImport
// 	// affects renameTop, but not vice versa, so do them in this order.
// 	renameTop(f, "exact", "constant"+suffix)
// 	rewriteImport(f, "golang.org/x/tools/go/exact", "go/constant")
// 	// renameTop will also rewrite the imported package name. Fix that;
// 	// we know it should be missing.
// 	importSpec.Name = nil
// 	return true
// }
