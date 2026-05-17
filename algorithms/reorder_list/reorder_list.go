package reorderlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReorderList reorders a linked list by interleaving first and second halves.
// Time O(n), since the iteration count depends on input linearly
// Space O(1), we don't use any extra space.
func ReorderList(head *ListNode) {
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	secondPart := slow.Next
	slow.Next = nil

	var secondReversedPart *ListNode

	for secondPart != nil {
		tmp := secondPart.Next
		secondPart.Next = secondReversedPart
		secondReversedPart = secondPart
		secondPart = tmp
	}

	curr := head
	for secondReversedPart != nil {
		tmp, tmp2 := curr.Next, secondReversedPart.Next
		curr.Next = secondReversedPart
		curr.Next.Next = tmp
		curr, secondReversedPart = curr.Next.Next, tmp2
	}
}
