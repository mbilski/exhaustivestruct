package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name:     "exhaustivestruct",
	Doc:      "Checks if all struct's fields are initialized",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		compositeLit := node.(*ast.CompositeLit)

		ident, ok := compositeLit.Type.(*ast.Ident)

		if !ok {
			return
		}

		tSpec, ok := ident.Obj.Decl.(*ast.TypeSpec)

		if !ok {
			return
		}

		sType, ok := tSpec.Type.(*ast.StructType)

		if !ok {
			return
		}

		if sType.Fields.NumFields() != len(compositeLit.Elts) {
			pass.Reportf(node.Pos(), "missing fields")
		}
	})

	return nil, nil
}
