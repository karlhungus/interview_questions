package main

func sortColors(nums []int) {
	quicksort(nums, 0, len(nums)-1)
}

func quicksort(nums []int, lo, hi int) {
	if lo < hi {
		p := partition(nums, lo, hi)
		quicksort(nums, lo, p-1)
		quicksort(nums, p+1, hi)
	}
}

func partition(nums []int, lo, hi int) int {
	p := lo
	pivot := nums[p]
	start := lo
	end := hi
	for start < end {
		for start < len(nums) && nums[start] <= pivot {
			start++
		}
		for nums[end] > pivot {
			end--
		}
		if start < end {
			swap(nums, start, end)
		}
	}
	swap(nums, end, p)
	return end
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
