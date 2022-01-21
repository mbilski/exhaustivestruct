package analyzer

import (
	"flag"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"path"
	"strings"

	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"golang.org/x/tools/go/analysis"
)

// Analyzer that checks if all struct's fields are initialized
var Analyzer = &analysis.Analyzer{
	Name:     "exhaustivestruct",
	Doc:      "Checks if all struct's fields are initialized",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Flags:    newFlagSet(),
}

// StructPatternList is a comma separated list of expressions to match struct packages and names
// The struct packages have the form example.com/package.ExampleStruct
// The matching patterns can use matching syntax from https://pkg.go.dev/path#Match
// If this list is empty, all structs are tested.
var StructPatternList string

func newFlagSet() flag.FlagSet {
	fs := flag.NewFlagSet("", flag.PanicOnError)
	fs.StringVar(&StructPatternList, "struct_patterns", "", "This is a comma separated list of expressions to match struct packages and names")
	return *fs
}

func run(pass *analysis.Pass) (interface{}, error) {
	structPatterns := strings.FieldsFunc(StructPatternList, func(c rune) bool { return c == ',' })
	// validate the pattern syntax
	for _, pattern := range structPatterns {
		_, err := path.Match(pattern, "")
		if err != nil {
			return nil, fmt.Errorf("invalid struct pattern %s: %w", pattern, err)
		}
	}

	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CompositeLit)(nil),
	}
	ins.WithStack(nodeFilter, func(node ast.Node, push bool, stack []ast.Node) bool {
		if !push {
			return true
		}

		compositeLit := node.(*ast.CompositeLit)
		if compositeLit.Type == nil {
			return true
		}

		var typeName string
		{
			i, ok := compositeLit.Type.(*ast.Ident)
			if ok {
				typeName = i.Name
			} else {
				s, ok := compositeLit.Type.(*ast.SelectorExpr)
				if !ok {
					return true
				}
				typeName = s.Sel.Name
			}
		}

		typeInfo := pass.TypesInfo.TypeOf(compositeLit.Type)
		if typeInfo == nil {
			return true
		}

		if len(structPatterns) > 0 {
			shouldLint := false
			for _, pattern := range structPatterns {
				// We check the patterns for vailidy ahead of time, so we don't need to check the error here
				if match, _ := path.Match(pattern, typeInfo.String()); match {
					shouldLint = true
					break
				}
			}
			if !shouldLint {
				return true
			}
		}

		structInfo, ok := typeInfo.Underlying().(*types.Struct)
		if !ok {
			return true
		}

		// Don't report an error if:
		// 1. This composite literal contains no fields and
		// 2. It's in a return statement and
		// 3. The return statement contains a non-nil error
		if len(compositeLit.Elts) == 0 {
			parentReturnStmt, ok := stack[len(stack)-2].(*ast.ReturnStmt)
			if ok {
				nonNilError := false
				// Check if any of the results has an error type and if that error is set to non-nil
				// (if it's set to nil, the type would be "untyped nil")
				for _, result := range parentReturnStmt.Results {
					if pass.TypesInfo.TypeOf(result).String() == "error" {
						nonNilError = true
					}
				}
				if nonNilError {
					return true
				}
			}
		}

		// Don't report an error if:
		// 1. This composite literal contains no fields and
		// 2. It is a type assertion like `var _ Interface = Impl{}`
		if len(compositeLit.Elts) == 0 {
			switch parent := stack[len(stack)-2].(type) {
			case *ast.ValueSpec:
				// for: var _ Interface = Impl{}
				if len(parent.Names) == 1 && parent.Names[0].Name == "_" {
					return true
				}
			case *ast.UnaryExpr:
				// for: var _ Interface = &Impl{}
				if parent.Op == token.AND && len(stack)-3 >= 0 {
					if parent2, ok := stack[len(stack)-3].(*ast.ValueSpec); ok {
						if len(parent2.Names) == 1 && parent2.Names[0].Name == "_" {
							return true
						}
					}
				}
			}
		}

		samePackage := strings.HasPrefix(typeInfo.String(), pass.Pkg.Path()+".")

		missing := []string{}

		for i := 0; i < structInfo.NumFields(); i++ {
			fieldName := structInfo.Field(i).Name()
			exists := false

			if !samePackage && !structInfo.Field(i).Exported() {
				continue
			}

			for eIndex, e := range compositeLit.Elts {
				if k, ok := e.(*ast.KeyValueExpr); ok {
					if i, ok := k.Key.(*ast.Ident); ok {
						if i.Name == fieldName {
							exists = true
							break
						}
					}
				} else {
					if eIndex == i {
						exists = true
						break
					}
				}
			}

			if !exists {
				missing = append(missing, fieldName)
			}
		}

		if len(missing) == 1 {
			pass.Reportf(node.Pos(), "%s is missing in %s", missing[0], typeName)
		} else if len(missing) > 1 {
			pass.Reportf(node.Pos(), "%s are missing in %s", strings.Join(missing, ", "), typeName)
		}

		return true
	})

	return nil, nil
}
