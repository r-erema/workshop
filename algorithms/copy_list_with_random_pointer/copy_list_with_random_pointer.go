package copylistwithrandompointer

import (
	linkedlist "github.com/r-erema/workshop/utils/data_structure/linked_list"
)

func IsClone(originalList, clonedList *linkedlist.Node) bool {
	for originalList != nil && clonedList != nil {
		if originalList == clonedList || originalList.Val != clonedList.Val {
			return false
		}

		if RandomNotValid(originalList, clonedList) {
			return false
		}

		originalList, clonedList = originalList.Next, clonedList.Next
	}

	return originalList == nil && clonedList == nil
}

func RandomNotValid(originalList, clonedList *linkedlist.Node) bool {
	randomHaveSamePointer := originalList.Prev == clonedList.Prev && originalList.Prev != nil
	randomNotNilAndHaveDiffVals := originalList.Prev != nil && clonedList.Prev != nil &&
		clonedList.Prev.Val != originalList.Prev.Val

	return randomHaveSamePointer || randomNotNilAndHaveDiffVals
}

// CopyPrevList creates a deep copy of a linked list with random pointers.
// Time O(3n), since we should iterate all the input 3 times
// Space O(1), sine we don't allocate any additional memory.
func CopyPrevList(head *linkedlist.Node) *linkedlist.Node {
	curr := head
	for curr != nil {
		tail := curr.Next
		curr.Next = &linkedlist.Node{Val: curr.Val, Next: tail, Prev: curr.Prev}
		curr = curr.Next.Next
	}

	curr = head
	for i := 0; curr != nil; i++ {
		if i%2 == 1 && curr.Prev != nil {
			curr.Prev = curr.Prev.Next
		}

		curr = curr.Next
	}

	dummy := &linkedlist.Node{}
	dummyCurr := dummy
	curr = head

	for i := 0; curr != nil; i++ {
		dummyCurr.Next = curr.Next

		if curr.Next != nil {
			curr.Next = curr.Next.Next
		}

		dummyCurr = dummyCurr.Next
		curr = curr.Next
	}

	return dummy.Next
}
