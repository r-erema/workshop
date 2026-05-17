package mergetwosortedlists

type (
	ListNode struct {
		Value int
		Next  *ListNode
	}
	TestCase struct {
		Name  string
		List1 *ListNode
		List2 *ListNode
		Want  *ListNode
	}
)

// MergeTwoSortedLists merges two sorted linked lists.
// Time O(n+m)
// n = number of nodes in list1
// m = number of nodes in list2
//
// Space O(1)
// we have a constant space, since we are just shifting the pointers.
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

func NormalLists() TestCase {
	const (
		val1 = 1
		val2 = 2
		val3 = 3
		val4 = 4
	)

	return TestCase{
		Name: "Normal lists",
		List1: &ListNode{
			Value: val1,
			Next: &ListNode{
				Value: val2,
				Next: &ListNode{
					Value: val4,
					Next:  nil,
				},
			},
		},
		List2: &ListNode{
			Value: val1,
			Next: &ListNode{
				Value: val3,
				Next: &ListNode{
					Value: val4,
					Next:  nil,
				},
			},
		},
		Want: &ListNode{
			Value: val1,
			Next: &ListNode{
				Value: val1,
				Next: &ListNode{
					Value: val2,
					Next: &ListNode{
						Value: val3,
						Next: &ListNode{
							Value: val4,
							Next: &ListNode{
								Value: val4,
								Next:  nil,
							},
						},
					},
				},
			},
		},
	}
}

func ListIsNil() TestCase {
	const val1 = 1

	return TestCase{
		Name: "List is nil",
		List1: &ListNode{
			Value: val1,
			Next:  nil,
		},
		List2: nil,
		Want: &ListNode{
			Value: val1,
			Next:  nil,
		},
	}
}

func ListHasNegativeNumber() TestCase {
	const (
		valNeg9 = -9
		val3    = 3
		val5    = 5
		val7    = 7
	)

	return TestCase{
		Name: "List has a negative number",
		List1: &ListNode{
			Value: valNeg9,
			Next: &ListNode{
				Value: val3,
				Next:  nil,
			},
		},
		List2: &ListNode{
			Value: val5,
			Next: &ListNode{
				Value: val7,
				Next:  nil,
			},
		},
		Want: &ListNode{
			Value: valNeg9,
			Next: &ListNode{
				Value: val3,
				Next: &ListNode{
					Value: val5,
					Next: &ListNode{
						Value: val7,
						Next:  nil,
					},
				},
			},
		},
	}
}
