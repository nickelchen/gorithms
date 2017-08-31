package tree

import "fmt"

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	Value  rune
}

func (node *Node) print(prefix string, right bool) {
	var lead string
	var space string

	if right {
		lead = "└── "
		space = "    "
	} else {
		lead = "├── "
		space = "│   "
	}

	fmt.Println(prefix + lead + string(node.Value))

	if node.left != nil {
		node.left.print(prefix+space, false)
	}
	if node.right != nil {
		node.right.print(prefix+space, true)
	}
}

func (node *Node) Print() {
	node.print("", true)
}

func (node *Node) AttachL(child *Node) {
	if node.left != nil {
		node.DetachL(node.left)
	}

	node.left = child
	child.parent = node
}

func (node *Node) AttachR(child *Node) {
	if node.right != nil {
		node.DetachR(node.right)
	}

	node.right = child
	child.parent = node
}

func (node *Node) DetachL(child *Node) {
	if node != nil {
		node.left = nil
	}
	child.parent = nil
}

func (node *Node) DetachR(child *Node) {
	if node != nil {
		node.right = nil
	}
	child.parent = nil
}
