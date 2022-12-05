package basicnew_test

import (
	"fmt"
	"testing"

	"github.com/jeremmfr/go-utils/basicnew"
)

func TestMapKeys(t *testing.T) {
	map1 := make(map[string]struct{})
	map1["bar"] = struct{}{}
	map1["foo"] = struct{}{}

	keys := basicnew.MapKeys(map1)

	if v := fmt.Sprintf("%v", keys); v != "[bar foo]" && v != "[foo bar]" {
		t.Errorf("MapKeys doesn't return expected slice [bar foo] or [foo bar]: %s", v)
	}
}

func TestMapValues(t *testing.T) {
	map1 := make(map[int]string)
	map1[1] = "bar"
	map1[0] = "foo"

	values := basicnew.MapValues(map1)

	if v := fmt.Sprintf("%v", values); v != "[bar foo]" && v != "[foo bar]" {
		t.Errorf("MapValues doesn't return expected slice [bar foo] or [foo bar]: %s", v)
	}
}
