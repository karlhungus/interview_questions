package main

func majorityElement(nums []int) int {
	counter := make(map[int]int)
	var cur, cnt int
	for _, i := range nums {
		n, ok := counter[i]
		if !ok {
			counter[i] = 1
		} else {
			counter[i] = n + 1
		}
		if cnt == 0 || cnt < counter[i] {
			cur = i
			cnt = counter[i]
		}
	}
	return cur
}
