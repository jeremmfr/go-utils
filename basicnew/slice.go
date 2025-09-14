package basicnew

// SliceIntersect generate a new slice with
// elements present in both slices.
//
// If an element is present multiple times in one of the argument lists,
// it will be present only once in the result.
func SliceIntersect[T comparable](a, b []T) []T {
	aElems := make(map[T]bool, len(a))
	for _, v := range a {
		aElems[v] = false
	}

	resultCap := len(a)
	if v := len(b); v < resultCap {
		resultCap = v
	}
	result := make([]T, 0, resultCap)

	for _, v := range b {
		if set, ok := aElems[v]; ok && !set {
			result = append(result, v)
			aElems[v] = true
		}
	}

	return result
}
