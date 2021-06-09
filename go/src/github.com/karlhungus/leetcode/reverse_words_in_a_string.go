package main

func reverseWords(s string) string {
	var result string
	split := strings.Split(s, " ")
	for i, w := range split {

		var r string
		for _, c := range w {
			r = string(c) + r
		}

		if i != 0 {
			result = result + " "
		}
		result = result + r
	}
	return result
}
