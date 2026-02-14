package keywords

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"

	"github.com/leoscrowi/loglint/internal/rules"
	"golang.org/x/tools/go/analysis"
)

type Rule struct {
	keywords []string
}

func NewRule(kws []string) *Rule {
	return &Rule{keywords: kws}
}

func (r *Rule) Handle(pass *analysis.Pass, call *ast.CallExpr, str string) {
	for _, arg := range call.Args {
		for _, s := range extractNames(arg) {
			lower := strings.ToLower(s)
			for _, kw := range r.keywords {
				if kw == "" {
					continue
				}
				if strings.Contains(lower, kw) {
					msg := "sensitive data not allowed: " + strconv.Quote(s)

					d := analysis.Diagnostic{
						Pos:     arg.Pos(),
						End:     arg.End(),
						Message: msg,
					}

					if lit, litValue := rules.AsStringLiteral(arg); lit != nil {
						if p, e, matched, ok := keywordPos(lit, litValue, r.keywords); ok {
							d.Pos = p
							d.End = e
							d.Message = "sensitive data not allowed: " + strconv.Quote(matched)
						} else {
							d.Pos = lit.Pos()
							d.End = lit.End()
						}

						fix := removeKeywords(litValue, r.keywords)
						d.SuggestedFixes = []analysis.SuggestedFix{
							{
								Message: "Remove sensitive keywords from string literal",
								TextEdits: []analysis.TextEdit{
									{
										Pos:     lit.Pos(),
										End:     lit.End(),
										NewText: []byte(strconv.Quote(fix)),
									},
								},
							},
						}
					}

					pass.Report(d)
				}
			}
		}
	}
}

func extractNames(e ast.Expr) []string {
	switch v := e.(type) {
	case *ast.BasicLit:
		if v.Kind == token.STRING {
			if unq, err := strconv.Unquote(v.Value); err == nil {
				return []string{unq}
			}
			return []string{v.Value}
		}
		return nil

	case *ast.Ident:
		return []string{v.Name}

	case *ast.SelectorExpr:
		return []string{v.Sel.Name}

	case *ast.BinaryExpr:
		if v.Op == token.ADD {
			out := extractNames(v.X)
			out = append(out, extractNames(v.Y)...)
			return out
		}
		return nil

	case *ast.ParenExpr:
		return extractNames(v.X)

	case *ast.UnaryExpr:
		return extractNames(v.X)

	case *ast.StarExpr:
		return extractNames(v.X)

	case *ast.CallExpr:
		var out []string
		for _, a := range v.Args {
			out = append(out, extractNames(a)...)
		}
		return out

	default:
		return nil
	}
}

func keywordPos(lit *ast.BasicLit, unquoted string, keywords []string) (pos, end token.Pos, matched string, ok bool) {
	raw := lit.Value
	lowerRaw := strings.ToLower(raw)
	lowerUnq := strings.ToLower(unquoted)

	for _, kw := range keywords {
		kw = strings.ToLower(kw)
		if kw == "" {
			continue
		}

		if i := strings.Index(lowerRaw, kw); i >= 0 {
			p := lit.Pos() + token.Pos(i)
			return p, p + token.Pos(len(kw)), kw, true
		}

		if i := strings.Index(lowerUnq, kw); i >= 0 {
			p := lit.Pos() + token.Pos(1+i)
			return p, p + token.Pos(len(kw)), kw, true
		}
	}

	return token.NoPos, token.NoPos, "", false
}

func removeKeywords(s string, keywords []string) string {
	out := s
	lower := strings.ToLower(out)

	for _, kw := range keywords {
		if kw == "" {
			continue
		}
		for {
			i := strings.Index(lower, kw)
			if i < 0 {
				break
			}
			out = out[:i] + out[i+len(kw):]
			lower = lower[:i] + lower[i+len(kw):]
		}
	}

	return out
}
