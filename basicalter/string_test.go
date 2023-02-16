package basicalter_test

import (
	"testing"
	"unicode"

	"github.com/jeremmfr/go-utils/basicalter"
)

const (
	foobar = "foobar"
)

func TestCutPrefixInString(t *testing.T) {
	if basicalter.CutPrefixInString(nil, "pre") {
		t.Errorf("CutPrefixInString returns true with nil for the string pointer")
	}
	str := foobar
	if !basicalter.CutPrefixInString(&str, "") {
		t.Errorf("CutPrefixInString returns false with empty prefix")
	} else if str != foobar {
		t.Errorf("CutPrefixInString has rewrite input string with empty prefix: %s", str)
	}
	if basicalter.CutPrefixInString(&str, "pre") {
		t.Errorf("CutPrefixInString returns true with other prefix")
	} else if str != foobar {
		t.Errorf("CutPrefixInString has rewrite input string with other prefix: %s", str)
	}
	if !basicalter.CutPrefixInString(&str, "foo") {
		t.Errorf("CutPrefixInString returns false with good prefix")
	} else if str != "bar" {
		t.Errorf("CutPrefixInString has rewrite input string with bad string: %s", str)
	}
}

func TestCutSuffixInString(t *testing.T) {
	if basicalter.CutSuffixInString(nil, "pre") {
		t.Errorf("CutSuffixInString returns true with nil for the string pointer")
	}
	str := foobar
	if !basicalter.CutSuffixInString(&str, "") {
		t.Errorf("CutSuffixInString returns false with empty suffix")
	} else if str != foobar {
		t.Errorf("CutSuffixInString has rewrite input string with empty suffix: %s", str)
	}
	if basicalter.CutSuffixInString(&str, "pre") {
		t.Errorf("CutSuffixInString returns true with other suffix")
	} else if str != foobar {
		t.Errorf("CutSuffixInString has rewrite input string with other suffix: %s", str)
	}
	if !basicalter.CutSuffixInString(&str, "bar") {
		t.Errorf("CutSuffixInString returns false with good suffix")
	} else if str != "foo" {
		t.Errorf("CutSuffixInString has rewrite input string with bad string: %s", str)
	}
}

func TestDelRuneInStringWith(t *testing.T) {
	s1 := "bar ! &foo"
	basicalter.DelRuneInStringWith(nil, unicode.IsDigit)
	basicalter.DelRuneInStringWith(&s1, nil)
	basicalter.DelRuneInStringWith(&s1, unicode.IsSpace)
	if s1 != "bar!&foo" {
		t.Errorf("DelRuneInStringWith doesn't change the original string with good value: %q", s1)
	}

	s2 := "bar ! &foo"
	basicalter.DelRuneInStringWith(&s2, unicode.IsSpace, unicode.IsPunct)
	if s2 != "barfoo" {
		t.Errorf("DelRuneInStringWith doesn't change the original string with good value: %q", s2)
	}

	s3 := "barfoo"
	basicalter.DelRuneInStringWith(&s3, unicode.IsSpace)
	if s3 != "barfoo" {
		t.Errorf("DelRuneInStringWith change the original string with other value: %q", s3)
	}

	var s4 string
	basicalter.DelRuneInStringWith(&s4, func(r rune) bool { return !unicode.IsPrint(r) })
	if s4 != "" {
		t.Errorf("DelRuneInStringWith change the original empty string with other value: %q", s4)
	}

	s5 := string('\n')
	basicalter.DelRuneInStringWith(&s5, func(r rune) bool { return !unicode.IsPrint(r) })
	if s5 != "" {
		t.Errorf("DelRuneInStringWith doesn't change the original string with no printable value: %q", s5)
	}
}

func TestFilterRuneInStringWith(t *testing.T) {
	s1 := "foo ! &baz"
	basicalter.FilterRuneInStringWith(nil, unicode.IsLetter)
	basicalter.FilterRuneInStringWith(&s1, nil)
	basicalter.FilterRuneInStringWith(&s1, unicode.IsLetter)
	if s1 != "foobaz" {
		t.Errorf("FilterRuneInStringWith doesn't change the original string with good value: %q", s1)
	}

	s2 := "foobar"
	basicalter.FilterRuneInStringWith(&s2, unicode.IsLetter)
	if s2 != "foobar" {
		t.Errorf("FilterRuneInStringWith change the original string with other value: %q", s2)
	}
}
