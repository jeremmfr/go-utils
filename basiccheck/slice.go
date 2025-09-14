package basiccheck

import "slices"

// SimilarSlice check if two slice is Similar:
// same length, same element (not necessarily in same order).
func SimilarSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range a {
		if !slices.Contains(b, v) {
			return false
		}
	}

	return true
}

// AllInSliceWith check if all elements in a slice
// return true with the function 'valid' passed in arguments.
//
// If 'valid' is nil, return true.
func AllInSliceWith[T any](list []T, valid func(T) bool) bool {
	if valid == nil {
		return true
	}
	for _, v := range list {
		if !valid(v) {
			return false
		}
	}

	return true
}
