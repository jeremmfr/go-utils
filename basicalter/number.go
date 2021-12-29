package basicalter

import (
	"math/bits"
)

// AbsoluteInt return absolute value of integer.
//
// If the 'num' is the minimum value of integers, the function panic
// due to the lack of a corresponding positive value with integer type.
func AbsoluteInt(num int) int {
	if num == -1<<(bits.UintSize-1) {
		panic("the absolute of the minimum value of integers is impossible (overflows int)")
	}

	y := num >> (bits.UintSize - 1)

	return (num ^ y) - y
}
