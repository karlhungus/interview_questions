package main

func largestNumber(nums []int) string {
	snums := []string{}
	for _, n := range nums {
		snums = append(snums, strconv.Itoa(n))
	}
	sort.Slice(snums, func(i, j int) bool {
		// concat them, then compare both representations
		a := snums[i] + snums[j]
		b := snums[j] + snums[i]
		return a > b
	})

	res := strings.Join(snums, "")
	// replace all zeros with single zero, maybe string->int->string would be better
	zreplacer := regexp.MustCompile("^0+$")
	res = zreplacer.ReplaceAllString(res, "0")
	return res
}
