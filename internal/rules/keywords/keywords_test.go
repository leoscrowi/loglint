package keywords

import (
	"strconv"
	"strings"
	"testing"

	"github.com/leoscrowi/loglint/internal/rules/testhelper"
)

func TestOnString(t *testing.T) {
	testVar := `log.Println("password:")`
	fset, _, call := testhelper.ParseFirstCall(t, testVar)

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := &Rule{}
	r.Handle(pass, call, "")

	d := testhelper.MustOneDiag(t, sink.Diags)

	if !strings.Contains(d.Message, "sensitive data not allowed") {
		t.Fatalf("unexpected message: %q", d.Message)
	}
	if !strings.Contains(d.Message, strconv.Quote("password")) && !strings.Contains(strings.ToLower(d.Message), "password") {
		t.Fatalf("message does not seem to mention keyword: %q", d.Message)
	}

	if len(d.SuggestedFixes) != 1 {
		t.Fatalf("expected 1 suggested fix, got %d", len(d.SuggestedFixes))
	}
	if len(d.SuggestedFixes[0].TextEdits) != 1 {
		t.Fatalf("expected 1 text edit, got %d", len(d.SuggestedFixes[0].TextEdits))
	}

	lit := testhelper.FindFirstStringLiteralInCall(call)
	if lit == nil {
		t.Fatalf("expected string literal in call args")
	}

	edit := d.SuggestedFixes[0].TextEdits[0]
	if edit.Pos != lit.Pos() || edit.End != lit.End() {
		t.Fatalf("edit positions do not match literal: edit[%v,%v] lit[%v,%v]",
			edit.Pos, edit.End, lit.Pos(), lit.End())
	}

	unquoted, err := strconv.Unquote(lit.Value)
	if err != nil {
		t.Fatalf("cannot unquote literal %q: %v", lit.Value, err)
	}
	wantFix := strconv.Quote(removeKeywords(unquoted))
	if string(edit.NewText) != wantFix {
		t.Fatalf("unexpected fix: got %s want %s", string(edit.NewText), wantFix)
	}
}

func TestOnStringUppercase(t *testing.T) {
	testVar := `log.Println("api:" + APIKEY)`
	fset, _, call := testhelper.ParseFirstCall(t, testVar)

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := &Rule{}
	r.Handle(pass, call, "")

	d := testhelper.MustOneDiag(t, sink.Diags)

	if !strings.Contains(d.Message, "sensitive data not allowed") {
		t.Fatalf("unexpected message: %q", d.Message)
	}
	if !strings.Contains(d.Message, strconv.Quote("apikey")) && !strings.Contains(strings.ToLower(d.Message), "apikey") {
		t.Fatalf("message does not seem to mention keyword: %q", d.Message)
	}

	if len(d.SuggestedFixes) != 1 {
		t.Fatalf("expected 1 suggested fix, got %d", len(d.SuggestedFixes))
	}
	if len(d.SuggestedFixes[0].TextEdits) != 1 {
		t.Fatalf("expected 1 text edit, got %d", len(d.SuggestedFixes[0].TextEdits))
	}

	lit := testhelper.FindFirstStringLiteralInCall(call)
	if lit == nil {
		t.Fatalf("expected string literal in call args")
	}

	edit := d.SuggestedFixes[0].TextEdits[0]
	if edit.Pos != lit.Pos() || edit.End != lit.End() {
		t.Fatalf("edit positions do not match literal: edit[%v,%v] lit[%v,%v]",
			edit.Pos, edit.End, lit.Pos(), lit.End())
	}

	unquoted, err := strconv.Unquote(lit.Value)
	if err != nil {
		t.Fatalf("cannot unquote literal %q: %v", lit.Value, err)
	}
	wantFix := strconv.Quote(removeKeywords(unquoted))
	if string(edit.NewText) != wantFix {
		t.Fatalf("unexpected fix: got %s want %s", string(edit.NewText), wantFix)
	}
}

func TestOnStringManyReportings(t *testing.T) {
	testVar := `log.Println("apikey:" + APIKEY)`
	fset, _, call := testhelper.ParseFirstCall(t, testVar)

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := &Rule{}
	r.Handle(pass, call, "")

	if len(sink.Diags) != 2 {
		t.Fatalf("expected 2 diagnostics, got %d: %#v", len(sink.Diags), sink.Diags)
	}
}
