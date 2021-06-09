package main

/*tion for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	out := []int{root.Val}
	for _, c := range root.Children {
		if c != nil {
			out = append(out, preorder(c)...)
		}

	}
	return out
}

/*
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorderSpeedImprovement(root *Node) []int {
	if root == nil {
		return []int{}
	}
	out := []int{root.Val}
	for _, root := range root.Children {
		out = append(out, preorder(root)...)
	}
	return out
}
