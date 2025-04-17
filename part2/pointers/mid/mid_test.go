package mid

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	//Create a linked list with nodes containing values 1, 2, 3, 4, and 5

	//drow this out...
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	middle := FindMiddle(head)
	if middle != nil {
		fmt.Println(middle.Val)
	}
	//fmt.Println(middle)
}

func Test2(t *testing.T) {
	//Create a linked list with nodes containing values 1, 2, 3, 4, 5,6

	//drow this out...
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	middle := FindMiddle(head)
	if middle != nil {
		fmt.Println(middle.Val)
	}
	//fmt.Println(middle)
}
