package main

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mappinga := make(map[byte]byte)
	mappingb := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		c := s[i]
		d := t[i]
		x, ok := mappinga[c]
		if !ok {
			mappinga[c] = d
			x = d
		}
		y, ok := mappingb[d]
		if !ok {
			mappingb[d] = c
			y = c
		}
		if c != y || d != x {
			return false
		}
	}
	return true
}
