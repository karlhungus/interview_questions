package main

func isValid(s string) bool {
	v := map[rune]rune{')': '(', '}': '{', ']': '['}
	var stack []rune
	for _, r := range s {
		if len(stack) > 0 && stack[len(stack)-1] == v[r] {
			stack = stack[:len(stack)-1]
		} else if _, ok := v[r]; ok {
			return false
		} else {
			stack = append(stack, r)
		}
	}
	return len(stack) == 0
}
