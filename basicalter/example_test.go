package basicalter_test

import (
	"fmt"
	"strings"

	"github.com/jeremmfr/go-utils/basicalter"
)

func Example() {
	input := []string{"foo", "bar", "foo"}
	output := basicalter.UniqueInSlice(input)
	fmt.Printf("%v", output)
	// Output: [foo bar]
}

func ExampleUniqueInSlice() {
	input := []string{"foo", "bar", "foo"}
	output := basicalter.UniqueInSlice(input)
	fmt.Printf("%v", output)
	// Output: [foo bar]
}

func ExampleDelInSlice() {
	input := []string{"foo", "baz", "bar"}
	output := basicalter.DelInSlice("baz", input)
	fmt.Printf("%v", output)
	// Output: [foo bar]
}

func ExampleFilterInSliceWith() {
	input := []string{"foo", "baz", "bar", "baz"}
	output := basicalter.FilterInSliceWith(input, func(s string) bool {
		return strings.HasPrefix(s, "ba")
	})
	fmt.Printf("%v", output)
	// Output: [baz bar baz]
}

func ExampleReverseSlice() {
	input := []string{"foo", "baz", "bar"}
	basicalter.ReverseSlice(input)
	fmt.Printf("%v", input)
	// Output: [bar baz foo]
}

func ExampleReplaceInSliceWith() {
	input := []string{"Foo", "Bar", "Baz"}
	basicalter.ReplaceInSliceWith(input, strings.ToLower)
	fmt.Printf("%v", input)
	// Output: [foo bar baz]
}
