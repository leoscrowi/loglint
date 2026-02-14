package lowercase

import (
	"go/ast"
	"strconv"
	"unicode"
	"unicode/utf8"

	"github.com/leoscrowi/loglint/internal/rules"
	"golang.org/x/tools/go/analysis"
)

type Rule struct{}

func NewRule() *Rule { return &Rule{} }

func (r *Rule) Handle(pass *analysis.Pass, call *ast.CallExpr, _ string) {
	for _, arg := range call.Args {
		cands := rules.ExtractStringCandidates(pass, arg)
		for _, c := range cands {
			if c.Value == "" {
				continue
			}

			ru, _ := utf8.DecodeRuneInString(c.Value)
			if ru == utf8.RuneError {
				continue
			}
			if !unicode.IsUpper(ru) {
				continue
			}

			msg := "first uppercase letter is not allowed: " + c.Value

			d := analysis.Diagnostic{
				Pos:     arg.Pos(),
				End:     arg.End(),
				Message: msg,
			}

			if c.Lit != nil {
				fix := toLowerFix(c.Value)
				d.SuggestedFixes = []analysis.SuggestedFix{
					{
						Message: "Make first letter lowercase",
						TextEdits: []analysis.TextEdit{
							{
								Pos:     c.Lit.Pos(),
								End:     c.Lit.End(),
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

func toLowerFix(s string) string {
	rs := []rune(s)
	if len(rs) == 0 {
		return s
	}
	rs[0] = unicode.ToLower(rs[0])
	return string(rs)
}
