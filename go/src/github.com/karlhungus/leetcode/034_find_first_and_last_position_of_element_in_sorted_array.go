package main

func searchRange(nums []int, target int) []int {
	// first find left most, then find right most
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{binarySearchLeft(nums, target), binarySearchRight(nums, target)}
}

func binarySearchLeft(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start < end && start >= 0 && end < len(nums) {
		pivot := start + (end-start)/2
		if nums[pivot] < target {
			start = pivot + 1
		} else {
			end = pivot
		}

	}
	if nums[start] == target {
		return start
	}
	return -1
}

func binarySearchRight(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start < end && start >= 0 && end < len(nums) {
		pivot := end - (end-start)/2
		if nums[pivot] <= target {
			start = pivot
		} else {
			end = pivot - 1
		}

	}
	if nums[end] == target {
		return end
	}
	return -1
}
