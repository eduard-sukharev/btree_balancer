package main

import "fmt"

type Node struct {
	value      int
	left       *Node
	right      *Node
	heightDiff int8
}

func NewNode(v int) *Node {
	return &Node{
		value:      v,
		heightDiff: 0,
	}
}

func (n Node) Print() {
	fmt.Printf("%v: [", n.value)
	if n.left != nil {
		n.left.Print()
	} else {
		fmt.Print("<nil>")
	}
	fmt.Print(", ")
	if n.right != nil {
		n.right.Print()
	} else {
		fmt.Print("<nil>")
	}
	fmt.Print("]")
}

func (t *Node) Insert(v int) {
	if t.value == v {
		return
	}

	// left
	if v < t.value {
		if t.left == nil {
			t.left = NewNode(v)
			return
		}
		t.left.Insert(v)
		return
	}
	// right
	if v > t.value {
		if t.right == nil {
			t.right = NewNode(v)
			return
		}
		t.right.Insert(v)
		return
	}
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(v int) {
	if t.root == nil {
		t.root = &Node{value: v}
		return
	}
	t.root.Insert(v)
}

func (t Tree) Print() {
	t.root.Print()
	fmt.Print("\n")
}

func generateTreeValues(upTo int) []int {
	rv := make([]int, 0, upTo-1)
	for i := 1; i <= upTo; i++ {
		rv = append(rv, i)
	}

	return rv
}

func main() {
	tree := Tree{}
	for _, v := range generateTreeValues(6) {
		tree.Insert(v)
	}

	tree.Print()
}
