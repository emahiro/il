package main

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/go/ast/inspector"
)

var a = &analysis.Analyzer{
	Name:     "samplechecher",
	Doc:      "sample code checker",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func main() {
	singlechecker.Main(a)
}

func run(pass *analysis.Pass) (interface{}, error) {
	i := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	filter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	i.Preorder(filter, func(n ast.Node) {
		// 必要であればテストファイルは検査から外す。
		if strings.HasSuffix(pass.Fset.File(n.Pos()).Name(), "_test.go") {
			return
		}

		switch n := n.(type) {
		case *ast.FuncDecl:
			name := n.Name.Name
			if name == "Sample" {
				pass.Report(
					analysis.Diagnostic{
						Pos:     n.Pos(),
						End:     n.End(),
						Message: "change Sample to Example",
						SuggestedFixes: []analysis.SuggestedFix{
							{
								Message: "Sample -> Example",
								TextEdits: []analysis.TextEdit{
									{
										Pos:     n.Pos(),
										End:     n.End(),
										NewText: []byte("Example"),
									},
								},
							},
						},
					},
				)
			}
			if name == "SampleWithContext" {
				pass.Report(analysis.Diagnostic{
					Pos:     n.Pos(),
					End:     n.End(),
					Message: "change SampleWithContext to ExampleContext",
					SuggestedFixes: []analysis.SuggestedFix{
						{
							Message: "Sample -> Example",
							TextEdits: []analysis.TextEdit{
								{
									Pos:     n.Pos(),
									End:     n.End(),
									NewText: []byte("ExampleWithContext"),
								},
							},
						},
					},
				})
			}
		}
	})
	return nil, nil
}
