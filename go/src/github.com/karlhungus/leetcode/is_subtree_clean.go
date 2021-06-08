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
	if root == nil {
		return false
	}
	return equal(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)

}

func equal(r *TreeNode, s *TreeNode) bool {
	if r == nil || s == nil {
		return r == s
	}

	return r.Val == s.Val && equal(r.Left, s.Left) && equal(r.Right, s.Right)
}
