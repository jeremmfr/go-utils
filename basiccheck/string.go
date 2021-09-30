package basiccheck

import (
	"strings"
)

func CheckStringHasOneOfPrefixes(s string, list []string) bool {
	for _, item := range list {
		if strings.HasPrefix(s, item) {
			return true
		}
	}

	return false
}
