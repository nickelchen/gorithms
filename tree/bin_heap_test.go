package tree

import (
	"fmt"
	"testing"
)

func DemoBinaryHeap() *BinaryHeap {
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

	root.AttachL(c11)
	root.AttachR(c12)

	c11.AttachL(c21)
	c11.AttachR(c22)

	c12.AttachL(c23)
	c12.AttachR(c24)

	c21.AttachL(c31)
	c21.AttachR(c32)

	c22.AttachL(c33)
	c22.AttachR(c34)

	return NewBinaryHeap(root)
}

func TestBinaryHeap(_ *testing.T) {
	heap := DemoBinaryHeap()
	heap.Print()

	node := &Node{value: 'Z'}
	heap.Insert(node)

	fmt.Printf("after insert Z:\n")
	heap.Print()

	heap.RemoveMax()

	fmt.Printf("after remove top:\n")
	heap.Print()

}
