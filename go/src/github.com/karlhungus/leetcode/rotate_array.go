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

func rotate2(nums []int, k int) {
	visited := 0
	// visit each position
	for i := 0; visited < len(nums); i++ {
		//store the first move
		prev := nums[i]
		// find where it should move
		j := (i + k) % len(nums)
		// while we haven't moved the one we are on continue
		for i != j {
			//do the swap, storing what was moved to
			nums[j], prev = prev, nums[j]
			//update what we've swapped
			visited++
			//calculate the next position
			j = (j + k) % len(nums)
		}
		// move the last umoved item back to i
		nums[i] = prev
		// update visited count
		visited++
	}
}
