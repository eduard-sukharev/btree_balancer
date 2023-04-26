package no_balance

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(v int) *Node {
	return &Node{
		value: v,
	}
}

func (n *Node) Print() {
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

func (n *Node) Insert(v int) {
	if n.value == v {
		return
	}

	// insert left
	if v < n.value {
		if n.left == nil {
			n.left = NewNode(v)
			return
		}
		n.left.Insert(v)
	}
	// insert right
	if v > n.value {
		if n.right == nil {
			n.right = NewNode(v)
			return
		}
		n.right.Insert(v)
	}
}
