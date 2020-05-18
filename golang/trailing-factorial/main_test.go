package main

import (
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	ret := trailingZeroes(3)
	if ret != 0 {
		t.Errorf("trailingZeroes(3) returned %d; wanted 0", ret)
	}

	ret = trailingZeroes(5)
	if ret != 1 {
		t.Errorf("trailingZeroes(5) returned %d; wanted 1", ret)
	}

	ret = trailingZeroes(10)
	if ret != 2 {
		t.Errorf("trailingZeroes(10) returned %d; wanted 2", ret)
	}
}
