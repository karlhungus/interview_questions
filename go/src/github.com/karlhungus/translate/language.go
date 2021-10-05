package translate

import (
	"sort"
	"strings"
)

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
			} else if r == "*" && !ok {
				set[s] = struct{}{}
				results = append(results, s)
			}
		}
	}
	return results
}

type PriorityPair struct {
	language string
	priority float32
}

type PriorityPairs []PriorityPair

func (p PriorityPairs) Len() int           { return len(p) }
func (p PriorityPairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PriorityPairs) Less(i, j int) bool { return p[i].priority < p[j].priority }

func languagesPriority(requested PriorityPairs, supported []string) []string {
	set := make(map[string]struct{})
	sort.Sort(requested)
	results := []string{}
	for _, r := range requested {
		for _, s := range supported {
			_, ok := set[s]
			if s == r.language && !ok {
				set[s] = struct{}{}
				results = append(results, s)
			} else if strings.HasPrefix(s, r.language) && !ok {
				set[s] = struct{}{}
				results = append(results, s)
			} else if r.language == "*" && !ok {
				set[s] = struct{}{}
				results = append(results, s)
			}
		}
	}
	return results
}
