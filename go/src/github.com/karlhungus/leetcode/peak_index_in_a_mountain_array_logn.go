package main

func peakIndexInMountainArray(arr []int) int {
	_, y := pmia(arr, 0)
	return y
}

func pmia(arr []int, p int) (int, int) {
	if len(arr) == 1 {
		return arr[0], p
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0], p
		} else {
			return arr[1], p + 1
		}
	}
	split := len(arr) / 2
	l := arr[:split]
	r := arr[split:]
	w, x := pmia(l, 0+p)
	y, z := pmia(r, split+p)
	if w < y {
		return y, z
	}
	return w, x
}
