package basiccheck

import (
	"maps"
)

// InSlice check if an element is present in a slice.
//
// Deprecated: use slices.Contains() from the standard 'slices' library instead.
func InSlice[T comparable](elem T, list []T) bool {
	return OneInSliceWith(list, func(v T) bool {
		return v == elem
	})
}

// EqualSlice check if two slice is Equal: same length, same element in same order.
//
// Deprecated: use slices.Equal() from the standard 'slices' library instead.
func EqualSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

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
