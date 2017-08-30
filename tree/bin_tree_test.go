package tree

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func DemoTree() *Node {
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

func visit(node *Node) {
	fmt.Printf("%s ", string(node.value))
}

func TestTraverse(_ *testing.T) {
	root := DemoTree()
	root.Print()

	bfsResult := []rune{'T', 'S', 'R', 'P', 'N', 'O', 'A', 'E', 'I', 'H', 'G'}
	dfsResult := []rune{'T', 'S', 'P', 'E', 'I', 'N', 'H', 'G', 'R', 'O', 'A'}

	var values []rune
	visit := func(node *Node) { values = append(values, node.value) }

	BFS(root, visit)
	if !reflect.DeepEqual(bfsResult, values) {
		log.Fatal("bfs not equal. want: ", string(bfsResult), ", get: ", string(values))
	}

	values = []rune{}

	DFS(root, visit)
	if !reflect.DeepEqual(dfsResult, values) {
		log.Fatal("dfs not equal. want: ", string(dfsResult), ", get: ", string(values))
	}
}
