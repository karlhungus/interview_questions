package translate

import "strings"

func languages(requested []string, supported []string) []string {
	set := make(map[string]struct{})
	results := []string{}
	for _, r := range requested {
		for _, s := range supported {
			_, ok := set[s]
			if s == r && !ok {
				set[s] = struct{}{}
				results = append(results, s)
			} else if strings.HasPrefix(s, r) && !ok {
				set[s] = struct{}{}
				results = append(results, s)
			}
		}
	}
	return results
}
