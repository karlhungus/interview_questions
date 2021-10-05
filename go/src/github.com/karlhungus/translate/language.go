package translate

func languages(requested []string, supported []string) []string {
	results := []string{}
	for _, s := range supported {
		for _, r := range requested {
			if s == r {
				results = append(results, r)
			}
		}

	}
	return results
}
