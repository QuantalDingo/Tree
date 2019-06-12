package tree

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) Insert(value int) {
	tmp := &Node{value, nil, nil}

	if t.root == nil {
		t.root = tmp
	} else {
		t.root.insertNode(tmp)
	}
}

func (t *Node) insertNode(node *Node) {
	if t.value < node.value {
		if t.right == nil {
			t.right = node
		} else {
			t.right.insertNode(node)
		}
	} else {
		if t.value > node.value {
			if t.left == nil {
				t.left = node
			} else {
				t.left.insertNode(node)
			}
		}
	}
}

func containsValue(node *Node, val int) bool {
	if node == nil {
		return false
	}

	if node.value == val {
		return true
	}

	if node.value < val {
		return containsValue(node.left, val)
	} else {
		return containsValue(node.right, val)
	}
}

func (t *Tree) contains(val int) bool {
	return containsValue(t.root, val)
}

func (t *Tree) Delete(value int) bool {
	return deleteValue(t.root, value)
}

func deleteValue(node *Node, value int) bool {
	if node == nil {
		return true
	}

	if node.value == value {
		if node.right == nil && node.left == nil {
			node = nil
			return true
		}

		if node.left == nil {
			node = node.right
			return true
		}

		if node.right == nil {
			node = node.left
			return true
		}

	}

	if value > node.value {
		deleteValue(node.right, value)
	} else {
		deleteValue(node.left, value)
	}

	return false
}

func findSmallestValue(node Node) int {
	if node.left == nil {
		return node.value
	} else {
		return findSmallestValue(*node.left)
	}
}

func (t *Tree) Traverse(f func(int)) {
	inOrderTraverse(t.root, f)
}

func inOrderTraverse(node *Node, f func(int)) {
	if node != nil {
		inOrderTraverse(node.left, f)
		f(node.value)
		inOrderTraverse(node.right, f)
	}
}
