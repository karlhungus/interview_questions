package main

func rotate(nums []int, k int) {
	visited := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		prev := nums[i]
		j := (i + k) % len(nums)
		for !visited[j] {

			nums[j], prev = prev, nums[j]

			visited[j] = true
			j = (j + k) % len(nums)
		}
	}
}
