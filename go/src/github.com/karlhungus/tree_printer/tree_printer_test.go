package main

import "testing"

func TestTreePrinter(t *testing.T) {
	/*
	*      5
	*    /  \
	*   2    8
	*  / \  / \
	* 1  3 6   9
	 */

	two := Node{Value: 2, Left: &Node{Value: 1}, Right: &Node{Value: 3}}
	eight := Node{Value: 8, Left: &Node{Value: 6}, Right: &Node{Value: 9}}
	five := Node{Value: 5, Left: &two, Right: &eight}
	tree := Tree{Root: &five}
	output := printTree(&tree)
	expected := "5,2,8,1,3,6,9"
	if output != expected {
		t.Error("Expected:\n %v for tree got %v", expected, output)
	}
}
