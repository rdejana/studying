package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {

	ptr := head
	for ptr != nil {
		fmt.Println(ptr.Val)
		ptr = ptr.Next
	}
}

func simpleRotate(head *ListNode, k int) *ListNode {

	//easy mode.///
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}

	//how long is the list?
	lastNode := head
	listLength := 1
	for lastNode.Next != nil {
		lastNode = lastNode.Next
		listLength++
	}

	fmt.Println("Lenght: ", listLength)
	lastNode.Next = head //circle me

	//trick one
	k %= listLength // no need to do rotations more than the length of the list
	fmt.Println("rotations: ", k)
	//now far to go
	lastNode.Next = head // connect the last node with the head to a circular list
	k %= listLength      // no need to do rotations more than the length of the list
	skipLength := listLength - k
	lastNodeOfRotatedList := head
	for i := 0; i < skipLength-1; i++ {
		lastNodeOfRotatedList = lastNodeOfRotatedList.Next
	}

	//split things
	head = lastNodeOfRotatedList.Next
	lastNodeOfRotatedList.Next = nil
	return head
}

func main() {

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	printList(head)

	head = simpleRotate(head, 3)

	printList(head)

}
