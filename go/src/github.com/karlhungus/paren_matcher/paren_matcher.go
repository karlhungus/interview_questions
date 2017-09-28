package main

import "os"
import "fmt"
import "regexp"
import "github.com/golang-collections/collections/stack"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expect one argument i.e.: ./brace_matcher ((()()))")
		return
	}

	validInput := regexp.MustCompile(`[\(\)]*`)
	if !validInput.MatchString(os.Args[1]) {
		fmt.Printf("%q Argument must contain only parens '(' and ')'.\n", os.Args[1])
		return
	}

	balenced := parenMatcher(os.Args[1])
	fmt.Printf("Balenced: %t\n", balenced)
}

/**
* Question
* Count matching braces open braces must come before closing braces
* eg input="()()((()))" output = 5, ")(" = 0
*
* I think the question here was really "do the parens balence out - ('s must come first"
 */
func parenMatcher(input string) bool {
	stack := stack.New()
	for _, char := range input {
		if char == '(' {
			stack.Push(char)
		} else if char == ')' {
			if stack.Len() == 0 || stack.Pop() != '(' {
				return false
			}
		}
	}
	return stack.Len() == 0
}
