package basicalter_test

import (
	"testing"

	"github.com/jeremmfr/go-utils/basicalter"
)

func TestAbsolutInt(t *testing.T) {
	i := -2

	if v := basicalter.AbsoluteInt(i); v < 0 {
		t.Errorf("AbsoluteInt(%d) return negative integer: %d", i, v)
	}
	if v := basicalter.AbsoluteInt(basicalter.AbsoluteInt(i)); v < 0 {
		t.Errorf("AbsoluteInt(AbsoluteInt(%d)) return negative integer: %d", i, v)
	}
}
