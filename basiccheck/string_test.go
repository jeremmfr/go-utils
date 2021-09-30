package basiccheck_test

import (
	"testing"

	"github.com/jeremmfr/go-utils/basiccheck"
)

func TestCheckStringHasOneOfPrefixes(t *testing.T) {
	prefStrings := []string{"foo", "bar"}

	if !basiccheck.CheckStringHasOneOfPrefixes("foobaz", prefStrings) {
		t.Errorf("CheckStringHasOneOfPrefixes not found prefix foo in footoo")
	}
	if basiccheck.CheckStringHasOneOfPrefixes("fozbar", prefStrings) {
		t.Errorf("CheckStringHasOneOfPrefixes found one of prefixes %v in fozbar", prefStrings)
	}
}
