package patterns

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/pattern"
)

type Pattern interface {
	GetPattern() pattern.Pattern
	HandleString(pass *analysis.Pass, call *ast.CallExpr) string
}
