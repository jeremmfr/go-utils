package basicalter

// AbsoluteInt return absolute value of integer.
func AbsoluteInt(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
