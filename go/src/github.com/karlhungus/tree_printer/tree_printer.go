package main

//import "os"
import "fmt"

func main() {
	tree := Tree{&Node{Value: 5, Left: &Node{Value: 4}, Right: &Node{Value: 6}}}
	output := printTree(&tree)
	fmt.Printf("Tree: \n%t\n", output)
}

func printTree(tree *Tree) string {
	return breadthTraversal(tree)
}

func breadthTraversal(tree *Tree) string {
	output := ""
	queue := make([]*Node, 0)
	set := make(map[*Node]bool)

	set[tree.Root] = true
	queue = append(queue, tree.Root)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if output != "" {
			output += fmt.Sprintf(",%d", current.Value)
		} else {
			output += fmt.Sprintf("%d", current.Value)
		}
		fmt.Printf("length: %d, value: %v", len(queue), current.Value)

		if current.Left != nil {
			if _, present := set[current.Left]; !present {
				set[current.Left] = true
				queue = append(queue, current.Left)
			}
		}
		if current.Right != nil {
			if _, present := set[current.Right]; !present {
				set[current.Right] = true
				queue = append(queue, current.Right)
			}
		}
	}
	return output
}
