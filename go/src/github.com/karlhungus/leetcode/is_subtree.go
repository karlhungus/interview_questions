package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil && subRoot != nil {
		return false
	}
	if root != nil && subRoot == nil {
		return false
	}
	if root.Val == subRoot.Val && equal(root, subRoot) {
		return true
	}
	if isSubtree(root.Left, subRoot) {
		return true
	}

	return isSubtree(root.Right, subRoot)
}

func equal(r *TreeNode, s *TreeNode) bool {
	if r == nil && s == nil {
		return true
	}
	if r == nil && s != nil {
		return false
	}
	if r != nil && s == nil {
		return false
	}
	if r.Val != s.Val {
		return false
	}
	return equal(r.Left, s.Left) && equal(r.Right, s.Right)
}
