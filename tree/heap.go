package tree

/*

max_heap.

Insert: insert a node to arbitrary position, then swim it up.
Remove: remove the top root. move arbitrary node to root, then sink it down.

After either operation, the result heap is still max heap.

*/

func Swim(node *Node) {
	// swim up node to proper position
	if node.parent == nil || node.parent.value > node.value {
		return
	}

	node.parent.value, node.value = node.value, node.parent.value
	Swim(node.parent)
}

func Sink(node *Node) {
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

	Sink(node)
}

// Insert a node value to the heap. then rebalance it.
func Insert(root *Node, node *Node) {
	leaf := root
	for leaf.left != nil {
		leaf = leaf.left
	}

	AttachL(leaf, node)

	Swim(node)
}

// Remove the max top node from the heap, then rebalance it.
func Remove(root *Node) {
	leaf := root

	for leaf.left != nil {
		leaf = leaf.left
	}
	leaf.value, root.value = root.value, leaf.value

	// leaf value now is swapped to root,
	// now it is useless, remove it.
	DetachL(leaf.parent, leaf)

	Sink(root)
}
