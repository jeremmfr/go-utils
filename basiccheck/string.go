package basiccheck

import (
	"strings"
)

// CheckStringHasOneOfPrefixes check if string has one of prefix list.
func CheckStringHasOneOfPrefixes(s string, list []string) bool {
	for _, item := range list {
		if strings.HasPrefix(s, item) {
			return true
		}
	}

	return false
}
