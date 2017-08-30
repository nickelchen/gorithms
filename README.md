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

func sortNumbers() {
	var arr := []int {100, 200, 1, 10}

	sort.QuickSort(arr)

	fmt.Printf("sorted arr: %s", arr)
}

```

tree & heap

```
import (
	"fmt"
	"github.com/nickelchen/gorithms/tree"
)

func treeDemo() {
	root := &tree.Node{value: 'T'}

	// child level 1
	c11 := &tree.Node{value: 'S'}
	c12 := &tree.Node{value: 'R'}

	root.AttachL(c11)
	root.AttachR(c12)

	visit := func(node *tree.Node) { fmt.Printf("%s "), node.value }

	bin_tree1 := tree.NewBinaryTree(root)
	bin_tree1.BFS(visit)
	bin_tree1.DFS(visit)
}

func heapDemo() {
	root := &tree.Node{value: 'T'}

	// child level 1
	c11 := &tree.Node{value: 'S'}
	c12 := &tree.Node{value: 'R'}

	root.AttachL(c11)
	root.AttachR(c12)

	node := &Node{value: 'Z'}

	heap1 := tree.NewHeap(root)
	heap1.Insert(node)
	heap1.RemoveMax()
}
```

node struct

```

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	value  rune
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

tree:

```
func (tree BinaryTree) BFS(visit func(*Node))
func (tree BinaryTree) DFS(visit func(*Node))
func (tree BinaryTree) Print()
```

heap:
```
func (heap Heap) RemoveMax()
func (heap Heap) Insert(node *Node)
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
```

### Test

```
go test ./...
```

### License

MIT
