package removenthnodefromendoflist

import (
	linkedlist "github.com/r-erema/workshop/utils/data_structure/linked_list"
)

// RemoveNthFromEnd removes nth node from the end of a linked list.
// Time O(n), since we need to iterate each node 1 time
// Space O(1), since we don't involve any additional data structure.
func RemoveNthFromEnd(head *linkedlist.Node, n int) *linkedlist.Node {
	dummy := &linkedlist.Node{Next: head}
	left, right := dummy, head

	for ; n > 0; n-- {
		right = right.Next
	}

	for right != nil {
		left, right = left.Next, right.Next
	}

	left.Next = left.Next.Next

	return dummy.Next
}
