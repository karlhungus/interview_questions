package main

//This is lazy
func findDisappearedNumbers(nums []int) []int {
	var missing []int
	for i := 0; i < len(nums); i++ {
		s := i + 1
		var f bool
		for _, n := range nums {
			if s == n {
				f = true
				break
			}
		}
		if !f {
			missing = append(missing, s)
		}
	}
	return missing

}
