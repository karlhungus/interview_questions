package main

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := [][]int{}
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if cur[1] >= intervals[i][0] {
			cur[1] = max(cur[1], intervals[i][1])
		} else {
			res = append(res, cur)
			cur = intervals[i]
		}
	}
	res = append(res, cur)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
