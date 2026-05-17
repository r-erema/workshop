package maxweightofthegoodsubset

import (
	"sort"
)

// MaxWeight finds the maximum weight of a good subset.
// Time O(N^2 * log(N)), since we have a nested loop
// (we need search an appropriate sum pair for each biggest number in iteration) and also we have sorting
// Space O(1), since we don't use any additional space.
func MaxWeight(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	sort.Ints(arr)

	var result int

	for i := len(arr) - 1; i > 0; i-- {
		examineElem := arr[i]
		leftPointer, rightPointer := i-1, i
		intermediateSum := arr[rightPointer]

		for arr[leftPointer]+arr[rightPointer] >= examineElem {
			intermediateSum += arr[leftPointer]
			leftPointer--
			rightPointer--

			if leftPointer < 0 {
				break
			}
		}

		result = max(result, intermediateSum)
	}

	return result
}
