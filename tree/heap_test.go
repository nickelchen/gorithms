package tree

import (
	"fmt"
	"testing"
)

func DemoHeap() *Node {
	/*
	   └── T
	       ├── S
	       │   ├── P
	       │   │   ├── E
	       │   │   └── I
	       │   └── N
	       │       ├── H
	       │       └── G
	       └── R
	           ├── O
	           └── A
	*/
	root := &Node{value: 'T'}

	// child level 1
	c11 := &Node{value: 'S'}
	c12 := &Node{value: 'R'}

	// child level 2
	c21 := &Node{value: 'P'}
	c22 := &Node{value: 'N'}
	c23 := &Node{value: 'O'}
	c24 := &Node{value: 'A'}

	// child level 3
	c31 := &Node{value: 'E'}
	c32 := &Node{value: 'I'}
	c33 := &Node{value: 'H'}
	c34 := &Node{value: 'G'}

	AttachL(root, c11)
	AttachR(root, c12)

	AttachL(c11, c21)
	AttachR(c11, c22)

	AttachL(c12, c23)
	AttachR(c12, c24)

	AttachL(c21, c31)
	AttachR(c21, c32)

	AttachL(c22, c33)
	AttachR(c22, c34)

	return root
}

func TestHeap(_ *testing.T) {
	root := DemoHeap()
	root.Print()

	node := &Node{value: 'Z'}
	Insert(root, node)

	fmt.Printf("after insert Z:\n")
	root.Print()

	Remove(root)

	fmt.Printf("after remove top:\n")
	root.Print()

}
