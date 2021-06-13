package main

func peakIndexInMountainArray(arr []int) int {
	for i, n := range arr {
		if i < len(arr) {
			if arr[i+1] < n {
				return i
			}
		}
	}
	return len(arr) - 1
}
