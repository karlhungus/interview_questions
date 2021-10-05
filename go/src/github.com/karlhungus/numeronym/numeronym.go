package numeronym

import (
	"fmt"
	"strconv"
	"strings"
)

func numeronymUrl(url string) string {
	split1 := strings.Split(url, "/")
	col1 := make([]string, len(split1))
	for i, s1 := range split1 {
		fmt.Printf("%d, s1: %s\n", i, s1)
		split2 := strings.Split(s1, ".")
		col2 := make([]string, len(split2))
		for j, s2 := range split2 {
			fmt.Printf("%d, s1: %s\n", j, s2)
			col2[j] = numeronym(s2)
		}
		col1[i] = strings.Join(col2, ".")
	}
	return strings.Join(col1, "/")
}

func numeronym(s string) string {
	switch len(s) {
	case 0:
		return ""
	case 1:
		return s
	case 2:
		return s
	default:
		x := strconv.Itoa(len(s) - 2)
		return strings.Join([]string{string(s[0]), x, string(s[len(s)-1])}, "")
	}
}
