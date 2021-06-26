package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	//post order last item is root node, pop that
	// use that to split between left and right subtree in preorder
	//recurse for both halves
	if len(postorder) == 0 {
		return nil
	}
	root := TreeNode{Val: postorder[len(postorder)-1]}
	left, right, idx := splitOn(inorder, root.Val)
	root.Left = buildTree(left, postorder[:idx])
	root.Right = buildTree(right, postorder[idx:len(postorder)-1])
	return &root
}

func splitOn(preorder []int, value int) (left []int, right []int, idx int) {

	for i, v := range preorder {
		if v == value {
			return preorder[:i], preorder[i+1:], i
		}
	}
	// this should be an error but, this is fine for now
	return []int{}, []int{}, -1
}

/*
[9,3,15,20,7], postorder = [9,15,7,20,3]
root=3 from post order, pop

left subtree [9]
root [3]
right subtree [15,20,7]
new post order [9,15,7,20]

repeat left
[9] only

repeat right
root = [20]
left [15]
right [7]
*/
