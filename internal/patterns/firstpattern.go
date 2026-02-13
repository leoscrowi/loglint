package patterns

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/pattern"
)

type FirstPattern struct {
	p pattern.Pattern
}

func (f FirstPattern) HandleString(pass *analysis.Pass, call *ast.CallExpr) string {
	if call == nil || len(call.Args) < 1 {
		return ""
	}

	str, _ := code.ExprToString(pass, call.Args[0])
	return str
}

func (f FirstPattern) GetPattern() pattern.Pattern {
	return f.p
}

func NewFirstPattern() *FirstPattern {
	return &FirstPattern{
		p: pattern.MustParse(`
	(CallExpr
		(Symbol
			name@(Or
				"log.Fatalf"
				"log.Panicf"
				"log.Printf"
				"(*log.Logger).Printf"
				"(*log.Logger).Panicf"
				"(*log.Logger).Fatalf"
				"log/slog.Debug"
				"log/slog.Error"
				"log/slog.Info"
				"log/slog.Warn"
				"(*log/slog.Logger).Debug"
				"(*log/slog.Logger).Error"
				"(*log/slog.Logger).Info"
				"(*log/slog.Logger).Warn"
				"(*go.uber.org/zap.Logger).DPanic"
				"(*go.uber.org/zap.Logger).Debug"
				"(*go.uber.org/zap.Logger).Error"
				"(*go.uber.org/zap.Logger).Fatal"
				"(*go.uber.org/zap.Logger).Info"
				"(*go.uber.org/zap.Logger).Panic"
				"(*go.uber.org/zap.Logger).Warn"
				"(*go.uber.org/zap.SugaredLogger).DPanicf"
				"(*go.uber.org/zap.SugaredLogger).DPanicw"
				"(*go.uber.org/zap.SugaredLogger).Debugf"
				"(*go.uber.org/zap.SugaredLogger).Debugw"
				"(*go.uber.org/zap.SugaredLogger).Errorf"
				"(*go.uber.org/zap.SugaredLogger).Errorw"
				"(*go.uber.org/zap.SugaredLogger).Fatalf"
				"(*go.uber.org/zap.SugaredLogger).Fatalw"
				"(*go.uber.org/zap.SugaredLogger).Infof"
				"(*go.uber.org/zap.SugaredLogger).Infow"
				"(*go.uber.org/zap.SugaredLogger).Panicf"
				"(*go.uber.org/zap.SugaredLogger).Panicw"
				"(*go.uber.org/zap.SugaredLogger).Warnf"
				"(*go.uber.org/zap.SugaredLogger).Warnw"))
		args)
`),
	}
}
