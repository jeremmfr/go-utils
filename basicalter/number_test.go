package basicalter_test

import (
	"math/bits"
	"testing"

	"github.com/jeremmfr/go-utils/basicalter"
)

const (
	maxInt = 1<<(bits.UintSize-1) - 1
	minInt = -1 << (bits.UintSize - 1)
)

func TestAbsolutInt(t *testing.T) {
	i := -2

	if v := basicalter.AbsoluteInt(i); v < 0 {
		t.Errorf("AbsoluteInt(%d) return negative integer: %d", i, v)
	}
	if v := basicalter.AbsoluteInt(basicalter.AbsoluteInt(i)); v < 0 {
		t.Errorf("AbsoluteInt(AbsoluteInt(%d)) return negative integer: %d", i, v)
	}

	i = minInt + 1
	if v := basicalter.AbsoluteInt(i); v != maxInt {
		t.Errorf("AbsoluteInt(%d) not return expected integer %d, found %d", i, maxInt, v)
	}
}

func TestAbsolutIntPanic(t *testing.T) {
	defer func() { _ = recover() }()
	basicalter.AbsoluteInt(minInt)
	t.Errorf("AbsoluteInt doesn't panic with the minimum value of integers")
}
