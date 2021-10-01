package basiccheck

import (
	"strings"
)

// StringHasOneOfPrefixes check if string has one of prefix list.
func StringHasOneOfPrefixes(s string, list []string) bool {
	for _, item := range list {
		if strings.HasPrefix(s, item) {
			return true
		}
	}

	return false
}
