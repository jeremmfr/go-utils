package basicalter_test

import (
	"testing"

	"github.com/jeremmfr/go-utils/basicalter"
)

func TestUniqueStrings(t *testing.T) {
	sliceOfString := []string{"foo", "bar"}

	if len(basicalter.UniqueStrings(sliceOfString)) != len(sliceOfString) {
		t.Errorf("UniqueStrings return bad slice with slice without duplicate entry")
	}
	sliceOfString = append(sliceOfString, "foo")

	if len(basicalter.UniqueStrings(sliceOfString)) == len(sliceOfString) {
		t.Errorf("UniqueStrings return bad slice with slice with duplicate entry")
	}
}

func TestDelEmptyStrings(t *testing.T) {
	sliceOfString := []string{"foo", "", "bar"}

	if v := basicalter.DelEmptyStrings(sliceOfString); len(v) != 2 {
		t.Errorf("DelEmptyStrings return bad slice %v", v)
	}
}

func TestDelStringInSlice(t *testing.T) {
	sliceOfString := []string{"foo", "baz", "bar", "baz"}

	if v := basicalter.DelStringInSlice("baz", sliceOfString); len(v) != 2 {
		t.Errorf("DelStringInSlice return bad slice %v", v)
	}
}

func TestSortStringsByLengthInc(t *testing.T) {
	s := []string{"Bravo", "Gopher", "Alpha", "Go", "Grin", "Delta"}
	basicalter.SortStringsByLengthInc(s)

	if s[0] != "Go" || s[1] != "Grin" {
		t.Errorf("SortStringsByLength fail to sort slice with smaller first")
	}
}

func TestSortStringsByLengthDec(t *testing.T) {
	s := []string{"Bravo", "Gopher", "Alpha", "Go", "Grin", "Delta"}
	basicalter.SortStringsByLengthDec(s)

	if s[len(s)-1] != "Go" || s[len(s)-2] != "Grin" {
		t.Errorf("SortStringsByLength fail to sort slice with smaller last")
	}
}
