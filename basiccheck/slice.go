package basiccheck

// InSlice check if an element is present in a slice.
func InSlice[T comparable](elem T, list []T) bool {
	return OneInSliceWith(list, func(v T) bool {
		return v == elem
	})
}

// EqualSlice check if two slice is Equal: same length, same element in same order.
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

// OneInSliceWith check if at least one element in a slice
// returns true with the function 'find' passed in arguments.
func OneInSliceWith[T any](list []T, find func(T) bool) bool {
	for _, v := range list {
		if find(v) {
			return true
		}
	}

	return false
}

// AllInSliceWith check if all elements in a slice
// return true with the function 'valid' passed in arguments.
func AllInSliceWith[T any](list []T, valid func(T) bool) bool {
	for _, v := range list {
		if !valid(v) {
			return false
		}
	}

	return true
}
