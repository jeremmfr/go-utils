package basicalter

import "strings"

// CutPrefixInString rewrites the value of 's' without the provided
// leading 'prefix' string and reports whether it found the prefix.
// If 's' is nil or if its value doesn't start with 'prefix', CutPrefixInString
// doesn't rewrite the value of 's' and returns false.
// If 'prefix' is the empty string, CutPrefixInString returns true without rewriting.
func CutPrefixInString(s *string, prefix string) bool {
	if s == nil {
		return false
	}
	if len(prefix) == 0 {
		return true
	}
	if !strings.HasPrefix(*s, prefix) {
		return false
	}
	str := *s
	*s = str[len(prefix):]

	return true
}

// CutSuffixInString rewrites the value of 's' without the provided
// ending 'suffix' string and reports whether it found the suffix.
// If 's' is nil or if its value doesn't end with 'suffix', CutSuffixInString
// doesn't rewrite the value of 's' and returns false.
// If 'suffix' is the empty string, CutSuffixInString returns true without rewriting.
func CutSuffixInString(s *string, suffix string) bool {
	if s == nil {
		return false
	}
	if len(suffix) == 0 {
		return true
	}
	if !strings.HasSuffix(*s, suffix) {
		return false
	}
	str := *s
	*s = str[:len(str)-len(suffix)]

	return true
}
