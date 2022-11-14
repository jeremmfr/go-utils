package basicalter_test

import (
	"strings"
	"testing"

	"github.com/jeremmfr/go-utils/basicalter"
	"github.com/jeremmfr/go-utils/basiccheck"
)

func TestUniqueInSlice(t *testing.T) {
	sliceOfString := []string{"foo", "bar"}
	if len(basicalter.UniqueInSlice(sliceOfString)) != len(sliceOfString) {
		t.Errorf("UniqueInSlice returned bad slice with slice without duplicate entry")
	}

	sliceOfString = append(sliceOfString, "foo")
	if v := basicalter.UniqueInSlice(sliceOfString); len(v) == len(sliceOfString) {
		t.Errorf("UniqueInSlice didn't remove duplicate foo in slice: %v", v)
	}

	// test empty slice
	_ = len(basicalter.UniqueInSlice([]string{}))

	// test nil slice
	var nilSlice []string
	if basicalter.UniqueInSlice(nilSlice) != nil {
		t.Errorf("UniqueInSlice didn't return nil slice with nil input slice")
	}

	sliceOfInt := []int{1, 2}
	if len(basicalter.UniqueInSlice(sliceOfInt)) != len(sliceOfInt) {
		t.Errorf("UniqueInSlice returned bad slice with slice without duplicate entry")
	}

	sliceOfInt = append(sliceOfInt, 1)
	if v := basicalter.UniqueInSlice(sliceOfInt); len(v) == len(sliceOfInt) {
		t.Errorf("UniqueInSlice didn't remove duplicate 1 in slice: %v", v)
	}

	// test slice with one element
	if v := basicalter.UniqueInSlice([]float64{1.2}); len(v) != 1 || v[0] != 1.2 {
		t.Errorf("UniqueInSlice didn't return the same slice with one element")
	}
}

func TestUniqueStrings(t *testing.T) {
	sliceOfString := []string{"foo", "bar"}

	if len(basicalter.UniqueStrings(sliceOfString)) != len(sliceOfString) {
		t.Errorf("UniqueStrings returned bad slice with slice without duplicate entry")
	}

	sliceOfString = append(sliceOfString, "foo")

	if v := basicalter.UniqueStrings(sliceOfString); len(v) == len(sliceOfString) {
		t.Errorf("UniqueStrings didn't remove duplicate foo in slice: %v", v)
	}
	// test empty list
	_ = len(basicalter.UniqueStrings([]string{}))
}

func TestDelEmptyStrings(t *testing.T) {
	sliceOfString := []string{"foo", "", "bar"}

	if v := basicalter.DelEmptyStrings(sliceOfString); len(v) != 2 {
		t.Errorf("DelEmptyStrings didn't remove empty string: %v", v)
	} else if !basiccheck.EqualSlice(v, []string{"foo", "bar"}) {
		t.Errorf("DelEmptyStrings didn't remove empty string: %v", v)
	}
}

func TestDelStringInSlice(t *testing.T) {
	sliceOfString := []string{"foo", "baz", "bar", "baz"}

	if v := basicalter.DelStringInSlice("baz", sliceOfString); len(v) != 2 {
		t.Errorf("DelStringInSlice didn't remove 'baz': %v", v)
	} else if !basiccheck.EqualSlice(v, []string{"foo", "bar"}) {
		t.Errorf("DelStringInSlice didn't remove 'baz': %v", v)
	}
}

func TestFilterInSliceWith(t *testing.T) {
	sliceOfString := []string{"foo", "baz", "bar", "baz"}

	if v := basicalter.FilterInSliceWith(sliceOfString, func(s string) bool {
		return strings.HasPrefix(s, "ba")
	}); len(v) != 3 {
		t.Errorf("FilterInSliceWith didn't remove foo (without prefix 'ba'): %v", v)
	} else if !basiccheck.EqualSlice(v, []string{"baz", "bar", "baz"}) {
		t.Errorf("FilterInSliceWith didn't remove foo (without prefix 'ba'): %v", v)
	}

	var nilSlice []string
	if v := basicalter.FilterInSliceWith(nilSlice, func(s string) bool {
		return true
	}); v != nil {
		t.Errorf("FilterInSliceWith didn't return nil slice with nil input slice")
	}
}

func TestFilterStringsWith(t *testing.T) {
	sliceOfString := []string{"foo", "baz", "bar", "baz"}

	if v := basicalter.FilterStringsWith(sliceOfString, func(s string) bool {
		return strings.HasPrefix(s, "ba")
	}); len(v) != 3 {
		t.Errorf("FilterStringsWith didn't remove foo (without prefix 'ba'): %v", v)
	} else if !basiccheck.EqualSlice(v, []string{"baz", "bar", "baz"}) {
		t.Errorf("FilterStringsWith didn't remove foo (without prefix 'ba'): %v", v)
	}
}

func TestReverseSlice(t *testing.T) {
	sliceOfString := []string{"foo", "baz", "bar", "World", "Hello"}

	basicalter.ReverseStrings(sliceOfString)

	desiredSlice := []string{"Hello", "World", "bar", "baz", "foo"}
	if !basiccheck.EqualSlice(sliceOfString, desiredSlice) {
		t.Errorf("ReverseStrings didn't reverse slice: %v expected %v", sliceOfString, desiredSlice)
	}
}

func TestSortStringsByLengthInc(t *testing.T) {
	s := []string{"Bravo", "Gopher", "Alpha", "Go", "Grin", "Delta"}

	basicalter.SortStringsByLengthInc(s)

	desiredSlice := []string{"Go", "Grin", "Alpha", "Bravo", "Delta", "Gopher"}
	if !basiccheck.EqualSlice(s, desiredSlice) {
		t.Errorf("SortStringsByLength didn't sort slice with smaller first and lexicographic order")
	}
}

func TestSortStringsByLengthDec(t *testing.T) {
	s := []string{"Bravo", "Gopher", "Alpha", "Go", "Grin", "Delta"}

	basicalter.SortStringsByLengthDec(s)

	desiredSlice := []string{"Gopher", "Alpha", "Bravo", "Delta", "Grin", "Go"}
	if !basiccheck.EqualSlice(s, desiredSlice) {
		t.Errorf("SortStringsByLength didn't sort slice with smaller last and lexicographic order")
	}
}

func TestReplaceStringsWith(t *testing.T) {
	sliceOfString := []string{"Foo", "Bar", "Baz"}

	basicalter.ReplaceStringsWith(sliceOfString, strings.ToLower)

	if !basiccheck.EqualSlice(sliceOfString, []string{"foo", "bar", "baz"}) {
		t.Errorf("ReplaceStringsWith didn't replace all strings in slice "+
			"with the lowercase version: %v", sliceOfString)
	}

	// test empty slice
	basicalter.ReplaceStringsWith([]string{}, strings.ToLower)
}
