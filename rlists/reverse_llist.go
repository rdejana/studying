package main

import "fmt"

type ListNode[T any] struct {
	Value T
	Next  *ListNode[T]
}

func printList[T any](head *ListNode[T]) {
	node := head
	for node != nil {
		fmt.Println(node.Value)
		node = node.Next
	}
}

func NewListNode[T any](value T) *ListNode[T] {
	return &ListNode[T]{Value: value, Next: nil}
}

func ReverseLinkedList[T any](head *ListNode[T]) *ListNode[T] {

	if head == nil {
		return nil
	}

	current := head
	var prev *ListNode[T] = nil
	//var next *ListNode[T] = nil

	for current != nil {
		//next is a temp ptr that captures the
		//current nodes (that is the one moving forward)
		//next.  We'll use this to update the ptr to current later
		next := current.Next
		//set the current node's next to the previous
		current.Next = prev
		//move prev to current as
		// current will be set to next
		prev = current
		current = next
	}
	//on the last loop, current is nil
	//prev is the LAST valid value, so the new head

	return prev
}

func main() {
	//simple list, 1 to 10
	head := NewListNode[int](1)
	current := head
	for i := 2; i <= 10; i++ {
		current.Next = NewListNode(i)
		current = current.Next
	}

	printList(head)

	head = ReverseLinkedList(head)
	printList(head)

}
