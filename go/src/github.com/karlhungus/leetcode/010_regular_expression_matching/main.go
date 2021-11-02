package main

import "fmt"

func main() {
	validate("aaaaa", "aaaaa", true)
	validate("aaaab", "aaaab", true)
	validate("aaaabc", "aaaab", false)
	validate("aaaa", "a*", true)
	validate("aaaabbb", "a*bbb", true)
}

func validate(s string, p string, want bool) {
	fmt.Printf("s: %s, p: %s, want: %t, got: %t\n", s, p, want, IsMatch(s, p))
}

func IsMatch(s string, p string) bool {
	// base case
	if len(p) == 0 {
		return len(s) == 0
	}
	first := len(s) != 0 && (s[0] == p[0] || p[0] == '.')
	// if the next one is a * we have to consume all matches, or move on because it's 0 or more matches
	if len(p) > 1 && p[1] == '*' {
		return (first && IsMatch(s[1:], p)) || (IsMatch(s, p[2:]))
	}
	return first && IsMatch(s[1:], p[1:])
}
