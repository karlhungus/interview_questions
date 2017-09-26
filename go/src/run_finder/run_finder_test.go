package main

import "testing"

func TestLengthCounter(t *testing.T) {
	total := lengthFinder(3, "111111000")
	if total != 5 {
		t.Error("Expected 5, for strings of length 3 in '1111000 got ", total)
	}
}
