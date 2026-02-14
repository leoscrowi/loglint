package lowercase

import (
	"strconv"
	"testing"

	"github.com/leoscrowi/loglint/internal/rules/testhelper"
)

func TestToLowerFix_Empty(t *testing.T) {
	if got := toLowerFix(""); got != "" {
		t.Fatalf("toLowerFix(\"\") = %q, want %q", got, "")
	}
}

func TestToLowerFixASCII(t *testing.T) {
	in := "Hello world"
	got := toLowerFix(in)
	want := "hello world"
	if got != want {
		t.Fatalf("toLowerFix(%q) = %q, want %q", in, got, want)
	}
}

func TestToLowerFixUnicode(t *testing.T) {
	in := "Привет"
	got := toLowerFix(in)
	want := "привет"
	if got != want {
		t.Fatalf("toLowerFix(%q) = %q, want %q", in, got, want)
	}
}

func TestNoDiagnostics(t *testing.T) {
	fset, _, call := testhelper.ParseFirstCall(t, `log.Println("x")`)

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := NewRule()
	r.Handle(pass, call, "")

	if len(sink.Diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %#v", len(sink.Diags), sink.Diags)
	}
}

func TestFirstRuneNotUpperNoDiagnostics(t *testing.T) {
	str := "hello"

	fset, _, call := testhelper.ParseFirstCall(t, `log.Println("hello")`)

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := NewRule()
	r.Handle(pass, call, str)

	if len(sink.Diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %#v", len(sink.Diags), sink.Diags)
	}
}

func TestFirstRuneUpper(t *testing.T) {
	str := "Hello"

	fset, _, call := testhelper.ParseFirstCall(t, `log.Println("Hello")`)
	lit := testhelper.FindFirstStringLiteralInCall(call)
	if lit == nil {
		t.Fatalf("expected to find a string literal in call args")
	}

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := NewRule()
	r.Handle(pass, call, str)

	d := testhelper.MustOneDiag(t, sink.Diags)

	wantMsg := "first uppercase letter is not allowed: " + str
	if d.Message != wantMsg {
		t.Fatalf("message mismatch:\n got: %q\nwant: %q", d.Message, wantMsg)
	}

	if len(d.SuggestedFixes) != 1 {
		t.Fatalf("expected 1 suggested fix, got %d: %#v", len(d.SuggestedFixes), d.SuggestedFixes)
	}

	fix := d.SuggestedFixes[0]
	if fix.Message != "Make first letter lowercase" {
		t.Fatalf("unexpected fix message: %q", fix.Message)
	}
	if len(fix.TextEdits) != 1 {
		t.Fatalf("expected 1 text edit, got %d: %#v", len(fix.TextEdits), fix.TextEdits)
	}

	edit := fix.TextEdits[0]
	wantNewText := []byte(strconv.Quote("hello"))
	if string(edit.NewText) != string(wantNewText) {
		t.Fatalf("unexpected NewText:\n got: %q\nwant: %q", string(edit.NewText), string(wantNewText))
	}
}

func TestReportsDiagnosticWithoutFix(t *testing.T) {
	fset, _, call := testhelper.ParseFirstCall(t, `hello := "Hello"; log.Println(hello)`)

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := NewRule()
	r.Handle(pass, call, "Hello")

	d := testhelper.MustOneDiag(t, sink.Diags)
	if len(d.SuggestedFixes) != 1 {
		t.Fatalf("expected no suggested fixes, got %d: %#v", len(d.SuggestedFixes), d.SuggestedFixes)
	}
}
