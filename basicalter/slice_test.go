package basicalter_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/jeremmfr/go-utils/basicalter"
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

func TestDelEmptyStrings(t *testing.T) {
	sliceOfString := []string{"foo", "", "bar"}

	if v := basicalter.DelEmptyStrings(sliceOfString); len(v) != 2 {
		t.Errorf("DelEmptyStrings didn't remove empty string: %v", v)
	} else if !slices.Equal(v, []string{"foo", "bar"}) {
		t.Errorf("DelEmptyStrings didn't remove empty string: %v", v)
	}
}

func TestDelInSlice(t *testing.T) {
	sliceOfString := []string{"foo", "baz", "bar", "baz"}

	if v := basicalter.DelInSlice("baz", sliceOfString); len(v) != 2 {
		t.Errorf("DelInSlice didn't remove 'baz': %v", v)
	} else if !slices.Equal(v, []string{"foo", "bar"}) {
		t.Errorf("DelInSlice didn't remove 'baz': %v", v)
	}
}

func TestFilterInSliceWith(t *testing.T) {
	sliceOfString := []string{"foo", "baz", "bar", "baz"}

	if v := basicalter.FilterInSliceWith(sliceOfString, nil); len(v) != len(sliceOfString) {
		t.Errorf("FilterInSliceWith didn't return same slice with nil filter: %v", v)
	}

	if v := basicalter.FilterInSliceWith(sliceOfString, func(s string) bool {
		return strings.HasPrefix(s, "ba")
	}); len(v) != 3 {
		t.Errorf("FilterInSliceWith didn't remove foo (without prefix 'ba'): %v", v)
	} else if !slices.Equal(v, []string{"baz", "bar", "baz"}) {
		t.Errorf("FilterInSliceWith didn't remove foo (without prefix 'ba'): %v", v)
	}

	var nilSlice []string
	if v := basicalter.FilterInSliceWith(nilSlice, func(_ string) bool {
		return true
	}); v != nil {
		t.Errorf("FilterInSliceWith didn't return nil slice with nil input slice")
	}
}

func TestReverseSlice(t *testing.T) {
	sliceOfString := []string{"foo", "baz", "bar", "World", "Hello"}

	basicalter.ReverseSlice(sliceOfString)

	desiredStringSlice := []string{"Hello", "World", "bar", "baz", "foo"}
	if !slices.Equal(sliceOfString, desiredStringSlice) {
		t.Errorf("ReverseSlice didn't reverse slice: %v expected %v", sliceOfString, desiredStringSlice)
	}

	sliceOfInt64 := []int64{3, 2, 1, 0}

	basicalter.ReverseSlice(sliceOfInt64)

	desiredInt64Slice := []int64{0, 1, 2, 3}
	if !slices.Equal(sliceOfInt64, desiredInt64Slice) {
		t.Errorf("ReverseSlice didn't reverse slice: %v expected %v", sliceOfInt64, desiredInt64Slice)
	}
}

func TestSortStringsByLengthInc(t *testing.T) {
	s := []string{"Bravo", "Gopher", "Alpha", "Go", "Grin", "Delta"}

	basicalter.SortStringsByLengthInc(s)

	desiredSlice := []string{"Go", "Grin", "Alpha", "Bravo", "Delta", "Gopher"}
	if !slices.Equal(s, desiredSlice) {
		t.Errorf("SortStringsByLength didn't sort slice with smaller first and lexicographic order")
	}
}

func TestSortStringsByLengthDec(t *testing.T) {
	s := []string{"Bravo", "Gopher", "Alpha", "Go", "Grin", "Delta"}

	basicalter.SortStringsByLengthDec(s)

	desiredSlice := []string{"Gopher", "Alpha", "Bravo", "Delta", "Grin", "Go"}
	if !slices.Equal(s, desiredSlice) {
		t.Errorf("SortStringsByLength didn't sort slice with smaller last and lexicographic order")
	}
}

func TestReplaceInSliceWith(t *testing.T) {
	sliceOfString := []string{"Foo", "Bar", "Baz"}

	basicalter.ReplaceInSliceWith(sliceOfString, strings.ToLower)

	if !slices.Equal(sliceOfString, []string{"foo", "bar", "baz"}) {
		t.Errorf("ReplaceInSliceWith didn't replace all strings in slice "+
			"with the lowercase version: %v", sliceOfString)
	}

	// test empty slice
	basicalter.ReplaceInSliceWith([]string{}, strings.ToLower)
}
