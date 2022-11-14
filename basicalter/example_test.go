package basicalter_test

import (
	"fmt"

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
