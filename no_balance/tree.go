package no_balance

import "fmt"

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

func (t *Tree) Print() {
	t.root.Print()
	fmt.Print("\n")
}

func BuildTree(values []int) *Tree {
	tree := &Tree{}
	for _, v := range values {
		tree.Insert(v)
		tree.Print()
	}

	return tree
}
