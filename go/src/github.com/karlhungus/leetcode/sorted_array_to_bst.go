package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	var median = int(len(nums) / 2)
	root := TreeNode{
		Val: nums[median],
	}
	root.Right = sortedArrayToBST(nums[median+1:])
	root.Left = sortedArrayToBST(nums[:median])

	return &root
}
