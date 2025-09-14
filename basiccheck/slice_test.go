package basiccheck_test

import (
	"strings"
	"testing"

	"github.com/jeremmfr/go-utils/basiccheck"
)

func TestSimilarSlice(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a            []string
		b            []string
		expectResult bool
	}

	tests := map[string]testCase{
		"nil|nil": {
			a:            nil,
			b:            nil,
			expectResult: true,
		},
		"foo_bar_baz": {
			a:            []string{"foo", "bar", "baz"},
			b:            []string{"foo", "baz", "bar"},
			expectResult: true,
		},
		"foo_bar+baz": {
			a:            []string{"foo", "bar", "baz"},
			b:            []string{"foo", "bar"},
			expectResult: false,
		},
		"foo_bar_baz_2": {
			a:            []string{"foo", "bar", "baz"},
			b:            []string{"foo", "bar", "baz"},
			expectResult: true,
		},
		"bar_baz+fo": {
			a:            []string{"foo", "bar", "baz"},
			b:            []string{"fo", "bar", "baz"},
			expectResult: false,
		},
		"foo_bar_baz|nil": {
			a:            []string{"foo", "bar", "baz"},
			b:            nil,
			expectResult: false,
		},
		"foo*2_bar_baz": {
			a:            []string{"foo", "bar", "baz", "foo"},
			b:            []string{"foo", "bar", "baz", "foo"},
			expectResult: true,
		},
		"foo*2_bar*2_baz*2": {
			a:            []string{"foo", "bar", "baz", "foo", "bar"},
			b:            []string{"foo", "bar", "baz", "foo", "baz"},
			expectResult: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if b := basiccheck.SimilarSlice(test.a, test.b); b != test.expectResult {
				t.Errorf("got unexpected result: want %v, got %v", test.expectResult, b)
			}
		})
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
