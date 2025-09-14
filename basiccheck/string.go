package basiccheck

import (
	"slices"
	"strings"
)

// StringHasOneOfPrefixes check if string has one of prefix list.
func StringHasOneOfPrefixes(s string, list []string) bool {
	return slices.ContainsFunc(list, func(v string) bool {
		return strings.HasPrefix(s, v)
	})
}

// StringHasOneOfSuffixes check if string has one of suffix list.
func StringHasOneOfSuffixes(s string, list []string) bool {
	return slices.ContainsFunc(list, func(v string) bool {
		return strings.HasSuffix(s, v)
	})
}
