package main

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	} else if len(s) == 2 && s[0] != s[1] {
		return string(s[0])

	}
	sp := pad(s)

	palrads := make(map[int]int, len(sp))
	for center, _ := range sp {
		palrads[center] = palrad(center, sp)
	}
	var maxpos, max int
	for k, v := range palrads {
		if v > max {
			maxpos = k
			max = v
		}
	}
	return unpad(sp[maxpos-max : maxpos+max+1])
}

func pad(s string) string {
	col := []byte("|")
	for i := 0; i < len(s); i++ {
		col = append(col, s[i])
		col = append(col, '|')
	}
	return string(col)
}

func unpad(s string) string {
	col := []byte("")
	for i := 0; i < len(s); i++ {
		if s[i] != '|' {
			col = append(col, s[i])
		}
	}
	return string(col)
}

func palrad(center int, sp string) int {
	var rad int
	for center-(rad+1) >= 0 &&
		center+(rad+1) < len(sp) &&
		sp[center-(rad+1)] == sp[center+(rad+1)] {
		rad++
	}
	return rad
}
