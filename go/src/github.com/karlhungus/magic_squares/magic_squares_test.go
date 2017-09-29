package main

import "testing"

func TestMagicSquare(t *testing.T) {
	row1 := []int{8, 1, 6}
	row2 := []int{3, 5, 7}
	row3 := []int{4, 9, 2}
	expected := [][]int{row1, row2, row3}
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
