package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	rs := []rune(s)
	rt := []rune(t)
	for _, c := range rs {
		var r bool
		for i, x := range rt {
			if x == c {
				rt = remove([]rune(rt), i)
				r = true
				break
			}
		}
		if !r {
			return false
		}
	}
	return len(rt) == 0
}

func remove(s []rune, i int) []rune {
	return append(s[0:i], s[i+1:]...)
}
