package englishcheck

import (
	"strconv"
	"testing"

	"github.com/leoscrowi/loglint/internal/rules/testhelper"
)

func TestIsEnglish(t *testing.T) {
	cases := []struct {
		r    rune
		want bool
	}{
		{'A', true},
		{'Z', true},
		{'a', true},
		{'z', true},
		{'G', true},
		{'я', false},
		{'Ж', false},
		{'1', false},
		{'_', false},
	}

	for _, tc := range cases {
		if got := isEnglish(tc.r); got != tc.want {
			t.Fatalf("isEnglish(%q) = %v, want %v", tc.r, got, tc.want)
		}
	}
}

func TestCheck(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", false},
		{"Hello world", false},
		{"123 !?", false},
		{"🙂", false},
		{"Hi Мир", true},
		{"Привет", true},
		{"Hello, 世界", true},
	}

	for _, tc := range cases {
		if got, _ := check(tc.in); got != tc.want {
			t.Fatalf("check(%q) = %v, want %v", tc.in, got, tc.want)
		}
	}
}

func TestRemoveLetters(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"Hello world", "Hello world"},
		{"Hi, Мир!", "Hi, !"},
		{"AБC", "AC"},
		{"123-Привет-456", "123--456"},
		{"Hello, 世界!", "Hello, !"},
		{"🙂Hi🙂", "🙂Hi🙂"},
		{"Go_язык", "Go_"},
	}

	for _, tc := range cases {
		if got := removeLetters(tc.in); got != tc.want {
			t.Fatalf("removeLetters(%q) = %q, want %q", tc.in, got, tc.want)
		}
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

func TestOnlyEnglishNoDiagnostics(t *testing.T) {
	str := "Hello world"

	fset, _, call := testhelper.ParseFirstCall(t, `log.Println("Hello world")`)

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := NewRule()
	r.Handle(pass, call, str)

	if len(sink.Diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %#v", len(sink.Diags), sink.Diags)
	}
}

func TestNonEnglishReportsDiagnosticWithFix(t *testing.T) {
	str := "Hi Мир!"

	fset, _, call := testhelper.ParseFirstCall(t, `println("Hi Мир!")`)
	lit := testhelper.FindFirstStringLiteralInCall(call)
	if lit == nil {
		t.Fatalf("expected to find a string literal in call args")
	}

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := NewRule()
	r.Handle(pass, call, str)

	d := testhelper.MustOneDiag(t, sink.Diags)

	wantMsg := "english check rule: " + str
	if d.Message != wantMsg {
		t.Fatalf("message mismatch:\n got: %q\nwant: %q", d.Message, wantMsg)
	}

	if len(d.SuggestedFixes) != 1 {
		t.Fatalf("expected 1 suggested fix, got %d: %#v", len(d.SuggestedFixes), d.SuggestedFixes)
	}
	fix := d.SuggestedFixes[0]

	if len(fix.TextEdits) != 1 {
		t.Fatalf("expected 1 text edit, got %d: %#v", len(fix.TextEdits), fix.TextEdits)
	}
	edit := fix.TextEdits[0]

	if edit.Pos != lit.Pos() || edit.End != lit.End() {
		t.Fatalf("expected edit range to match literal; got [%v,%v), want [%v,%v)",
			edit.Pos, edit.End, lit.Pos(), lit.End())
	}

	wantNewText := []byte(strconv.Quote("Hi !"))
	if string(edit.NewText) != string(wantNewText) {
		t.Fatalf("unexpected NewText:\n got: %q\nwant: %q", string(edit.NewText), string(wantNewText))
	}
}

func TestReportsDiagnosticWithoutFix(t *testing.T) {
	fset, _, call := testhelper.ParseFirstCall(t, `s := "Hi Мир!"; log.Println(s)`)

	var sink testhelper.DiagSink
	pass := sink.Pass(fset)

	r := NewRule()
	r.Handle(pass, call, "Hi Мир!")

	d := testhelper.MustOneDiag(t, sink.Diags)

	wantMsg := "english check rule: " + "Hi Мир!"
	if d.Message != wantMsg {
		t.Fatalf("message mismatch:\n got: %q\nwant: %q", d.Message, wantMsg)
	}

	if len(d.SuggestedFixes) != 1 {
		t.Fatalf("expected no suggested fixes, got %d: %#v", len(d.SuggestedFixes), d.SuggestedFixes)
	}
}
