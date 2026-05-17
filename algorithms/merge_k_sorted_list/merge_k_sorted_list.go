package mergeksortedlist

type ListNode struct {
	Value int
	Next  *ListNode
}

// MergeKSortedLists merges k sorted linked lists.
// Time O(N * logK) since we merge K lists with N numbers of nodes
// Space O(N) since mergedLists doesn't consume more memory than input.
func MergeKSortedLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var list1, list2 *ListNode

	for len(lists) > 1 {
		mergedLists := make([]*ListNode, 0)

		for i := 0; i < len(lists); i += 2 {
			list1, list2 = lists[i], nil
			if i+1 < len(lists) {
				list2 = lists[i+1]
			}

			mergedLists = append(mergedLists, MergeTwoSortedLists(list1, list2))
		}

		lists = mergedLists
	}

	return lists[0]
}

func MergeTwoSortedLists(list1, list2 *ListNode) *ListNode {
	dummyNode := &ListNode{}
	tail := dummyNode

	for list1 != nil && list2 != nil {
		if list1.Value < list2.Value {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	}

	if list2 != nil {
		tail.Next = list2
	}

	return dummyNode.Next
}
