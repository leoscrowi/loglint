package loglint

import (
	"go/ast"
	"strings"

	"github.com/leoscrowi/loglint/internal/patterns"
	"github.com/leoscrowi/loglint/internal/rules"
	"github.com/leoscrowi/loglint/internal/rules/englishcheck"
	"github.com/leoscrowi/loglint/internal/rules/keywords"
	"github.com/leoscrowi/loglint/internal/rules/lowercase"
	"github.com/leoscrowi/loglint/internal/rules/specialsymbols"

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
		patterns.NewAnyLnPattern(),
		patterns.NewAnySecondLnPattern(),
		patterns.NewAnySecondPattern(),
		patterns.NewAnyPattern(),
		patterns.NewFirstPattern(),
		patterns.NewSecondPattern(),
		patterns.NewThirdPattern(),
	}

	var checkedRules = []rules.Rule{
		lowercase.NewRule(),
		englishcheck.NewRule(),
		keywords.NewRule(),
		specialsymbols.NewRule(),
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

		if !isImported {
			continue
		}

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
					for _, rule := range checkedRules {
						rule.Handle(pass, callExpr, str)
					}
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
