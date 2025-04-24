package data

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	fmt.Println("Breadth-First Traversal:")
	BFT(root) // Output: 1 2 3 4 5 6

}

func Test2(t *testing.T) {
	bst := &BinarySearchTree{}
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(3)
	bst.Insert(4)
	bst.Insert(5)
	bst.Insert(6)

	bst.BFT()

}
