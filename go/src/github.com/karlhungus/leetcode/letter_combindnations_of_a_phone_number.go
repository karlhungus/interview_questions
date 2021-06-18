package main

var phone = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	// base case no digits left
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return phone[digits[0]]
	}
	// recursive case peal off the first digit, prepend it to each return value from leterCombindinations
	var col []string
	for _, c := range phone[digits[0]] {
		postfixes := letterCombinations(digits[1:])
		for _, p := range postfixes {
			x := append([]byte(c), p...)
			col = append(col, string(x))
		}
	}
	return col
}
