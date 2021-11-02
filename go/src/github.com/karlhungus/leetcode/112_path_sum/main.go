

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return hasPathSum2(root, 0, targetSum)
}

func hasPathSum2(root *TreeNode, curSum int, targetSum int) bool {
	if root.Left == nil && root.Right == nil {
		return (curSum + root.Val) == targetSum
	}
	var l, r bool
	if root.Left != nil {
		l = hasPathSum2(root.Left, root.Val+curSum, targetSum)
	}
	if root.Right != nil {
		r = hasPathSum2(root.Right, root.Val+curSum, targetSum)
	}
	return l || r
}


