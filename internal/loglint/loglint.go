package loglint

import (
	"flag"
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

var rulesCSV string
var keywordsCSV string

func init() {
	Analyzer.Flags = *flag.NewFlagSet("loglint", flag.ContinueOnError)

	Analyzer.Flags.StringVar(
		&rulesCSV,
		"rules",
		"",
		"list of rules (e.g. englishcheck,keywords)",
	)

	Analyzer.Flags.StringVar(
		&keywordsCSV,
		"keywords",
		"",
		"list of keywords (for rule keywords)",
	)
}

func RulesFactories() []rules.Rule {
	var kw []string
	for _, dkw := range DefaultKeyWords {
		kw = append(kw, strings.ToLower(dkw))
	}

	extra := splitCSV(keywordsCSV)
	if len(extra) > 0 {
		for _, ekw := range extra {
			kw = append(kw, strings.ToLower(ekw))
		}
	}

	ruleFactories := map[string]func() rules.Rule{
		"englishcheck":   func() rules.Rule { return englishcheck.NewRule() },
		"specialsymbols": func() rules.Rule { return specialsymbols.NewRule() },
		"lowercase":      func() rules.Rule { return lowercase.NewRule() },
		"keywords":       func() rules.Rule { return keywords.NewRule(kw) },
	}

	enabled := splitCSV(rulesCSV)
	if len(enabled) == 0 {
		return []rules.Rule{
			ruleFactories["englishcheck"](),
			ruleFactories["specialsymbols"](),
			ruleFactories["keywords"](),
			ruleFactories["lowercase"](),
		}
	}

	var checked []rules.Rule
	for _, name := range enabled {
		f, ok := ruleFactories[name]
		if !ok {
			continue
		}
		checked = append(checked, f())
	}
	return checked
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

	enabledRules := RulesFactories()

	for _, file := range pass.Files {
		isImported := false
		for _, imp := range file.Imports {
			if imp.Path != nil {
				if contains(LogPackages, strings.Trim(imp.Path.Value, "\"")) {
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
					for _, rule := range enabledRules {
						rule.Handle(pass, callExpr, str)
					}
				}
			}
			return true
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

func splitCSV(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
