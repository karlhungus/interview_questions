package main

func checkPossibility(nums []int) bool {
	var modified bool
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if modified {
				return false
			} else {
				//if first, or can decrese left decrese
				if i == 0 || nums[i-1] <= nums[i+1] {
					nums[i] = nums[i+1]
				} else { // increase right
					nums[i+1] = nums[i]
				}

				modified = true
			}
		}
	}

	return true
}
