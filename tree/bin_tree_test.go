package tree

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func DemoTree() *BinaryTree {
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

	return &BinaryTree{root: root}
}

func visit(node *Node) {
	fmt.Printf("%s ", string(node.value))
}

func TestBFS(_ *testing.T) {
	tree := DemoTree()
	tree.Print()

	bfsResult := []rune{'T', 'S', 'R', 'P', 'N', 'O', 'A', 'E', 'I', 'H', 'G'}

	var values []rune
	visit := func(node *Node) { values = append(values, node.value) }

	tree.BFS(visit)
	if !reflect.DeepEqual(bfsResult, values) {
		log.Fatal("bfs not equal. want: ", string(bfsResult), ", get: ", string(values))
	}
}

func TestDFS(_ *testing.T) {
	tree := DemoTree()
	tree.Print()

	dfsResult := []rune{'T', 'S', 'P', 'E', 'I', 'N', 'H', 'G', 'R', 'O', 'A'}

	var values []rune
	visit := func(node *Node) { values = append(values, node.value) }

	tree.DFS(visit)
	if !reflect.DeepEqual(dfsResult, values) {
		log.Fatal("dfs not equal. want: ", string(dfsResult), ", get: ", string(values))
	}
}
