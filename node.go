package main

import (
	"fmt"
	"math"
)

type Node struct {
	value         int
	left          *Node
	right         *Node
	balanceFactor int
}

func NewNode(v int) *Node {
	return &Node{
		value:         v,
		balanceFactor: 0,
	}
}

func (n *Node) Print() {
	fmt.Printf("%v{%v}: [", n.value, n.balanceFactor)
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

func (n *Node) Height() int {
	return int(math.Max(float64(n.childHeight(n.left)), float64(n.childHeight(n.right))))
}

func (n *Node) Insert(v int) {
	if n.value == v {
		return
	}

	defer n.updateBalanceFactor()

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

func (n *Node) childHeight(child *Node) int {
	if child == nil {
		return 0
	}

	return child.Height() + 1
}

func (n *Node) updateBalanceFactor() {
	n.balanceFactor = n.childHeight(n.right) - n.childHeight(n.left)
}

func NodesAreEqual(n1 *Node, n2 *Node) bool {
	if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
		return false
	}
	if n1 == nil && n2 == nil {
		return true
	}

	if n1.value != n2.value {
		return false
	}

	return NodesAreEqual(n1.left, n2.left) && NodesAreEqual(n1.right, n2.right)
}
