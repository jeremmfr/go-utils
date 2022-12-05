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

func ExampleAbsoluteInt() {
	input := -10
	output := basicalter.AbsoluteInt(input)
	fmt.Printf("%d", output)
	// Output: 10
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

func ExampleSortStringsByLengthInc() {
	input := []string{"a1", "a10", "a11", "a100", "a2", "a3"}
	basicalter.SortStringsByLengthInc(input)
	fmt.Printf("%v", input)
	// Output: [a1 a2 a3 a10 a11 a100]
}

func ExampleSortStringsByLengthDec() {
	input := []string{"a1", "a10", "a11", "a100", "a2", "a3"}
	basicalter.SortStringsByLengthDec(input)
	fmt.Printf("%v", input)
	// Output: [a100 a10 a11 a1 a2 a3]
}

func ExampleReplaceInSliceWith() {
	input := []string{"Foo", "Bar", "Baz"}
	basicalter.ReplaceInSliceWith(input, strings.ToLower)
	fmt.Printf("%v", input)
	// Output: [foo bar baz]
}

func ExampleMergeMaps() {
	m1 := map[string]string{"foo": "baz"}
	m2 := map[string]string{"bar": "baz"}
	basicalter.MergeMaps(m1, false, m2)
	fmt.Printf("%+v", m1)
	// Output: map[bar:baz foo:baz]
}
