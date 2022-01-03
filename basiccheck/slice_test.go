package basiccheck_test

import (
	"testing"

	"github.com/jeremmfr/go-utils/basiccheck"
)

func TestStringInSlice(t *testing.T) {
	sliceOfString := []string{"foo", "bar"}

	if !basiccheck.StringInSlice("bar", sliceOfString) {
		t.Errorf("CheckIfStringInSlice didn't find bar in %v", sliceOfString)
	}
	if basiccheck.StringInSlice("baz", sliceOfString) {
		t.Errorf("CheckIfStringInSlice found baz in %v", sliceOfString)
	}
}

func TestEqualStringSlice(t *testing.T) {
	sliceA := []string{"foo", "bar", "baz"}
	sliceB := []string{"foo", "baz", "bar"}
	sliceC := []string{"foo", "bar"}

	if basiccheck.EqualStringSlice(sliceA, sliceB) {
		t.Errorf("EqualStringSlice found equal slice %v, %v", sliceA, sliceB)
	}
	if basiccheck.EqualStringSlice(sliceA, sliceC) {
		t.Errorf("EqualStringSlice found equal slice %v, %v", sliceA, sliceC)
	}

	sliceC = append(sliceC, "baz")
	if !basiccheck.EqualStringSlice(sliceA, sliceC) {
		t.Errorf("EqualStringSlice didn't find equal slice %v, %v", sliceA, sliceC)
	}
}

func TestIntInSlice(t *testing.T) {
	sliceOfInt := []int{2, 4, 10}

	if !basiccheck.IntInSlice(4, sliceOfInt) {
		t.Errorf("IntInSlice didn't find 4 in %v", sliceOfInt)
	}
	if basiccheck.IntInSlice(6, sliceOfInt) {
		t.Errorf("IntInSlice found 6 in %v", sliceOfInt)
	}
}

func TestInt64InSlice(t *testing.T) {
	sliceOfInt := []int64{2, 4, 10}

	if !basiccheck.Int64InSlice(4, sliceOfInt) {
		t.Errorf("Int64InSlice didn't find 4 in %v", sliceOfInt)
	}
	if basiccheck.Int64InSlice(6, sliceOfInt) {
		t.Errorf("Int64InSlice found 6 in %v", sliceOfInt)
	}
}
