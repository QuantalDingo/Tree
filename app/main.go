package main

import (
	"fmt"
)

func main() {
	tree := new(Tree)

	tree.insert(5)
	tree.insert(4)
	tree.insert(8)

	tree.traverse(func(i int) { fmt.Println(i) })
}
