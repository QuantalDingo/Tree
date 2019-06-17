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

func (t *Tree) Delete(value int) {
	t.root = deleteNode(t.root, value)
}

func deleteNode(node *Node, value int) *Node {
	if value > node.value {
		node.right = deleteNode(node.right, value)
	} else if value < node.value {
		node.left = deleteNode(node.left, value)
	} else {
		if node.right == nil && node.left == nil {
			return nil
		}

		if node.left == nil {
			return node.right
		}

		if node.right == nil {
			return node.left
		}

		node.value = findSmallestValue(*node.right)
		node.right = deleteNode(node.right, node.value)
	}

	return node // node == nil
}

func findSmallestValue(node Node) int {
	if node.left == nil {
		return node.value
	}
	return findSmallestValue(*node.left)
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
