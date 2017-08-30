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

func main() {
	var arr := []int {100, 200, 1, 10}

	sort.QuickSort(arr)
	
	fmt.Printf("sorted arr: %s", arr)
}

```

tree

```
import (
	"fmt"
	"github.com/nickelchen/gorithms/tree"
)

func main () {
	root := &tree.Node{value: 'T'}
	
	// child level 1
	c11 := &tree.Node{value: 'S'}
	c12 := &tree.Node{value: 'R'}
	
	tree.AttachL(root, c11)
	tree.AttachR(root, c12)
	
	visit := func(node *tree.Node) { fmt.Printf("%s "), node.value }
	
	tree.BFS(root, visit)
	tree.DFS(root, visit)
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

### Test

```
go test -v ./...
```

### License

MIT
