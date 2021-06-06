package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	ptr1 := m - 1
	ptr2 := n - 1
	pos := m + n - 1

	for pos >= 0 {
		if ptr2 < 0 {
			return
		}
		if ptr1 < 0 || nums1[ptr1] < nums2[ptr2] {
			nums1[pos] = nums2[ptr2]
			ptr2--

		} else {
			nums1[pos] = nums1[ptr1]
			ptr1--
		}
		pos--
	}

}
