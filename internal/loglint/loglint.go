package loglint

import (
	"go/ast"
	"strings"

	"github.com/leoscrowi/loglint/internal/patterns"
	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/pattern"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "reports log messages",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	var patt = []patterns.Pattern{
		patterns.NewAnySecondPattern(),
		patterns.NewAnyPattern(),
		patterns.NewFirstPattern(),
		patterns.NewSecondPattern(),
		patterns.NewThirdPattern(),
	}

	for _, file := range pass.Files {
		isImported := false
		for _, imp := range file.Imports {
			if imp.Path != nil {
				if contains(logPackages, strings.Trim(imp.Path.Value, "\"")) {
					isImported = true
					break
				}
			}
		}

		// skip if it's no imports
		if !isImported {
			continue
		}

		// we check functions with imports
		matcher := &pattern.Matcher{
			TypesInfo: pass.TypesInfo,
		}

		ast.Inspect(file, func(n ast.Node) bool {
			callExpr, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			for _, p := range patt {
				ok := matcher.Match(p.GetPattern(), callExpr)
				if ok {
					str := p.HandleString(pass, callExpr)
					pass.Reportf(callExpr.Pos(), "%s", str)
				}
			}
			return false
		})
	}
	return nil, nil
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
