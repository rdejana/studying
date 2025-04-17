package mid

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindMiddle(head *ListNode) *ListNode {

	slow := head
	fast := head

	//need to base on fast
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
