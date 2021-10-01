package basiccheck_test

import (
	"testing"

	"github.com/jeremmfr/go-utils/basiccheck"
)

func TestStringHasOneOfPrefixes(t *testing.T) {
	prefStrings := []string{"foo", "bar"}

	if !basiccheck.StringHasOneOfPrefixes("foobaz", prefStrings) {
		t.Errorf("StringHasOneOfPrefixes not found prefix foo in footoo")
	}
	if basiccheck.StringHasOneOfPrefixes("fozbar", prefStrings) {
		t.Errorf("StringHasOneOfPrefixes found one of prefixes %v in fozbar", prefStrings)
	}
}
