package main

import (
	"container/list"
	"errors"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//simple queue

type Queue struct {
	data *list.List //use the builtin list
}

func NewQueue() *Queue {
	return &Queue{data: list.New()}
}

func (q *Queue) Depth() int {
	return q.data.Len()
}

func (q *Queue) Empty() bool {
	return q.data.Len() == 0
}

func (q *Queue) Peek() (*TreeNode, error) {
	if q.Empty() {
		return nil, errors.New("queue is empty")
	}
	return q.data.Front().Value.(*TreeNode), nil
}

func (q *Queue) Push(value *TreeNode) {
	q.data.PushBack(value)
}

func (q *Queue) Pop() (*TreeNode, error) {
	if q.Empty() {
		return nil, errors.New("queue is empty")
	}
	v := q.data.Remove(q.data.Front())
	return v.(*TreeNode), nil
}

func traverse(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	//queue not recursion...
	queue := NewQueue()
	//push the root
	queue.Push(root)
	for !queue.Empty() {
		lengthOfLevel := queue.Depth()
		fmt.Println("Depth of queue/width of level:", lengthOfLevel)
		currentLevel := make([]int, 0, lengthOfLevel)
		for i := 0; i < lengthOfLevel; i++ {
			value, _ := queue.Pop()
			currentLevel = append(currentLevel, value.Val)
			if value.Left != nil {
				queue.Push(value.Left)
			}

			if value.Right != nil {
				queue.Push(value.Right)
			}
		}
		result = append(result, currentLevel)
	}

	return result
}

func main() {
	/*
	   Node root = new Node(5);
	      root.left = new Node(12);
	      root.right = new Node(13);

	      root.left.left = new Node(7);
	      root.left.right = new Node(14);

	      root.right.right = new Node(2);

	      root.left.left.left = new Node(17);
	      root.left.left.right = new Node(23);

	      root.left.right.left = new Node(27);
	      root.left.right.right = new Node(3);

	      root.right.right.left = new Node(8);
	      root.right.right.right = new Node(11);
	*/
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 12}
	root.Right = &TreeNode{Val: 13}

	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: 14}

	root.Right.Right = &TreeNode{Val: 2}

	root.Left.Left.Left = &TreeNode{Val: 17}
	root.Left.Left.Right = &TreeNode{Val: 23}

	root.Left.Right.Left = &TreeNode{Val: 27}
	root.Left.Right.Right = &TreeNode{Val: 13}

	root.Right.Right.Left = &TreeNode{Val: 8}
	root.Right.Right.Right = &TreeNode{Val: 11}
	/*
		root.Left = &TreeNode{Val: 7}
		root.Right = &TreeNode{Val: 1}
		root.Left.Left = &TreeNode{Val: 9}
		root.Right.Left = &TreeNode{Val: 10}
		root.Right.Right = &TreeNode{Val: 5}
	*/

	output := traverse(root)
	fmt.Println(output)

}
