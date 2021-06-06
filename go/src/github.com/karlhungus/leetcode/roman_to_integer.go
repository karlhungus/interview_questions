package main

import "fmt"

func main() {
}

func romanToInt(s string) int {
	result := 0
	for i, c := range s {
		switch c {
		case 'I':
			if isNegative(c, i, s) {
				result = result - 1
			} else {
				result = result + 1
			}

		case 'V':

			result = result + 5
		case 'X':
			if isNegative(c, i, s) {
				result = result - 10
			} else {
				result = result + 10
			}
		case 'L':

			result = result + 50

		case 'C':
			if isNegative(c, i, s) {
				result = result - 100
			} else {
				result = result + 100
			}
		case 'D':
			result = result + 500
		case 'M':
			result = result + 1000
		}
	}
	return result
}

func isNegative(c rune, pos int, s string) bool {
	if pos+1 >= len(s) {
		return false
	}
	n := s[pos+1]
	switch c {
	case 'I':
		if n == 'V' || n == 'X' {
			return true
		}
	case 'X':
		if n == 'L' || n == 'C' {
			return true
		}
	case 'C':
		if n == 'D' || n == 'M' {
			return true
		}

	default:
		return false
	}
	return false
}
