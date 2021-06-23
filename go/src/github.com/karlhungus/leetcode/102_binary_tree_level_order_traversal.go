package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	vals := [][]int{}
	levelOrderDepth(root, 0, &vals)
	return vals
}

func levelOrderDepth(root *TreeNode, depth int, vals *[][]int) {
	if root == nil {
		return
	}
	if len(*vals) < depth+1 {
		*vals = append(*vals, []int{})
	}
	(*vals)[depth] = append((*vals)[depth], root.Val)
	levelOrderDepth(root.Left, depth+1, vals)
	levelOrderDepth(root.Right, depth+1, vals)
}
