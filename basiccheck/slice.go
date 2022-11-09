package basiccheck

// InSlice check if element is present in a slice.
func InSlice[T comparable](elm T, list []T) bool {
	return OneOfSliceWith(list, func(v T) bool {
		return v == elm
	})
}

// OneOfSliceWith check if at least one element in a slice
// returns true with the function 'find' passed in arguments.
func OneOfSliceWith[T any](list []T, find func(T) bool) bool {
	for _, v := range list {
		if find(v) {
			return true
		}
	}

	return false
}

// StringInSlice check if a string is present in a slice of string.
//
// Deprecated: use InSlice() instead.
func StringInSlice(str string, list []string) bool {
	return OneOfStringsWith(list, func(s string) bool {
		return s == str
	})
}

// EqualStringSlice check if two slice of string is Equal: same length, same element in same order.
func EqualStringSlice(a, b []string) bool {
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

// IntInSlice check if int is present in slice of int.
//
// Deprecated: use InSlice() instead.
func IntInSlice(num int, list []int) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}

	return false
}

// Int64InSlice check if int64 is present in slice of int64.
//
// Deprecated: use InSlice() instead.
func Int64InSlice(num int64, list []int64) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}

	return false
}

// OneOfStringsWith check if at least one string in a slice
// returns true with the function 'find' passed in arguments.
//
// Deprecated: use OneOfSliceWith() instead.
func OneOfStringsWith(list []string, find func(string) bool) bool {
	for _, v := range list {
		if find(v) {
			return true
		}
	}

	return false
}

// AllStringsWith check if all strings in a slice
// return true with the function 'valid' passed in arguments.
func AllStringsWith(list []string, valid func(string) bool) bool {
	for _, v := range list {
		if !valid(v) {
			return false
		}
	}

	return true
}
