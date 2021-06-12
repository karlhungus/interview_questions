package main

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := 1
	curmax := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			curmax++
		} else {
			curmax = 1
		}
		if curmax > max {
			max = curmax
		}
	}
	return max

}
