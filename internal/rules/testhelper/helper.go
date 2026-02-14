package testhelper

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"testing"

	"golang.org/x/tools/go/analysis"
)

type DiagSink struct {
	Diags []analysis.Diagnostic
}

func (s *DiagSink) Pass(fset *token.FileSet) *analysis.Pass {
	return &analysis.Pass{
		Fset: fset,
		Report: func(d analysis.Diagnostic) {
			s.Diags = append(s.Diags, d)
		},
	}
}

func MustOneDiag(t *testing.T, diags []analysis.Diagnostic) analysis.Diagnostic {
	t.Helper()
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %#v", len(diags), diags)
	}
	return diags[0]
}

func FindFirstStringLiteralInCall(call *ast.CallExpr) *ast.BasicLit {
	for _, a := range call.Args {

		var lit *ast.BasicLit

		ast.Inspect(a, func(n ast.Node) bool {
			bl, ok := n.(*ast.BasicLit)
			if ok && bl.Kind == token.STRING {
				lit = bl
				return false
			}
			return true
		})
		//
		if lit != nil {
			_, err := strconv.Unquote(lit.Value)
			if err == nil {
				return lit
			}
		}
	}
	return nil
}

func ParseFirstCall(t *testing.T, body string) (*token.FileSet, *ast.File, *ast.CallExpr) {
	t.Helper()

	src := "package log\nfunc _(){\n" + body + "\n}\n"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "p.go", src, 0)
	if err != nil {
		t.Fatalf("parse error: %v\nsrc:\n%s", err, src)
	}

	var call *ast.CallExpr
	ast.Inspect(f, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok && call == nil {
			call = c
			return false
		}
		return true
	})
	if call == nil {
		t.Fatalf("no CallExpr found in src:\n%s", src)
	}
	return fset, f, call
}
