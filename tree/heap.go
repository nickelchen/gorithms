package tree

/*

max heap.

Insert: insert a node to arbitrary position, then swim it up.
Remove: remove the top root. move arbitrary node to root, then sink it down.

After either operation, the result heap is still max heap.

*/

type Heap struct {
	// heap is a special binary tree. so imbeded it here.
	BinaryTree
}

// Remove the max top node from the heap, then rebalance it.
func (heap Heap) RemoveMax() {
	if heap.root == nil {
		return
	}
	root := heap.root

	var nodes []*Node
	heap.BFS(func(node *Node) {
		nodes = append(nodes, node)
	})
	if len(nodes) == 1 {
		// only one root node.
		heap.root = nil
	}

	// get last node. swap it with root. then delete the last node.
	leaf := nodes[len(nodes)-1]
	leaf.value, root.value = root.value, leaf.value

	// leaf value now is swapped to root,
	// it is useless, remove it.
	leaf.parent.DetachL(leaf)

	sink(root)
}

// Insert a node value to the heap. then rebalance it.
func (heap Heap) Insert(node *Node) {
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

func swim(node *Node) {
	// swim up node to proper position
	if node.parent == nil || node.parent.value > node.value {
		return
	}

	node.parent.value, node.value = node.value, node.parent.value
	swim(node.parent)
}

func sink(node *Node) {
	// sink down node to proper position
	var next *Node

	if node.left != nil && node.right != nil {
		if node.left.value > node.right.value {
			next = node.left
		} else {
			next = node.right
		}

	} else if node.left != nil {
		next = node.left
	} else if node.right != nil {
		next = node.right
	}

	if next == nil || node.value > next.value {
		return
	}

	node.value, next.value = next.value, node.value
	node = next

	sink(node)
}
