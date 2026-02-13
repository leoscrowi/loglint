package patterns

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/pattern"
)

type SecondPattern struct {
	p pattern.Pattern
}

func (s SecondPattern) HandleString(pass *analysis.Pass, call *ast.CallExpr) string {
	if call == nil || len(call.Args) < 2 {
		return ""
	}

	str, _ := code.ExprToString(pass, call.Args[1])
	return str
}

func (s SecondPattern) GetPattern() pattern.Pattern {
	return s.p
}

func NewSecondPattern() *SecondPattern {
	return &SecondPattern{
		p: pattern.MustParse(`
	(CallExpr
		(Symbol
			name@(Or
				"log/slog.DebugContext"
				"log/slog.ErrorContext"
				"log/slog.InfoContext"
				"log/slog.WarnContext"
				"(*log/slog.Logger).DebugContext"
				"(*log/slog.Logger).ErrorContext"
				"(*log/slog.Logger).InfoContext"
				"(*log/slog.Logger).WarnContext"
				"(*go.uber.org/zap.Logger).Log"
				"(*go.uber.org/zap.SugaredLogger).Logf"
				"(*go.uber.org/zap.SugaredLogger).Logw"))
		args)
`),
	}
}
