package patterns

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/analysis/code"
	"honnef.co/go/tools/pattern"
)

type AnyLnPattern struct {
	p pattern.Pattern
}

func (a AnyLnPattern) HandleString(pass *analysis.Pass, call *ast.CallExpr) string {
	if call == nil || len(call.Args) < 1 {
		return ""
	}

	args := make([]string, 0, len(call.Args))
	for _, arg := range call.Args {
		str, _ := code.ExprToString(pass, arg)
		args = append(args, str)
	}

	return strings.Join(args, " ")
}

func (a AnyLnPattern) GetPattern() pattern.Pattern {
	return a.p
}

func NewAnyLnPattern() *AnyLnPattern {
	return &AnyLnPattern{
		p: pattern.MustParse(`
	(CallExpr
		(Symbol
			name@(Or
				"log.Fatalln"
				"log.Panicln"
				"log.Println"
				"(*log.Logger).Fatalln"
				"(*log.Logger).Panicln"
				"(*go.uber.org/zap.SugaredLogger).DPanicln"
				"(*log.Logger).Println"
				"(*go.uber.org/zap.SugaredLogger).Debugln"
				"(*go.uber.org/zap.SugaredLogger).Errorln"
				"(*go.uber.org/zap.SugaredLogger).Fatalln"
				"(*go.uber.org/zap.SugaredLogger).Infoln"
				"(*go.uber.org/zap.SugaredLogger).Panicln"
				"(*go.uber.org/zap.SugaredLogger).Warnln"))
		args)
`)}
}
