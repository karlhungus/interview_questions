package main

import "testing"

func TestTreePrinter(t *testing.T) {
	/*
	*  5
	* / \
	*4   6
	*
	 */

	four := Node{Value: 4}
	six := Node{Value: 6}
	five := Node{Value: 5, Left: &four, Right: &six}
	tree := Tree{Root: &five}
	output := printTree(&tree)
	if output != "5\n6,4\n" {
		t.Error("Expected:\n 5\n6,4\n, for tree got ", output)
	}
}
