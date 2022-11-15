package basiccheck_test

import (
	"fmt"
	"strings"

	"github.com/jeremmfr/go-utils/basiccheck"
)

func Example() {
	input := []string{"foo", "bar"}
	if basiccheck.InSlice("bar", input) {
		fmt.Printf("bar found in %v", input)
		// Output: bar found in [foo bar]
	} else {
		fmt.Printf("bar not found in %v", input)
	}
}

func ExampleInSlice() {
	input := []string{"foo", "bar"}
	if basiccheck.InSlice("bar", input) {
		fmt.Printf("bar found in %v", input)
		// Output: bar found in [foo bar]
	} else {
		fmt.Printf("bar not found in %v", input)
	}
}

func ExampleEqualSlice() {
	a := []string{"foo", "bar", "baz"}
	b := []string{"foo", "bar", "baz"}
	if basiccheck.EqualSlice(a, b) {
		fmt.Print("a == b")
		// Output: a == b
	} else {
		fmt.Print("a != b")
	}
}

func ExampleOneInSliceWith() {
	input := []string{"foo", "bar", "baz"}
	if basiccheck.OneInSliceWith(input, func(s string) bool {
		return strings.HasPrefix(s, "b")
	}) {
		fmt.Printf("one of %v has prefix 'b'", input)
		// Output: one of [foo bar baz] has prefix 'b'
	} else {
		fmt.Printf("no one of %v has prefix 'b'", input)
	}
}

func ExampleAllInSliceWith() {
	input := []string{"boo", "bar", "baz"}
	if basiccheck.AllInSliceWith(input, func(s string) bool {
		return strings.HasPrefix(s, "b")
	}) {
		fmt.Printf("all of %v has prefix 'b'", input)
		// Output: all of [boo bar baz] has prefix 'b'
	} else {
		fmt.Printf("one of %v has not prefix 'b'", input)
	}
}
