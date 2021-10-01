package basicalter_test

import (
	"testing"

	"github.com/jeremmfr/go-utils/basicalter"
	"github.com/jeremmfr/go-utils/basiccheck"
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

func TestReverseStrings(t *testing.T) {
	sliceOfString := []string{"foo", "baz", "bar", "World", "Hello"}

	newSlice := basicalter.ReverseStrings(sliceOfString)

	desiredSlice := []string{"Hello", "World", "bar", "baz", "foo"}
	if !basiccheck.EqualStringSlice(newSlice, desiredSlice) {
		t.Errorf("ReverseStrings fail to reverse slice: %v expected %v", newSlice, desiredSlice)
	}
}

func TestSortStringsByLengthInc(t *testing.T) {
	s := []string{"Bravo", "Gopher", "Alpha", "Go", "Grin", "Delta"}

	basicalter.SortStringsByLengthInc(s)

	desiredSlice := []string{"Go", "Grin", "Alpha", "Bravo", "Delta", "Gopher"}
	if !basiccheck.EqualStringSlice(s, desiredSlice) {
		t.Errorf("SortStringsByLength fail to sort slice with smaller first and lexicographic order")
	}
}

func TestSortStringsByLengthDec(t *testing.T) {
	s := []string{"Bravo", "Gopher", "Alpha", "Go", "Grin", "Delta"}

	basicalter.SortStringsByLengthDec(s)

	desiredSlice := []string{"Gopher", "Alpha", "Bravo", "Delta", "Grin", "Go"}
	if !basiccheck.EqualStringSlice(s, desiredSlice) {
		t.Errorf("SortStringsByLength fail to sort slice with smaller last and lexicographic order")
	}
}
