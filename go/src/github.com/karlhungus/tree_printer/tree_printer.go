package main

//import "os"
import "fmt"

func main() {
	/*	if len(os.Args) < 2 {
			fmt.Println("Expect one argument i.e.: ./brace_matcher ((()()))")
			return
		}

		validInput := regexp.MustCompile(`[\(\)]*`)
		if !validInput.MatchString(os.Args[1]) {
			fmt.Printf("%q Argument must contain only parens '(' and ')'.\n", os.Args[1])
			return
		}
	*/
	tree := Tree{&Node{Value: 5, Left: &Node{Value: 4}, Right: &Node{Value: 6}}}
	output := printTree(&tree)
	fmt.Printf("Tree: \n%t\n", output)
}

func printTree(tree *Tree) string {
	output := ""
	queue := make([]Node, 0)
	set := make(map[int]bool)

	set[tree.Root.Value] = true
	queue = append(queue, *tree.Root)
	even := true
	counter := 4

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		output += fmt.Sprintf("%d,", current.Value)
		fmt.Printf("length: %d, value: %v, even: %t\n", len(queue), current.Value, even)
		counter--
		if counter < 0 {
			return output
		}

		if even {
			fmt.Printf("even\n")
			if current.Left != nil {

				fmt.Printf("current left not null!\n")
				if _, present := set[current.Left.Value]; !present {
					fmt.Printf("LEFT!\n")
					set[current.Left.Value] = true
					queue = append(queue, *current.Left)
				}
			}
			if current.Right != nil {
				if _, present := set[current.Right.Value]; !present {
					fmt.Printf("RIGHT!\n")
					set[current.Right.Value] = true
					queue = append(queue, *current.Right)
				}
			}
		} else {
			fmt.Printf("ODD!\n")

			if current.Right != nil {
				if _, present := set[current.Right.Value]; !present {
					fmt.Printf("RIGHT!\n")
					set[current.Right.Value] = true
					queue = append(queue, *current.Right)
				}
			}
			if current.Left != nil {
				if _, present := set[current.Left.Value]; !present {
					fmt.Printf("LEFT!\n")
					set[current.Left.Value] = true
					queue = append(queue, *current.Left)
				}
			}
		}
		even = !even
	}
	return output
}

//
// queue := make([]int, 0)
//Push
//queue := append(queue, 1)
// Top (just get next element, don't remove it)
// x = queue[0]
// // Discard top element
// queue = queue[1:]
// // Is empty ?
// if len(queue) == 0 {
//     fmt.Println("Queue is empty !")
//     }
//
//
//
//     to imitate a set ints in Go, we define a map:
//
//     set := make(map[int]bool)
//     To add something is as easy as:
//
//     i := valueToAdd()
//     set[i] = true
//     Deleting something is just
//
//     delete(set, i)
