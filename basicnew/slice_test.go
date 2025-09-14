package basicnew_test

import (
	"slices"
	"testing"

	"github.com/jeremmfr/go-utils/basicnew"
)

func TestSliceIntersect(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a            []string
		b            []string
		expectResult []string
	}

	tests := map[string]testCase{
		"nil": {
			a:            nil,
			b:            nil,
			expectResult: []string{},
		},
		"with_0_common": {
			a:            []string{"a", "b"},
			b:            []string{"d", "c"},
			expectResult: []string{},
		},
		"with_1_common": {
			a:            []string{"a", "b"},
			b:            []string{"b", "c"},
			expectResult: []string{"b"},
		},
		"with_1_common_multiple": {
			a:            []string{"a", "b", "b"},
			b:            []string{"b", "c"},
			expectResult: []string{"b"},
		},
		"with_1_common_multiple_b": {
			a:            []string{"a", "b", "b"},
			b:            []string{"b", "c", "b"},
			expectResult: []string{"b"},
		},
		"with_2_commons": {
			a:            []string{"a", "b", "zz"},
			b:            []string{"b", "c", "zz"},
			expectResult: []string{"b", "zz"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			result := basicnew.SliceIntersect(test.a, test.b)
			if !slices.Equal(result, test.expectResult) {
				t.Errorf("got unexpected result: want %v, got %v", test.expectResult, result)
			}
		})
	}
}
