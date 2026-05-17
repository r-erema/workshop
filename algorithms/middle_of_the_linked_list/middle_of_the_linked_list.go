package middleofthelinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// MiddleNode finds the middle node of a linked list.
// Time O(n), since we walk trough the list 1 time in the worst case
// Space O(1), we don't use any extra space.
func MiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	return slow
}
