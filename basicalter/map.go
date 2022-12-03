package basicalter

// MergeMaps copies all key/value pairs in 'srcs' maps to 'dst' map.
// Set 'overwrite' to true to allow copy value when the key is already set
// in 'dst' or previous element of 'srcs'.
// If 'dst' is nil, MergeMaps is a no-op.
func MergeMaps[K comparable, V any, M ~map[K]V](dst M, overwrite bool, srcs ...M) {
	if dst == nil {
		return
	}
	for _, m := range srcs {
		for k, v := range m {
			if _, ok := dst[k]; ok {
				if overwrite {
					dst[k] = v
				}
			} else {
				dst[k] = v
			}
		}
	}
}
