algorithms and data structure, implemented in golang.

### Install

```
// install all packages under this repository
go get -u github.com/nickelchen/gorithms/...
```

### Usage

sort

```
import (
	"fmt"
	"github.com/nickelchen/gorithms/sort"
)

func sortDemo() {
	var arr := []int {100, 200, 1, 10}

	sort.QuickSort(arr)

	fmt.Printf("sorted arr: %s", arr)
}

```

binary tree & heap

```
import (
	"fmt"
	"github.com/nickelchen/gorithms/tree"
)

func treeDemo() {
	root := &tree.Node{Value: 'T'}

	// child level 1
	c11 := &tree.Node{Value: 'S'}
	c12 := &tree.Node{Value: 'R'}

	root.AttachL(c11)
	root.AttachR(c12)

	visit := func(node *tree.Node) { fmt.Printf("%s "), node.Value }

	bin_tree1 := tree.NewBinaryTree(root)
	bin_tree1.BFS(visit)
	bin_tree1.DFS(visit)
}

func heapDemo() {
	root := &tree.Node{Value: 'T'}

	// child level 1
	c11 := &tree.Node{Value: 'S'}
	c12 := &tree.Node{Value: 'R'}

	root.AttachL(c11)
	root.AttachR(c12)

	// create an empty heap.
	heap1 := tree.NewBinaryHeap()
	// insert the root node. now heap has 3 nodes
	heap1.Insert(root)

	node := &Node{Value: 'Z'}
	// insert to last. then swim up to right position('Z' is the top)
	heap1.Insert(node)

	// remove the top
	node2 := heap1.RemoveMax()
	fmt.Printf("%s", node2.Value)
}
```

node struct

```

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	Value  rune
}

```

print tree

```
root.Print()

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

```

on a same level, the upper node is left child, the below node is right child.
eg: `S` is left, `R` is right child of node `T` ; `E` is left, `I` is right child of node `P`

### API

binary tree:

```
func (tree *BinaryTree) Print()
func (tree *BinaryTree) BFS(visit func(*Node))
func (tree *BinaryTree) DFS(visit func(*Node))
```

binary heap, subclass binary tree:
```
func (heap *BinaryHeap) RemoveMax() *Node
func (heap *BinaryHeap) Insert(node *Node)
```

node:
```
func (node *Node) Print()
func (node *Node) AttachL(child *Node)
func (node *Node) AttachR(child *Node)
func (node *Node) DetachL(child *Node)
func (node *Node) DetachR(child *Node)
```

sort:
```
func QuickSort(numbers []int)
func InsertSort(numbers []int)
func ShellSort(numbers []int)
func SelectSort(numbers []int)
func BubbleSort(numbers []int)
func MergeSort(numbers []int) []int
func HeapSort(numbers []int) []int
```

### Test

```
go test ./...
```

### License

MIT
