package tests

import (
	"testing"

	programming "../src"
)

func TestSumForMultipleValues(t *testing.T) {
	type test struct {
		a      int
		b      int
		result int
	}

	tests := []test{
		test{2, 3, 5},
		test{9, 5, 14},
		test{9, -4, 5},
		test{0, 0, 0},
	}

	for _, v := range tests {
		x := programming.Sum(v.a, v.b)
		if x != v.result {
			t.Errorf("Expected %d, got %d", v.result, x)
		}
	}
}
