package main

func maxArea(height []int) int {
	m1 := 0
	m2 := 0
	ptrL := 0
	ptrR := len(height) - 1

	for ptrL < ptrR && ptrR > ptrL {
		hR := height[ptrR]
		hL := height[ptrL]
		h := min(hR, hL)
		l := ptrR - ptrL
		if h*l > m1*m2 {
			m1 = h
			m2 = l
		}
		if hR < hL {
			ptrR--
		} else {
			ptrL++
		}
	}
	return m1 * m2

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
