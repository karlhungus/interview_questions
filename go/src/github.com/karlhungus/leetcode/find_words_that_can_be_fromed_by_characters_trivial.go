package main

func countCharacters(words []string, chars string) int {
	var cnt int
	for _, w := range words {
		if match(w, chars) {
			cnt += len(w)
		}
	}
	return cnt
}

func match(w string, chars string) bool {
	x := []byte(chars)
	for i := 0; i < len(w); i++ {
		idx := strings.IndexByte(string(x), w[i])
		if idx < 0 {
			return false
		}
		x = append(x[:idx], x[idx+1:]...)
	}
	return true
}
