package englishcheck

import (
	"go/ast"
	"strconv"
	"strings"
	"unicode"

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
			if !check(c.Value) {
				continue
			}

			msg := "english check rule: " + c.Value

			d := analysis.Diagnostic{
				Pos:     arg.Pos(),
				End:     arg.End(),
				Message: msg,
			}

			if c.Lit != nil {
				fix := removeLetters(c.Value)

				d.SuggestedFixes = []analysis.SuggestedFix{
					{
						Message: "Remove non-english letters from string literal",
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

func removeLetters(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			if isEnglish(r) {
				b.WriteRune(r)
			}
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

func check(str string) bool {
	for _, r := range str {
		if unicode.IsLetter(r) && !isEnglish(r) {
			return true
		}
	}
	return false
}

func isEnglish(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}
