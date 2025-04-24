package t1

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}

	expected := 6
	result := SumNodes(root1)

	fmt.Println("Sum of nodes in Example 1:", result) // Output: 6

	if expected != result {
		t.Errorf("unexpected result %d != %d\n", expected, result)
	}
}

func Test2(t *testing.T) {
	root2 := &TreeNode{Val: 4}
	root2.Left = &TreeNode{Val: 9}
	root2.Right = &TreeNode{Val: 7}
	root2.Left.Left = &TreeNode{Val: 2}
	root2.Left.Right = &TreeNode{Val: 6}
	expected := 28
	result := SumNodes(root2)
	fmt.Println("Sum of nodes in Example:", result)
	if expected != result {
		t.Errorf("unexpected result %d != %d\n", expected, result)
	}

}

func Test3(t *testing.T) {
	root3 := &TreeNode{Val: 10}
	root3.Left = &TreeNode{Val: 5}
	root3.Left.Left = &TreeNode{Val: 3}
	root3.Left.Right = &TreeNode{Val: 7}
	root3.Left.Right.Right = &TreeNode{Val: 9}

	expected := 34
	result := SumNodes(root3)
	fmt.Println("Sum of nodes in Example:", result)
	if expected != result {
		t.Errorf("unexpected result %d != %d\n", expected, result)
	}
}

// Example 1

// Example 2

// Example 3
