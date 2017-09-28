package main

//import "os"
import "fmt"

type NodeExecuter func(node Node, depth int)

func main() {
	tree := Tree{&Node{Value: 5, Left: &Node{Value: 4}, Right: &Node{Value: 6}}}
	output := printTree(&tree)
	fmt.Printf("Tree: \n%t\n", output)
}

func printTree(tree *Tree) string {
	printArray := make([][]int, 0)
	breadthTraversal(tree, func(node Node, depth int) {
		if len(printArray) <= depth {
			printArray = append(printArray, make([]int, 0))
		}
		printArray[depth] = append(printArray[depth], node.Value)
	})

	output := ""
	for _, array := range printArray {
		for i, value := range array {
			if i >= 0 && i < len(array)-1 {
				output += fmt.Sprintf("%v,", value)
			} else {
				output += fmt.Sprintf("%v", value)
			}
		}
		output += fmt.Sprintf("\n")
	}
	return output
}

func breadthTraversal(tree *Tree, executor NodeExecuter) string {
	output := ""
	queue := make([]*Node, 0)
	set := make(map[*Node]int)

	set[tree.Root] = 0
	queue = append(queue, tree.Root)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		executor(*current, set[current])

		if current.Left != nil {
			if _, present := set[current.Left]; !present {
				set[current.Left] = set[current] + 1
				queue = append(queue, current.Left)
			}
		}
		if current.Right != nil {
			if _, present := set[current.Right]; !present {
				set[current.Right] = set[current] + 1
				queue = append(queue, current.Right)
			}
		}
	}
	return output
}
