package main

func getSum(a int, b int) int {
	dir := -1
	if a < 0 {
		dir = 1
	}
	ans := b
	for i := a; i != 0; i = i + dir {
		if dir == -1 {
			ans++
		} else {
			ans--
		}
	}
	return ans
}
