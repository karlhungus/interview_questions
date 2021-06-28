package main

func reverseWords(s string) string {
	stack := []string{}
	var i int
	for i < len(s) {
		for i < len(s) && s[i] == ' ' {
			i++
		}
		end := i
		for end < len(s) && s[end] != ' ' {
			end++
		}
		if end != i {
			stack = append(stack, s[i:end])
			i = end
		}
	}
	var rs string
	for i := len(stack) - 1; i >= 0; i-- {
		rs = rs + stack[i]
		if i != 0 {
			rs = rs + " "
		}
	}
	return rs
}
