package array_test

import (
	"testing"
	"array"
)
func TestMin(t *testing.T) {

	_min := array.Min(1,3,5,7,9)
	_expect := 1
	if _min != _expect {
		t.Error("expected", _expect, "got", _min)
	}
}