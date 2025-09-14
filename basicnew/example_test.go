package basicnew_test

import (
	"fmt"
	"sort"

	"github.com/jeremmfr/go-utils/basicnew"
)

func Example() {
	aMap := map[string]int{
		"bar": 1,
		"foo": 2,
	}

	keys := basicnew.MapKeys(aMap)
	sort.Strings(keys)
	fmt.Println(keys)
	// Output: [bar foo]
}

func ExampleMapKeys() {
	aMap := map[string]string{
		"bar": "baz",
		"foo": "baz",
	}

	keys := basicnew.MapKeys(aMap)
	sort.Strings(keys)
	fmt.Println(keys)
	// Output: [bar foo]
}

func ExampleMapValues() {
	aMap := map[string]string{
		"bar": "baz",
		"foo": "baz",
	}

	values := basicnew.MapValues(aMap)
	fmt.Println(values)
	// Output: [baz baz]
}

func ExampleSliceIntersect() {
	result := basicnew.SliceIntersect(
		[]string{"foo", "bar"},
		[]string{"baz", "foobar", "bar", "bar"},
	)

	fmt.Println(result)
	// Output: [bar]
}
