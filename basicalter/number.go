package basicalter

import (
	"math/bits"
)

// AbsoluteInt return absolute value of integer.
func AbsoluteInt(num int) int {
	y := num >> (bits.UintSize - 1)

	return (num ^ y) - y
}
