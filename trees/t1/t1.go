package t1

import (
	"container/list"
	"errors"
)

// simple bin tree
// TreeNode represents a node in the binary tree
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

func SumNodes(root *TreeNode) int {

	if root == nil {
		return 0
	}
	queue := NewQueue()
	sum := 0
	queue.Push(root)

	for !queue.Empty() {
		v, _ := queue.Pop()
		sum += v.Val

		if v.Left != nil {
			queue.Push(v.Left)
		}

		if v.Right != nil {
			queue.Push(v.Right)
		}

	}

	return sum
}
