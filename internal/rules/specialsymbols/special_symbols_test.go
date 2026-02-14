package specialsymbols

import (
	"strconv"
	"testing"

	"github.com/leoscrowi/loglint/internal/rules/testhelper"
)

func Test_EmptyString(t *testing.T) {
	fset, _, call := testhelper.ParseFirstCall(t, `log.Println("")`)

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := NewRule()
	r.Handle(pass, call, "")

	if len(sink.Diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %#v", len(sink.Diags), sink.Diags)
	}
}

func TestNoDiagnostics(t *testing.T) {
	str := "Hello 123"

	fset, _, call := testhelper.ParseFirstCall(t, `log.Println("Hello 123")`)

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := NewRule()
	r.Handle(pass, call, str)

	if len(sink.Diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %#v", len(sink.Diags), sink.Diags)
	}
}

func TestReportsDiagnosticWithSuggestedFix(t *testing.T) {
	str := "Hi🙂!,;..."

	fset, _, call := testhelper.ParseFirstCall(t, `log.Println("Hi🙂!,;...")`)
	lit := testhelper.FindFirstStringLiteralInCall(call)
	if lit == nil {
		t.Fatalf("expected to find a string literal in call args")
	}

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := NewRule()
	r.Handle(pass, call, str)

	d := testhelper.MustOneDiag(t, sink.Diags)

	wantMsg := "special symbols are not allowed: 🙂!,;... in " + strconv.Quote(str)
	if d.Message != wantMsg {
		t.Fatalf("message mismatch:\n got: %q\nwant: %q", d.Message, wantMsg)
	}

	if len(d.SuggestedFixes) != 1 {
		t.Fatalf("expected 1 suggested fix, got %d: %#v", len(d.SuggestedFixes), d.SuggestedFixes)
	}

	fix := d.SuggestedFixes[0]
	if fix.Message != "Remove special symbols and emojis" {
		t.Fatalf("unexpected fix message: %q", fix.Message)
	}
	if len(fix.TextEdits) != 1 {
		t.Fatalf("expected 1 text edit, got %d: %#v", len(fix.TextEdits), fix.TextEdits)
	}

	edit := fix.TextEdits[0]
	if edit.Pos != lit.Pos() || edit.End != lit.End() {
		t.Fatalf("expected edit range to match literal; got [%v,%v), want [%v,%v)",
			edit.Pos, edit.End, lit.Pos(), lit.End())
	}

	wantNewText := []byte(strconv.Quote("Hi"))
	if string(edit.NewText) != string(wantNewText) {
		t.Fatalf("unexpected NewText:\n got: %q\nwant: %q", string(edit.NewText), string(wantNewText))
	}
}
