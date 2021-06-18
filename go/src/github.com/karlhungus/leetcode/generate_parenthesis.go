package main

func generateParenthesis(n int) []string {
	combinations := &[]string{}
	return *backtrack(combinations, "", 0, 0, n)
}

func backtrack(combinations *[]string, str string, open int, close int, n int) *[]string {
	if len(str) == 2*n {
		*combinations = append(*combinations, str)
		return combinations
	}

	if open < n {
		combinations = backtrack(combinations, str+"(", open+1, close, n)
	}

	if close < open {
		combinations = backtrack(combinations, str+")", open, close+1, n)
	}
	return combinations
}
