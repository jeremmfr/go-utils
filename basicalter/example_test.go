package basicalter_test

import (
	"fmt"

	"github.com/jeremmfr/go-utils/basicalter"
)

func Example() {
	sliceOfString := []string{"foo", "bar", "foo"}

	fmt.Printf("%v", basicalter.UniqueStrings(sliceOfString))
	// Output: [foo bar]
}
