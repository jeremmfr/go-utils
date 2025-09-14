package basiccheck

import "slices"

// InSlice check if an element is present in a slice.
//
// Deprecated: use slices.Contains() from the standard 'slices' library instead.
func InSlice[T comparable](elem T, list []T) bool {
	return OneInSliceWith(list, func(v T) bool {
		return v == elem
	})
}

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

// OneInSliceWith check if at least one element in a slice
// returns true with the function 'find' passed in arguments.
//
// If 'find' is nil, return false.
//
// Deprecated: use slices.ContainsFunc() from the standard 'slices' library instead.
func OneInSliceWith[T any](list []T, find func(T) bool) bool {
	if find == nil {
		return false
	}
	for _, v := range list {
		if find(v) {
			return true
		}
	}

	return false
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
