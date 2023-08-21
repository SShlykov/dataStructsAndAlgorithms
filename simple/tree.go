package main

import (
	"fmt"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func tree() {
	root := &TreeNode{Value: 10}

	root.Left = &TreeNode{Value: 5}
	root.Right = &TreeNode{Value: 15}

	root.Left.Left = &TreeNode{Value: 3}
	root.Left.Right = &TreeNode{Value: 7}

	root.Right.Left = &TreeNode{Value: 12}
	root.Right.Right = &TreeNode{Value: 18}

	printTree(root, "", "root")
	fmt.Print("\n")
}

func printTree(node *TreeNode, indent string, direction string) {
	if node == nil {
		return
	}
	prefix := indent
	if direction == "left" {
		prefix = prefix + "--"
	} else if direction == "right" {
		prefix = prefix + "--"
	} else {
		prefix = prefix + " "
	}
	printTree(node.Right, indent+"   |", "right")
	fmt.Println("tree: ", prefix, node.Value)
	if direction == "left" || direction == "right" {
		fmt.Println("tree: ", indent)
	}
	printTree(node.Left, indent+"   |", "left")
}
