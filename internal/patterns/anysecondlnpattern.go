package patterns

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/pattern"
)

type AnySecondLnPattern struct {
	p pattern.Pattern
}

func (a AnySecondLnPattern) HandleString(pass *analysis.Pass, call *ast.CallExpr) string {
	if call == nil || len(call.Args) < 2 {
		return ""
	}

	args := make([]string, 0, len(call.Args)-1)
	for _, arg := range call.Args[1:] {
		str, _ := code.ExprToString(pass, arg)
		args = append(args, str)
	}

	return strings.Join(args, " ")
}

func (a AnySecondLnPattern) GetPattern() pattern.Pattern {
	return a.p
}

func NewAnySecondLnPattern() *AnySecondLnPattern {
	return &AnySecondLnPattern{
		p: pattern.MustParse(`
	(CallExpr
		(Symbol
			name@(Or
				"(*go.uber.org/zap.SugaredLogger).Logln"))
		args)
`),
	}
}
