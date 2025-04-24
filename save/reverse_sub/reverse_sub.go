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

func Reverse[T any](p, q int, head *ListNode[T]) *ListNode[T] {
	//the example uses 1 index, but I want o use 0
	fmt.Println("P is ", p)
	pIndex := p - 1
	//qIndex := q - 1
	//find
	if p == q {
		return head
	}

	var current, previous *ListNode[T] = head, nil
	for i := 0; i < pIndex && current != nil; i++ {
		previous = current
		current = current.Next
	}

	fmt.Println(previous.Value)
	fmt.Println(current.Value)

	//so current is the start of where we go
	start := previous  //this is where we'll join
	subList := current //start of the sublist

	for i := 0; current != nil && i < q-p+1; i++ {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	//now relink
	if start != nil {
		start.Next = previous
	} else {
		head = previous
	}

	subList.Next = current

	return head
}

func ReverseLSubinkedList[T any](p, q int, head *ListNode[T]) *ListNode[T] {

	if head == nil {
		return nil
	}
	//find pa and q

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

	printList(Reverse(2, 4, head))

}
