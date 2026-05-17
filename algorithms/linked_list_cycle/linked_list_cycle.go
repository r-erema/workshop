package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycle detects if a linked list has a cycle using Floyd's algorithm.
// Time O(n), since we should iterate all the input
// Space O(1), we don't allocate additional memory.
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
