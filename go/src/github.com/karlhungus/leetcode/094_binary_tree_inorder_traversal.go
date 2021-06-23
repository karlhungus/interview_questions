package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	vals := []int{}
	vals = append(vals, inorderTraversal(root.Left)...)
	vals = append(vals, root.Val)
	vals = append(vals, inorderTraversal(root.Right)...)
	return vals
}
