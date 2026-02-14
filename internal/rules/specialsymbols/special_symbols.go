package specialsymbols

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
	"unicode"

	"github.com/leoscrowi/loglint/internal/rules"
	"golang.org/x/tools/go/analysis"
)

type Rule struct{}

func NewRule() *Rule { return &Rule{} }

var allowed = []*unicode.RangeTable{
	unicode.L,
	unicode.N,
	unicode.Z,
}

func (r *Rule) Handle(pass *analysis.Pass, call *ast.CallExpr, _ string) {
	for _, arg := range call.Args {
		cands := rules.ExtractStringCandidates(pass, arg)
		for _, c := range cands {
			if c.Value == "" {
				continue
			}

			bad, posDx := findBadRunes(c.Value)
			if len(bad) == 0 {
				continue
			}

			msg := "special symbols are not allowed: " + string(bad) + " in " + strconv.Quote(c.Value)

			d := analysis.Diagnostic{
				Pos:     arg.Pos(),
				End:     arg.End(),
				Message: msg,
			}

			if c.Lit != nil {
				d.Pos = c.Lit.Pos() + token.Pos(posDx) + 1
				d.End = c.Lit.End()
				fix := removeSpecials(c.Value)
				d.SuggestedFixes = []analysis.SuggestedFix{
					{
						Message: "Remove special symbols and emojis",
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

func findBadRunes(s string) ([]rune, int) {
	var bad []rune
	posDx := 0
	pos := 0
	f := false
	for _, ru := range s {
		if !unicode.IsOneOf(allowed, ru) {
			bad = append(bad, ru)
			if !f {
				f = true
				posDx = pos
			}
		}
		pos++
	}
	return bad, posDx
}

func removeSpecials(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, ru := range s {
		if unicode.IsOneOf(allowed, ru) {
			b.WriteRune(ru)
		}
	}
	return b.String()
}
