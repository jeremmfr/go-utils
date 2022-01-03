package basicalter_test

import (
	"strings"
	"testing"

	"github.com/jeremmfr/go-utils/basicalter"
	"github.com/jeremmfr/go-utils/basiccheck"
)

func TestUniqueStrings(t *testing.T) {
	sliceOfString := []string{"foo", "bar"}

	if len(basicalter.UniqueStrings(sliceOfString)) != len(sliceOfString) {
		t.Errorf("UniqueStrings returned bad slice with slice without duplicate entry")
	}

	sliceOfString = append(sliceOfString, "foo")

	if v := basicalter.UniqueStrings(sliceOfString); len(v) == len(sliceOfString) {
		t.Errorf("UniqueStrings didn't remove duplicate foo in slice: %v", v)
	}
}

func TestDelEmptyStrings(t *testing.T) {
	sliceOfString := []string{"foo", "", "bar"}

	if v := basicalter.DelEmptyStrings(sliceOfString); len(v) != 2 {
		t.Errorf("DelEmptyStrings didn't remove empty string: %v", v)
	} else if !basiccheck.EqualStringSlice(v, []string{"foo", "bar"}) {
		t.Errorf("DelEmptyStrings didn't remove empty string: %v", v)
	}
}

func TestDelStringInSlice(t *testing.T) {
	sliceOfString := []string{"foo", "baz", "bar", "baz"}

	if v := basicalter.DelStringInSlice("baz", sliceOfString); len(v) != 2 {
		t.Errorf("DelStringInSlice didn't remove 'baz': %v", v)
	} else if !basiccheck.EqualStringSlice(v, []string{"foo", "bar"}) {
		t.Errorf("DelStringInSlice didn't remove 'baz': %v", v)
	}
}

func TestFilterStringsWith(t *testing.T) {
	sliceOfString := []string{"foo", "baz", "bar", "baz"}

	if v := basicalter.FilterStringsWith(sliceOfString, func(s string) bool {
		return strings.HasPrefix(s, "ba")
	}); len(v) != 3 {
		t.Errorf("FilterStringsWith didn't remove foo (without prefix 'ba'): %v", v)
	} else if !basiccheck.EqualStringSlice(v, []string{"baz", "bar", "baz"}) {
		t.Errorf("FilterStringsWith didn't remove foo (without prefix 'ba'): %v", v)
	}
}

func TestReverseSlice(t *testing.T) {
	sliceOfString := []string{"foo", "baz", "bar", "World", "Hello"}

	basicalter.ReverseStrings(sliceOfString)

	desiredSlice := []string{"Hello", "World", "bar", "baz", "foo"}
	if !basiccheck.EqualStringSlice(sliceOfString, desiredSlice) {
		t.Errorf("ReverseStrings didn't reverse slice: %v expected %v", sliceOfString, desiredSlice)
	}
}

func TestSortStringsByLengthInc(t *testing.T) {
	s := []string{"Bravo", "Gopher", "Alpha", "Go", "Grin", "Delta"}

	basicalter.SortStringsByLengthInc(s)

	desiredSlice := []string{"Go", "Grin", "Alpha", "Bravo", "Delta", "Gopher"}
	if !basiccheck.EqualStringSlice(s, desiredSlice) {
		t.Errorf("SortStringsByLength didn't sort slice with smaller first and lexicographic order")
	}
}

func TestSortStringsByLengthDec(t *testing.T) {
	s := []string{"Bravo", "Gopher", "Alpha", "Go", "Grin", "Delta"}

	basicalter.SortStringsByLengthDec(s)

	desiredSlice := []string{"Gopher", "Alpha", "Bravo", "Delta", "Grin", "Go"}
	if !basiccheck.EqualStringSlice(s, desiredSlice) {
		t.Errorf("SortStringsByLength didn't sort slice with smaller last and lexicographic order")
	}
}
