package rules

import (
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

type Rule interface {
	Handle(pass *analysis.Pass, call *ast.CallExpr, str string)
}

type StringCandidate struct {
	Value string
	Lit   *ast.BasicLit
}

func FindFirstStringLiteral(e ast.Expr) *ast.BasicLit {
	var lit *ast.BasicLit
	ast.Inspect(e, func(n ast.Node) bool {
		bl, ok := n.(*ast.BasicLit)
		if ok && bl.Kind == token.STRING {
			lit = bl
			return false
		}
		return true
	})
	return lit
}

func AsStringLiteral(e ast.Expr) (*ast.BasicLit, string) {
	lit := FindFirstStringLiteral(e)
	if lit == nil {
		return nil, ""
	}
	v, err := strconv.Unquote(lit.Value)
	if err != nil {
		return nil, ""
	}
	return lit, v
}

func AsDirectStringLiteral(e ast.Expr) (*ast.BasicLit, string) {
	bl, ok := e.(*ast.BasicLit)
	if !ok || bl.Kind != token.STRING {
		return nil, ""
	}
	v, err := strconv.Unquote(bl.Value)
	if err != nil {
		return nil, ""
	}
	return bl, v
}

func ExtractStringCandidates(pass *analysis.Pass, e ast.Expr) []StringCandidate {
	if lit, v := AsDirectStringLiteral(e); lit != nil {
		return []StringCandidate{{Value: v, Lit: lit}}
	}

	if v, ok := constStringValue(pass, e); ok {
		if lit := backingStringLiteralFromDecl(e); lit != nil {
			if unq, err := strconv.Unquote(lit.Value); err == nil {
				return []StringCandidate{{Value: unq, Lit: lit}}
			}
			return []StringCandidate{{Value: v, Lit: lit}}
		}
		return []StringCandidate{{Value: v, Lit: nil}}
	}

	if lit := backingStringLiteralFromDecl(e); lit != nil && lit.Kind == token.STRING {
		if unq, err := strconv.Unquote(lit.Value); err == nil {
			return []StringCandidate{{Value: unq, Lit: lit}}
		}
		return []StringCandidate{{Value: lit.Value, Lit: lit}}
	}

	if be, ok := e.(*ast.BinaryExpr); ok && be.Op == token.ADD {
		lx, okx := constStringValue(pass, be.X)
		ly, oky := constStringValue(pass, be.Y)
		if okx && oky {
			return []StringCandidate{{Value: lx + ly, Lit: nil}}
		}
	}

	switch v := e.(type) {
	case *ast.CallExpr:
		var out []StringCandidate
		for _, a := range v.Args {
			out = append(out, ExtractStringCandidates(pass, a)...)
		}
		return out
	case *ast.ParenExpr:
		return ExtractStringCandidates(pass, v.X)
	case *ast.UnaryExpr:
		return ExtractStringCandidates(pass, v.X)
	case *ast.StarExpr:
		return ExtractStringCandidates(pass, v.X)
	default:
		return nil
	}
}

func constStringValue(pass *analysis.Pass, e ast.Expr) (string, bool) {
	if pass == nil || pass.TypesInfo == nil {
		return "", false
	}
	tv, ok := pass.TypesInfo.Types[e]
	if !ok || tv.Value == nil || tv.Type == nil {
		return "", false
	}
	b, ok := tv.Type.Underlying().(*types.Basic)
	if !ok || b.Kind() != types.String {
		return "", false
	}
	return constant.StringVal(tv.Value), true
}

func backingStringLiteralFromDecl(e ast.Expr) *ast.BasicLit {
	id, ok := e.(*ast.Ident)
	if !ok || id == nil || id.Obj == nil {
		return nil
	}

	switch d := id.Obj.Decl.(type) {
	case *ast.AssignStmt:
		for i, lhs := range d.Lhs {
			lid, ok := lhs.(*ast.Ident)
			if !ok || lid.Name != id.Name {
				continue
			}
			if i < len(d.Rhs) {
				if lit, _ := AsDirectStringLiteral(d.Rhs[i]); lit != nil {
					return lit
				}
				if lit := FindFirstStringLiteral(d.Rhs[i]); lit != nil {
					return lit
				}
			}
		}

	case *ast.ValueSpec:
		for i, name := range d.Names {
			if name == nil || name.Name != id.Name {
				continue
			}
			if i < len(d.Values) {
				if lit, _ := AsDirectStringLiteral(d.Values[i]); lit != nil {
					return lit
				}
				if lit := FindFirstStringLiteral(d.Values[i]); lit != nil {
					return lit
				}
			}
		}
	}

	return nil
}
