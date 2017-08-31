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
	root := &Node{Value: 'T'}

	// child level 1
	c11 := &Node{Value: 'S'}
	c12 := &Node{Value: 'R'}

	// child level 2
	c21 := &Node{Value: 'P'}
	c22 := &Node{Value: 'N'}
	c23 := &Node{Value: 'O'}
	c24 := &Node{Value: 'A'}

	// child level 3
	c31 := &Node{Value: 'E'}
	c32 := &Node{Value: 'I'}
	c33 := &Node{Value: 'H'}
	c34 := &Node{Value: 'G'}

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

	heap := NewBinaryHeap()
	heap.Insert(root)

	return heap
}

func TestBinaryHeap(_ *testing.T) {
	heap := DemoBinaryHeap()
	heap.Print()

	node := &Node{Value: 'Z'}
	heap.Insert(node)

	fmt.Printf("after insert Z:\n")
	heap.Print()

	heap.RemoveMax()

	fmt.Printf("after remove top:\n")
	heap.Print()

}
