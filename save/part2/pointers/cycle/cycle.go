package cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindCycleLength(head *ListNode) int {
	slow := head
	fast := head

	for fast != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return calculateCycleLength(slow)
		}
	}

	return 0
}

func calculateCycleLength(node *ListNode) int {
	//walk to I'm back to my self
	ptr := node
	length := 0
	for {
		ptr = ptr.Next
		length++
		if ptr == node {
			break
		}
	}
	return length

}

func HasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil {

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
