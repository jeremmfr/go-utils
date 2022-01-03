package basiccheck_test

import (
	"testing"

	"github.com/jeremmfr/go-utils/basiccheck"
)

func TestStringHasOneOfPrefixes(t *testing.T) {
	prefStrings := []string{"foo", "bar"}

	if !basiccheck.StringHasOneOfPrefixes("foobaz", prefStrings) {
		t.Errorf("StringHasOneOfPrefixes didn't find prefix foo in foobaz")
	}
	if basiccheck.StringHasOneOfPrefixes("fozbar", prefStrings) {
		t.Errorf("StringHasOneOfPrefixes found one of prefixes %v in fozbar", prefStrings)
	}
}
