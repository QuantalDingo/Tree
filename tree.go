package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) insert(value int) {
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

func (t *Tree) traverse() {
	inOrderTraverse(t.root)
}

func inOrderTraverse(node *Node) {
	if node != nil {
		inOrderTraverse(node.left)
		fmt.Println(node.value)
		inOrderTraverse(node.right)
	}
}

func main() {
	tree := new(Tree)
	tree.insert(5)
	tree.insert(4)
	tree.insert(8)
	tree.traverse()
}
