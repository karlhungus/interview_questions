package main

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, t := range tokens {
		if !strings.Contains("*+/-", t) {
			x, _ := strconv.Atoi(t)
			stack = append(stack, x)
		} else {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, eval(a, b, t))
		}
	}
	return stack[0]
}

func eval(a int, b int, sym string) int {
	switch sym {
	case "+":
		return a + b
	case "*":
		return a * b
	case "/":
		return a / b
	case "-":
		return a - b
	}
	// this should be error, same with / 0, but whatever
	return 0
}
