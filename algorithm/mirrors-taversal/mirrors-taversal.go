package main

import "fmt"

// https://www.educative.io/edpresso/what-is-morris-traversal

type node struct {
	data      interface{}
	leftNode  *node
	rightNode *node
}

func mirros(root *node) {
	var curr, prev *node
	if root == nil {
		return
	}

	curr = root

	for curr != nil {
		if curr.leftNode == nil {
			fmt.Printf("%v\n", curr.data)
			curr = curr.rightNode
			continue
		}

		// Find the prevous (prev) of curr
		prev = curr.leftNode
		for prev.rightNode != nil && prev.rightNode != curr {
			prev = prev.rightNode
		}

		// Make curr as the right child of its previous
		if prev.rightNode == nil {
			prev.rightNode = curr
			curr = curr.leftNode
			continue
		}

		// Fix the right child of previous
		prev.rightNode = nil
		fmt.Printf("%v\n", curr.data)
		curr = curr.rightNode
	}
}

func main() {
	root := &node{data: 4}
	root.leftNode = &node{data: 2}
	root.rightNode = &node{data: 5}
	root.leftNode.leftNode = &node{data: 1}
	root.leftNode.rightNode = &node{data: 3}

	mirros(root)
}
