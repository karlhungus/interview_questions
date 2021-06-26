package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	nodes := [][]*Node{}
	levelOrderDepth(root, 0, &nodes)
	for _, nl := range nodes {
		for j, n := range nl {
			if len(nl) > j+1 {
				n.Next = nl[j+1]
			}
		}
	}
	return root
}

func levelOrderDepth(root *Node, depth int, nodes *[][]*Node) {
	if root == nil {
		return
	}
	if len(*nodes) < depth+1 {
		*nodes = append(*nodes, []*Node{})
	}
	(*nodes)[depth] = append((*nodes)[depth], root)
	levelOrderDepth(root.Left, depth+1, nodes)
	levelOrderDepth(root.Right, depth+1, nodes)
}
