package main

func fractionToDecimal(numerator int, denominator int) string {
	// determine if it's negative
	// divide take the bit before the decimal
	// modulus the bit after to get the decimal part
	if numerator == 0 {
		return "0"
	}
	var res string

	// negative
	if (numerator < 0) != (denominator < 0) {
		res = "-"
	}

	// elimnate negatives
	if numerator < 0 {
		numerator = numerator * -1
	}
	if denominator < 0 {
		denominator = denominator * -1
	}
	n := numerator
	d := denominator

	// integer part
	res = fmt.Sprintf("%s%d", res, n/d)
	n %= d
	if n == 0 {
		return res
	}
	res = fmt.Sprintf("%s.", res)

	// Map will store remainders to position
	m := map[int]int{}
	// decimal part
	for n != 0 {
		n *= 10
		res = fmt.Sprintf("%s%d", res, n/d)
		n %= d

		if p, ok := m[n]; ok {
			// plop a '(' at p
			// append an ')'
			return fmt.Sprintf("%s(%s)", res[:p], res[p:])
		} else {
			m[n] = len(res)
		}
	}
	return res
}
