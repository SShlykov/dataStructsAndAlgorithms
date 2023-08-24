package simple_structs

import "fmt"

type BSTNode struct {
	Value int
	Left  *BSTNode
	Right *BSTNode
}

type BinarySearchTree struct {
	Root *BSTNode
}

func (bst *BinarySearchTree) Insert(value int) {
	bst.Root = insert(bst.Root, value)
}

func insert(node *BSTNode, value int) *BSTNode {
	if node == nil {
		return &BSTNode{Value: value}
	}

	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else if value > node.Value {
		node.Right = insert(node.Right, value)
	}
	return node
}

func (bst *BinarySearchTree) InOrder() {
	inOrder(bst.Root)
}

func inOrder(BSTNode *BSTNode) {
	if BSTNode != nil {
		inOrder(BSTNode.Left)
		fmt.Print(BSTNode.Value, ",")
		inOrder(BSTNode.Right)
	}

}

func Bst() {
	tree := &BinarySearchTree{}
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(6)
	fmt.Print("bst: [ ")
	tree.InOrder()
	fmt.Print("] \n")
}
