package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1

	for start <= end {
		pivot := (end + start) / 2

		if nums[pivot] == target {
			return pivot
		}
		if nums[start] <= nums[pivot] {
			// we know that the rotation begins on the right, and can reason about the left side
			if target >= nums[start] && target < nums[pivot] {
				end = pivot - 1
			} else {
				start = pivot + 1
			}

		} else {
			// we know that the rotation begins on the left, and can reason about the right side
			if target <= nums[end] && target > nums[pivot] {
				start = pivot + 1
			} else {
				end = pivot - 1
			}

		}
	}
	return -1
}
