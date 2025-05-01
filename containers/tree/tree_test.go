package tree

import (
	"fmt"
	"testing"
)

func TestInOrder(t *testing.T) {

	// Create a sample BST:
	//         8
	//       /   \
	//      3     10
	//     / \      \
	//    1   6      14
	//       /
	//      4

	root := &TreeNode{Val: 8}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 10}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 6}
	root.Left.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 14}
	fmt.Println("Inorder Traversal (Sorted Order):")
	InOrder(root) // Expected output: 1 3 4 6 8 10 14
	fmt.Println()
}

func TestPreOrder(t *testing.T) {
	// Create a sample BST:
	//         8
	//       /   \
	//      3     10
	//     / \      \
	//    1   6      14
	//       /
	//      4
	root := &TreeNode{Val: 8}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 10}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 6}
	root.Left.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 14}

	fmt.Println("Preorder Traversal (Root First):")
	PreOrder(root) // Expected output: 8 3 1 6 4 10 14
	fmt.Println()
}

func TestPostOrder(t *testing.T) {
	// Create a sample BST:
	//         8
	//       /   \
	//      3     10
	//     / \      \
	//    1   6      14
	//       /
	//      4
	root := &TreeNode{Val: 8}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 10}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 6}
	root.Left.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 14}

	fmt.Println("Postorder Traversal (Root Last):")
	PostOrder(root) // Expected output: 1 4 6 3 14 10 8
	fmt.Println()
}
