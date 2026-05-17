package addtwonumbers

import (
	linkedlist "github.com/r-erema/workshop/utils/data_structure/linked_list"
)

// AddTwoNumbers adds two numbers represented as linked lists.
// Time O(m+n), since we should iterate both lists
// Space O(1), sine we don't allocate any additional memory.
func AddTwoNumbers(list1, list2 *linkedlist.Node) *linkedlist.Node {
	curr1, curr2, carry := list1, list2, 0

	const dec = 10

	for curr1 != nil && curr2 != nil {
		val1, val2 := curr1.Val, curr2.Val
		curr1.Val += val2
		curr2.Val += val1

		curr1.Val += carry
		curr2.Val += carry
		carry = 0

		if curr1.Val >= dec {
			curr1.Val %= dec
			curr2.Val %= dec
			carry = 1
		}

		if curr1.Next == nil && curr2.Next == nil && carry == 1 {
			curr1.Next = &linkedlist.Node{Val: 1}

			return list1
		}

		curr1, curr2 = curr1.Next, curr2.Next
	}

	if list := FinishList(curr1, list1, carry); list != nil {
		return list
	}

	if list := FinishList(curr2, list2, carry); list != nil {
		return list
	}

	return list1
}

func FinishList(curr, list *linkedlist.Node, carry int) *linkedlist.Node {
	const dec = 10

	for curr != nil {
		curr.Val += carry
		carry = 0

		if curr.Val >= dec {
			curr.Val %= dec
			carry = 1
		}

		if curr.Next == nil {
			if carry == 1 {
				curr.Next = &linkedlist.Node{Val: 1}
			}

			return list
		}

		curr = curr.Next
	}

	return nil
}
