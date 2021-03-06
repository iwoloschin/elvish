package ui

import (
	"testing"

	"github.com/elves/elvish/pkg/tt"
)

func TestMarkLines(t *testing.T) {
	stylesheet := RuneStylesheet{
		'-': Inverse,
		'x': Stylings(FgBlue, BgGreen),
	}
	tt.Test(t, tt.Fn("MarkLines", MarkLines), tt.Table{
		tt.Args("foo  bar foobar").Rets(T("foo  bar foobar")),
		tt.Args(
			"foo  bar foobar", stylesheet,
			"---  xxx ------",
		).Rets(
			T("foo", Inverse).
				ConcatText(T("  ")).
				ConcatText(T("bar", FgBlue, BgGreen)).
				ConcatText(T(" ")).
				ConcatText(T("foobar", Inverse)),
		),
		tt.Args(
			"foo  bar foobar", stylesheet,
			"---",
		).Rets(
			T("foo", Inverse).
				ConcatText(T("  bar foobar")),
		),
		tt.Args(
			"plain1",
			"plain2",
			"foo  bar foobar\n", stylesheet,
			"---  xxx ------",
			"plain3",
		).Rets(
			T("plain1").
				ConcatText(T("plain2")).
				ConcatText(T("foo", Inverse)).
				ConcatText(T("  ")).
				ConcatText(T("bar", FgBlue, BgGreen)).
				ConcatText(T(" ")).
				ConcatText(T("foobar", Inverse)).
				ConcatText(T("\n")).
				ConcatText(T("plain3")),
		),
	})
}
