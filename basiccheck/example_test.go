package basiccheck_test

import (
	"fmt"

	"github.com/jeremmfr/go-utils/basiccheck"
)

func Example() {
	sliceOfString := []string{"foo", "bar"}

	if basiccheck.StringInSlice("bar", sliceOfString) {
		fmt.Printf("bar found in %v", sliceOfString)
		// Output: bar found in [foo bar]
	} else {
		fmt.Printf("bar not found in %v", sliceOfString)
	}
}
