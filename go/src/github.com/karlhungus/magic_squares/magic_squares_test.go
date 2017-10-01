package main

import "testing"

func TestMagicSquare(t *testing.T) {
	expected := [][]int{
		[]int{8, 1, 6},
		[]int{3, 5, 7},
		[]int{4, 9, 2},
	}

	result := magicSquare(3)

	if len(result) != 3 {
		t.Error("result is not long enough got: %v", result)
	}

	for i := range expected {
		if len(result[i]) != len(expected[i]) {
			t.Error("result is not long enough got: %v", result)
		}
		for j := range expected[i] {
			if expected[i][j] != result[i][j] {
				t.Error("Expected:\n %v for magic square got %v", expected, result)
			}
		}
	}
}
