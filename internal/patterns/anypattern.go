package patterns

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/pattern"
)

type AnyPattern struct {
	p pattern.Pattern
}

func (a AnyPattern) HandleString(pass *analysis.Pass, call *ast.CallExpr) string {
	if call == nil || len(call.Args) < 1 {
		return ""
	}

	args := make([]string, 0, len(call.Args))
	for _, arg := range call.Args {
		str, _ := code.ExprToString(pass, arg)
		args = append(args, str)
	}

	return strings.Join(args, "")
}

func (a AnyPattern) GetPattern() pattern.Pattern {
	return a.p
}

func NewAnyPattern() *AnyPattern {
	return &AnyPattern{
		p: pattern.MustParse(`
	(CallExpr
		(Symbol
			name@(Or
				"log.Fatal"
				"log.Fatalln"
				"log.Panic"
				"log.Panicln"
				"log.Print"
				"log.Println"
				"(*log.Logger).Fatal"
				"(*log.Logger).Fatalln"
				"(*log.Logger).Panic"
				"(*log.Logger).Panicln"
				"(*log.Logger).Print"
				"(*log.Logger).Println"
				"(*go.uber.org/zap.SugaredLogger).DPanic"
				"(*go.uber.org/zap.SugaredLogger).DPanicln"
				"(*go.uber.org/zap.SugaredLogger).Debug"
				"(*go.uber.org/zap.SugaredLogger).Debugln"
				"(*go.uber.org/zap.SugaredLogger).Error"
				"(*go.uber.org/zap.SugaredLogger).Errorln"
				"(*go.uber.org/zap.SugaredLogger).Fatal"
				"(*go.uber.org/zap.SugaredLogger).Fatalln"
				"(*go.uber.org/zap.SugaredLogger).Info"
				"(*go.uber.org/zap.SugaredLogger).Infoln"
				"(*go.uber.org/zap.SugaredLogger).Panic"
				"(*go.uber.org/zap.SugaredLogger).Panicln"
				"(*go.uber.org/zap.SugaredLogger).Warn"
				"(*go.uber.org/zap.SugaredLogger).Warnln"))
		args)
`)}
}
