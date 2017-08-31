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

	return NewBinaryTree(root)
}

func visit(node *Node) {
	fmt.Printf("%s ", string(node.Value))
}

func TestBFS(_ *testing.T) {
	tree := DemoTree()
	tree.Print()

	bfsResult := []rune{'T', 'S', 'R', 'P', 'N', 'O', 'A', 'E', 'I', 'H', 'G'}

	var values []rune
	visit := func(node *Node) { values = append(values, node.Value) }

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
	visit := func(node *Node) { values = append(values, node.Value) }

	tree.DFS(visit)
	if !reflect.DeepEqual(dfsResult, values) {
		log.Fatal("dfs not equal. want: ", string(dfsResult), ", get: ", string(values))
	}
}
