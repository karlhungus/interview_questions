package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 0 {
		return n <= 0
	}
	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		return n <= 1
	}
	if len(flowerbed) == 2 && flowerbed[0] == 0 && flowerbed[1] == 0 {
		return n < 2
	}

	var total, count int
	if len(flowerbed) >= 2 && flowerbed[0] == 0 && flowerbed[1] == 0 {
		total++
	}
	for i := 1; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			count++
		} else {
			total = total + int((count-1)/2)
			count = 0
		}
	}

	if count != 0 {
		total = total + int(count/2)
	}

	return total >= n

}
