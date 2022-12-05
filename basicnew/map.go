package basicnew

// MapKeys generate a new slice with all keys of a map.
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, len(m))
	i := 0
	for k := range m {
		r[i] = k
		i++
	}

	return r
}

// MapValues generate a new slice with all values of a map.
func MapValues[K comparable, V any](m map[K]V) []V {
	r := make([]V, len(m))
	i := 0
	for _, v := range m {
		r[i] = v
		i++
	}

	return r
}
