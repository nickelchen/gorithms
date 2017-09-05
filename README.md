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

lru 

```
import (
	"fmt"
	"github.com/nickelchen/gorithms/lru"
)

func lruDemo() {
	l := lru.NewLRU(3)
	l.Put("k1", "v1")
	l.Put("k2", "v2")
	l.Put("k3", "v3")
	l.Put("k4", "v4")

	keys := l.Keys() // []string{"k4", "k3", "k2"}

	l.Get("k2")

	keys = l.Keys() // []string{"k2", "k4", "k3"}
}
```



### API

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

lru:
```
func NewLRU(capacity int)
func (lru *LRU) Put(k, v string)
func (lru *LRU) Get(k string) (v string, err error)
func (lru *LRU) Keys() []string
func (lru *LRU) Head() (v string, err error)
```

### Test

```
go test ./...
```

### License

MIT
