package data

import "fmt"

/* very very basic LinkedList */
type ListNode struct {
	Val  int
	Next *ListNode
}

/** equally basic binary tree */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BFT(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]  // Dequeue the first node
		queue = queue[1:] //lazy way to move the queue ptr

		fmt.Printf("Value %d\n", node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left) // Enqueue left child
		}
		if node.Right != nil {
			queue = append(queue, node.Right) // Enqueue right child
		}

	}

}

type BinarySearchTree struct {
	root *TreeNode
}

func (bst *BinarySearchTree) BFT() {
	BFT(bst.root)
}

func (bst *BinarySearchTree) Insert(value int) {
	newNode := &TreeNode{Val: value}
	if bst.root == nil {
		bst.root = newNode
	} else {
		bstInsert(bst.root, newNode)
	}
}

func bstInsert(node, newNode *TreeNode) {
	if newNode.Val < node.Val {
		if node.Left == nil {
			node.Left = newNode
		} else {
			bstInsert(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			bstInsert(node.Right, newNode)
		}
	}
}

//we'll use methods on this one on this one
