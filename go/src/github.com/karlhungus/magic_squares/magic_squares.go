package main

import "os"
import "fmt"
import "strconv"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expect one odd argument i.e.: magic_squares 3")
		return
	}
	size, err := strconv.Atoi(os.Args[1])
	if size%2 != 1 || err != nil {
		fmt.Println("Expect one odd argument i.e.: magic_squares 3 got", size)
		return
	}

	square := magicSquare(size)
	fmt.Printf("Magic square:\n")
	for _, value := range square {
		fmt.Printf("%v\n", value)
	}
}

/**
*
* Odd algorithm
* 1. start in central position of first row with 1
* 2. from your position go up and right, wrapping when going past an edge
* 3. if a square is already filled drop down one from your existing position
* 4. stop after filling n^2 positions
 */
func magicSquare(size int) [][]int {
	square := make([][]int, size)
	for i := range square {
		square[i] = make([]int, size)
	}

	x := int(size / 2)
	y := 0
	newX := 0
	newY := 0
	square[y][x] = 1
	for i := 2; i <= size*size; i++ {
		newX = (x + 1) % size
		if y-1 < 0 {
			newY = size - 1
		} else {
			newY = y - 1
		}
		if square[newY][newX] != 0 {
			newX = x
			newY = y + 1%size
		}
		x = newX
		y = newY
		square[y][x] = i
	}
	return square
}
