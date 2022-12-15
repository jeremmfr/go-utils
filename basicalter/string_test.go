package basicalter_test

import (
	"testing"

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
