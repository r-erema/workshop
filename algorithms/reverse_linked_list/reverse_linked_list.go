package reverselinkedlist

type (
	ListNode struct {
		Val  int
		Next *ListNode
	}
)

// ReverseList reverses a linked list.
// Time O(n), since we should iterate all the input
// Space O(1), we allocate a new linked list, but we reduce an input linked list.
func ReverseList(head *ListNode) *ListNode {
	dummyNode := &ListNode{}

	for head != nil {
		dummyNode.Next = &ListNode{Val: head.Val, Next: dummyNode.Next}
		head = head.Next
	}

	return dummyNode.Next
}
