package patterns

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/pattern"
)

type ThirdPattern struct {
	p pattern.Pattern
}

func (t ThirdPattern) HandleString(pass *analysis.Pass, call *ast.CallExpr) string {
	if call == nil || len(call.Args) < 3 {
		return ""
	}

	str, _ := code.ExprToString(pass, call.Args[2])
	return str
}

func (t ThirdPattern) GetPattern() pattern.Pattern {
	return t.p
}

func NewThirdPattern() *ThirdPattern {
	return &ThirdPattern{
		p: pattern.MustParse(`
		(CallExpr
		(Symbol
			name@(Or
				"log/slog.LogAttrs"
				"log/slog.Log"
				"(*log/slog.Logger).Log"
				"(*log/slog.Logger).LogAttrs"))
		args)
`),
	}
}
