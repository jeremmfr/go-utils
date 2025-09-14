package basiccheck

import (
	"maps"
)

// SimilarSlice check if two slice is Similar:
// same length, same element (not necessarily in same order).
//
// Elements that appear multiple times must appear same times between the two slices.
func SimilarSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	aElems := make(map[T]int, len(a))
	for _, v := range a {
		aElems[v]++
	}

	bElems := make(map[T]int, len(a))
	for _, v := range b {
		if _, ok := aElems[v]; !ok {
			return false
		}
		bElems[v]++
	}

	return maps.Equal(aElems, bElems)
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
