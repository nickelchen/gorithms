package tree

/*

max heap.

Insert: insert a node to arbitrary position, then swim it up.
Remove: remove the top root. move arbitrary node to root, then sink it down.

After either operation, the result heap is still max heap.

*/

type BinaryHeap struct {
	// heap is a special binary tree. so imbeded it here.
	BinaryTree
}

// Remove the max top node from the heap, then rebalance it.
func (heap *BinaryHeap) RemoveMax() *Node {
	if heap.root == nil {
		return nil
	}

	var nodes []*Node
	heap.BFS(func(node *Node) {
		nodes = append(nodes, node)
	})

	root := heap.root
	if len(nodes) == 1 {
		// only one root node.
		heap.root = nil
		return root
	}

	// get last node. swap it with root. then delete the last node.
	leaf := nodes[len(nodes)-1]
	leaf.Value, root.Value = root.Value, leaf.Value

	// leaf Value now is swapped to root,
	// it is useless, remove it.
	leaf.parent.DetachL(leaf)

	sink(root)

	return leaf
}

// Insert a node Value to the heap. then rebalance it.
func (heap *BinaryHeap) Insert(node *Node) {
	if heap.root == nil {
		heap.root = node
		return
	}

	var nodes []*Node
	heap.BFS(func(node *Node) {
		nodes = append(nodes, node)
	})

	leaf := nodes[len(nodes)-1]
	leaf.AttachL(node)

	swim(node)
}

func NewBinaryHeap() *BinaryHeap {
	heap := BinaryHeap{}
	return &heap
}

func swim(node *Node) {
	// swim up node to proper position
	if node.parent == nil || node.parent.Value > node.Value {
		return
	}

	node.parent.Value, node.Value = node.Value, node.parent.Value
	swim(node.parent)
}

func sink(node *Node) {
	// sink down node to proper position
	var next *Node

	if node.left != nil && node.right != nil {
		if node.left.Value > node.right.Value {
			next = node.left
		} else {
			next = node.right
		}

	} else if node.left != nil {
		next = node.left
	} else if node.right != nil {
		next = node.right
	}

	if next == nil || node.Value > next.Value {
		return
	}

	node.Value, next.Value = next.Value, node.Value
	node = next

	sink(node)
}
