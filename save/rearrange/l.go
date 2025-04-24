package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// reorder reorders the linked list as per the given logic
func reorder(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// find the middle of the LinkedList
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// slow is now pointing to the middle node
	headSecondHalf := reverse(slow) // reverse the second half
	headFirstHalf := head

	// rearrange to produce the LinkedList in the required order
	for headFirstHalf != nil && headSecondHalf != nil {
		temp := headFirstHalf.Next
		headFirstHalf.Next = headSecondHalf
		headFirstHalf = temp

		temp = headSecondHalf.Next
		headSecondHalf.Next = headFirstHalf
		headSecondHalf = temp
	}

	// set the next of the last node to 'null'
	if headFirstHalf != nil {
		headFirstHalf.Next = nil
	}

	return head
}

// reverse reverses the linked list
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func main() {
	head := &ListNode{Val: 2}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 8}
	head.Next.Next.Next.Next = &ListNode{Val: 10}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 12}
	reorder(head)
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}
