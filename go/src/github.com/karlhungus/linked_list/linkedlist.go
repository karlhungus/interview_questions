package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Value string
	Next  *Node
}

type LinkedList struct {
	Root *Node
}

func main() {
	ll := LinkedList{}
	for i := 0; i < 10; i++ {
		ll.Add(&Node{
			Value: fmt.Sprintf("node %d", i),
		})
	}
	fmt.Printf("%s\n", ll)
}

func (ll LinkedList) String() string {
	var b strings.Builder
	b.WriteString("LinkedList [")
	cur := ll.Root
	if cur == nil {
		b.WriteString("]")
		return b.String()

	}
	b.WriteString(fmt.Sprintf("Root: %s", cur.Value))

	for {
		cur = cur.Next
		if cur == nil {
			b.WriteString("]")
			return b.String()
		} else {
			b.WriteString(fmt.Sprintf("--> %s", cur.Value))
		}

	}
}
func (ll *LinkedList) Add(n *Node) {

	if ll.Root == nil {
		ll.Root = n
		return
	}
	cur := ll.Root
	for {
		if cur.Next == nil {
			cur.Next = n
			return
		}
		cur = cur.Next
	}
}
