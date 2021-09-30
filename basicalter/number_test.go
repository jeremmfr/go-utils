package basicalter_test

import (
	"testing"

	"github.com/jeremmfr/go-utils/basicalter"
)

func TestAbsolutint(t *testing.T) {
	i := -2

	if basicalter.AbsoluteInt(i) < 0 {
		t.Errorf("AbsoluteInt return negative integer")
	}
	if basicalter.AbsoluteInt(basicalter.AbsoluteInt(i)) < 0 {
		t.Errorf("AbsoluteInt return negative integer")
	}
}
