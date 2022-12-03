package basicalter_test

import (
	"fmt"
	"testing"

	"github.com/jeremmfr/go-utils/basicalter"
)

func TestMergeMaps(t *testing.T) {
	var m map[string]string
	m2 := map[string]string{"foo": "bar"}
	// try merge in nil map
	basicalter.MergeMaps(m, false, m2)

	// try merge in empty map
	m = make(map[string]string)
	m["baz"] = "baz"
	basicalter.MergeMaps(m, false, m2)
	if len(m) != 2 {
		t.Errorf("MergeMaps doesn't merge correctly")
	}
	if v := fmt.Sprintf("%v", m); v != "map[baz:baz foo:bar]" {
		t.Errorf("MergeMaps doesn't merge correctly: %s", v)
	}

	// try merge in with same keys in two maps (srcs)
	m["foo"] = "baz"
	mm := make(map[string]string)
	basicalter.MergeMaps(mm, false, m2, m)
	if len(mm) != 2 {
		t.Errorf("MergeMaps doesn't merge correctly")
	}
	if v := fmt.Sprintf("%v", mm); v != "map[baz:baz foo:bar]" {
		t.Errorf("MergeMaps doesn't merge correctly: %s", v)
	}

	basicalter.MergeMaps(mm, true, m)
	if len(mm) != 2 {
		t.Errorf("MergeMaps doesn't merge correctly")
	}
	if v := fmt.Sprintf("%v", mm); v != "map[baz:baz foo:baz]" {
		t.Errorf("MergeMaps doesn't merge correctly: %s", v)
	}
}
