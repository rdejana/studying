package cycle

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	result := HasCycle(head)

	fmt.Println(result)

}

func Test2(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	head.Next.Next.Next.Next.Next.Next = head.Next.Next

	result := HasCycle(head)

	fmt.Println(result)

}

func Test3(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next.Next = head.Next.Next

	result := FindCycleLength(head)

	fmt.Println(result)

	head.Next.Next.Next.Next.Next.Next = head.Next.Next.Next

	result = FindCycleLength(head)

	fmt.Println(result)

}
