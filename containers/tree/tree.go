package tree

import "fmt"

type TreeNode struct {
	Val   int       // Value stored in the node
	Left  *TreeNode // Pointer to left child
	Right *TreeNode // Pointer to right child
}

func InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder(root.Left)       // Traverse Left subtree
	fmt.Print(root.Val, " ") // Visit current node
	InOrder(root.Right)
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ") // Visit current node
	PreOrder(root.Left)      // Traverse left subtree
	PreOrder(root.Right)     // Traverse right subtree
}

func PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrder(root.Left)     // Traverse left subtree
	PostOrder(root.Right)    // Traverse right subtree
	fmt.Print(root.Val, " ") // Visit current node
}
