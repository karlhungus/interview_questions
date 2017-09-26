package main

import "os"
import "fmt"
import "strconv"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Expect two argument i.e.: ./run_finder 3 1111000111")
		return
	}

	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q First param must be an int.\n", os.Args[1])
		return
	}

	total := lengthFinder(length, os.Args[2])
	fmt.Printf("Total: %d\n", total)
}

// In a string like "11111000" find runs of length x 3 would = 6 for last string
func lengthFinder(length int, input string) int {
	count := 0
	total := 0
	var currentChar rune
	for _, char := range input {
		if char != currentChar {
			currentChar = char
			count = 1
		} else {
			count++
			if count >= length {
				total++
			}
		}
	}
	return total
}
