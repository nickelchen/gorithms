package tree

type Node struct {
	parent *Node
	left   *Node
	right  *Node
	value  rune
}

func (node *Node) print(prefix string, right bool) {
	var lead string
	var space string

	if right {
		lead = "└── "
		space = "    "
	} else {
		lead = "├── "
		space = "│   "
	}

	fmt.Println(prefix + lead + string(node.value))

	if node.left != nil {
		node.left.print(prefix+space, false)
	}
	if node.right != nil {
		node.right.print(prefix+space, true)
	}
}

func (node *Node) Print() {
	node.print("", true)
}

func AttachL(root *Node, child *Node) {
	if root.left != nil {
		DetachL(root, root.left)
	}

	root.left = child
	child.parent = root
}

func AttachR(root *Node, child *Node) {
	if root.right != nil {
		DetachR(root, root.right)
	}

	root.right = child
	child.parent = root
}

func DetachL(root *Node, node *Node) {
	if root != nil {
		root.left = nil
	}
	node.parent = nil
}

func DetachR(root *Node, node *Node) {
	if root != nil {
		root.right = nil
	}
	node.parent = nil
}

func BFS(root *Node, visit func(*Node)) {
	queue := []*Node{root}

	for 0 < len(queue) {
		var nextQueue []*Node
		for _, n := range queue {
			visit(n)
			if n.left != nil {
				nextQueue = append(nextQueue, n.left)
			}
			if n.right != nil {
				nextQueue = append(nextQueue, n.right)
			}
		}
		queue = nextQueue
	}
}

func DFS(root *Node, visit func(*Node)) {
	visit(root)
	if root.left != nil {
		DFS(root.left, visit)
	}
	if root.right != nil {
		DFS(root.right, visit)
	}
}
