package main

func longestPalindrome(s string) string {
	sp := s
	if len(s)%2 == 0 {
		//it's even we need to make it odd
		sp = makeOdd(s)
	}
	palrad := make(map[int]int, len(sp))
	//storage := map[int]string
	//var cap []rune
	for center := 0; center < len(sp); center++ {
		var rad int
		for center-(rad+1) >= 0 && center+(rad+1) < len(sp) && sp[center-(rad+1)] == sp[center+(rad+1)] {
			rad++
		}
		palrad[center] = rad
	}
	var maxpos, max int
	for k, v := range palrad {
		if v > max {
			maxpos = k
			max = v
		}
	}
	fmt.Printf("maxpos:%d, max:%d, s:%s", maxpos, max, s)
	return s[maxpos-max : maxpos+max+1]
}

func makeOdd(s string) string {
	col := []byte("|")
	for i := 0; i < len(s); i++ {
		col = append(col, s[i])
		col = append(col, '|')
	}
	return string(col)
}
