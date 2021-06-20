package main

func permute(nums []int) [][]int {
	//take each number, make it the first and permute the rest

	//base case
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	col := [][]int{}
	for i := 0; i < len(nums); i++ {
		pre := make([]int, i)
		post := make([]int, len(nums)-(i+1))
		copy(pre, nums[:i])
		copy(post, nums[i+1:])
		perms := permute(append(pre, post...))
		for _, t := range perms {
			x := append(t, nums[i])
			col = append(col, x)
		}
	}
	return col

}
