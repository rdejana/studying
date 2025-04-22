package binary_tree

import (
	"container/list"
	"errors"
	"fmt"
)

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

type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree() *TreeNode {
	return &TreeNode{}
}

func NewTreeWithKey(key int) *TreeNode {
	return &TreeNode{Key: key}
}

func (node *TreeNode) SetKey(key int) *TreeNode {
	node.Key = key
	return node
}

func Tranverse(root *TreeNode) [][]int {
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
			currentLevel = append(currentLevel, value.Key)
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

func (node *TreeNode) Search(key int) bool {
	if node.Key == key {
		return true
	}
	if node.Key < key {
		if node.Right == nil {
			return false
		}
		return node.Right.Search(key)
	} else if node.Key > key {
		if node.Left == nil {
			return false
		}
		return node.Left.Search(key)
	}

	return false
}

func (node *TreeNode) Insert(key int) {
	if node.Key < key {
		//go right
		if node.Right == nil {
			node.Right = &TreeNode{Key: key}
		} else {
			//something is here, so ask the child..
			node.Right.Insert(key)
		}
	} else if node.Key > key {
		//go left
		if node.Left == nil {
			node.Left = &TreeNode{Key: key}
		} else {
			node.Left.Insert(key)
		}
	}
}
