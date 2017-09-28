package main

type Tree struct {
	Root *Node
}

type Node struct {
	Value       int
	Left, Right *Node
}

func (this *Tree) Clear() {
	this.Root = nil
}
