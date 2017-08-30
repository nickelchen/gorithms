package tree

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree) BFS(visit func(*Node)) {
	if tree.root == nil {
		return
	}

	root := tree.root
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

func (tree *BinaryTree) DFS(visit func(*Node)) {
	if tree.root == nil {
		return
	}

	dfs(tree.root, visit)
}

func NewBinaryTree(root *Node) *BinaryTree {
	bintree := BinaryTree{
		root: root,
	}
	return &bintree
}

func dfs(root *Node, visit func(*Node)) {
	visit(root)
	if root.left != nil {
		dfs(root.left, visit)
	}
	if root.right != nil {
		dfs(root.right, visit)
	}
}

func (tree *BinaryTree) Print() {
	if tree.root == nil {
		return
	}

	tree.root.Print()
}
