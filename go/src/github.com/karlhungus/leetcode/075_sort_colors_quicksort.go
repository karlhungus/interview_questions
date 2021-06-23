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
	pivot := nums[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			swap(nums, i, j)
			i++
		}

	}
	swap(nums, i, hi)
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
