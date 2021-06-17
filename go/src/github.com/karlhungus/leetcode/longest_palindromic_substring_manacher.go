package main

//https://tarokuriyama.com/projects/palindrome2.php

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	} else if len(s) == 2 && s[0] != s[1] {
		return string(s[0])

	}
	sp := pad(s)

	palLen := make(map[int]int, len(sp))
	for center := 0; center < len(sp); center++ {
		if _, ok := palLen[center]; ok {
			continue
		}
		rad := palrad(center, sp)
		palLen[center] = rad
		right := min(center+rad, len(sp))
		for mirror := 1; center+mirror < right; mirror++ {
			// fmt.Println(mirror)
			pLeft := center - mirror
			pRight := center + mirror
			distToEdge := right - pRight

			if palLen[pLeft] < distToEdge {
				palLen[pRight] = palLen[pLeft]
			} else if palLen[pLeft] > distToEdge {
				palLen[pRight] = distToEdge
			} else if palLen[pLeft] == distToEdge {
				palLen[pRight] = palrad(pRight, sp)
			}
		}

		//calc as length plen = palrad*2 +1
		//for each center calc
		// pRight (next center until last rad)
		// distToEdge (last palrad)
		// pLeft = mirror of pRight
		// Cases:
		// 1 palLen[pLeft] < distToEdge : palLen[pRight]=palLen[pLeft] next
		// 2 palLen[pLeft] > distToEdge : palLen[pRight] = distToEdget next
		// 3 palLen[pleft] == distToEdget : calc expansion
	}
	var maxpos, max int
	for k, v := range palLen {
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

func pallength(center int, sp string) int {
	var rad int
	for center-(rad+1) >= 0 &&
		center+(rad+1) < len(sp) &&
		sp[center-(rad+1)] == sp[center+(rad+1)] {
		rad++
	}
	if rad == 0 {
		return 0
	}
	return (rad * 2) + 1
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
