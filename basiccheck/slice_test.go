package basiccheck_test

import (
	"strings"
	"testing"

	"github.com/jeremmfr/go-utils/basiccheck"
)

func TestSimilarSlice(t *testing.T) {
	sliceA := []string{"foo", "bar", "baz"}
	sliceB := []string{"foo", "baz", "bar"}
	sliceC := []string{"foo", "bar"}

	if !basiccheck.SimilarSlice(sliceA, sliceB) {
		t.Errorf("SimilarSlice didn't find similar slice %v, %v", sliceA, sliceB)
	}
	if basiccheck.SimilarSlice(sliceA, sliceC) {
		t.Errorf("SimilarSlice found similar slice %v, %v", sliceA, sliceC)
	}

	sliceC = append(sliceC, "baz")
	if !basiccheck.SimilarSlice(sliceA, sliceC) {
		t.Errorf("SimilarSlice didn't find similar slice %v, %v", sliceA, sliceC)
	}

	sliceC[0] = "fo"
	if basiccheck.SimilarSlice(sliceA, sliceC) {
		t.Errorf("SimilarSlice found similar slice %v, %v", sliceA, sliceC)
	}

	sliceC = nil
	if basiccheck.SimilarSlice(sliceA, sliceC) {
		t.Errorf("SimilarSlice found similar slice %v, %v", sliceA, sliceC)
	}
	sliceA = nil
	if !basiccheck.SimilarSlice(sliceA, sliceC) {
		t.Errorf("SimilarSlice didn't find similar slice %v, %v", sliceA, sliceC)
	}
}

func TestAllInSliceWith(t *testing.T) {
	sliceOfString := []string{}

	if !basiccheck.AllInSliceWith(sliceOfString, nil) {
		t.Errorf("AllInSliceWith return false with nil find")
	}

	if !basiccheck.AllInSliceWith(sliceOfString, func(s string) bool {
		return strings.HasPrefix(s, "b")
	}) {
		t.Errorf("AllInSliceWith found prefix 'b' in empty slice")
	}

	sliceOfString = append(sliceOfString, []string{"baz", "bar"}...)
	// check if all elements contains 'a'
	if !basiccheck.AllInSliceWith(sliceOfString, func(s string) bool {
		return strings.Contains(s, "a")
	}) {
		t.Errorf("AllInSliceWith found one of %v without 'a'", sliceOfString)
	}

	sliceOfString = append(sliceOfString, "foo")
	// check if all elements contains 'a'
	if basiccheck.AllInSliceWith(sliceOfString, func(s string) bool {
		return strings.Contains(s, "a")
	}) {
		t.Errorf("AllInSliceWith didn't find 'foo' (without 'a') in %v", sliceOfString)
	}

	// check if all elements have lowercase letters
	if !basiccheck.AllInSliceWith(sliceOfString, func(s string) bool {
		return strings.ToLower(s) == s
	}) {
		t.Errorf("AllInSliceWith found a capital letter in one of %v", sliceOfString)
	}
}
