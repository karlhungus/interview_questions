package main

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	cnt := 0
	max := int(math.Sqrt(float64(n)))
	notPrime := make([]bool, n)

	for i := 2; i <= max; i++ {
		if notPrime[i] == false {
			for j := i * i; j < n; j += i {
				notPrime[j] = true
			}
		}

	}
	for i := 2; i < n; i++ {
		if notPrime[i] == false {
			cnt++
		}
	}

	return cnt
}
