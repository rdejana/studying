package main

//// append the current level at the beginning
//		result = append([][]int{currentLevel}, result...)

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

func Traverse(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	queue := NewQueue()
	queue.Push(root)
	level := 1
	for !queue.Empty() {
		//so now, coming back must be the level...
		levelSize := queue.Depth() //look at this here so we know how many
		fmt.Printf("Level %d Size: %d\n", level, levelSize)
		level++
		//we have.  also helps with the size of the next array
		currentLevel := make([]int, 0, levelSize)
		for i := 0; i < levelSize; i++ {
			//pop
			value, _ := queue.Pop()
			currentLevel = append(currentLevel, value.Val)
			if value.Left != nil {
				queue.Push(value.Left)
			}
			if value.Right != nil {
				queue.Push(value.Right)
			}

		}
		//this acts like a prepend ..
		result = append([][]int{currentLevel}, result...)
		//result = append(result, currentLevel)
	}

	return result
}

func main() {
	root := &TreeNode{Val: 12}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 5}

	result := Traverse(root)

	fmt.Println(result)
}
