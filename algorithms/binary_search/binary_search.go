package binarysearch

// BinarySearch searches for a needle in a sorted array and returns its index.
// Average, Worst: O(log n) time | O(1) space.
func BinarySearch(array []int, needle int) int {
	leftPointer, rightPointer := 0, len(array)-1

	const two = 2

	for leftPointer <= rightPointer {
		cutPoint := (leftPointer + rightPointer) / two
		potentialResult := array[cutPoint]

		switch {
		case potentialResult > needle:
			rightPointer = cutPoint - 1
		case potentialResult < needle:
			leftPointer = cutPoint + 1
		default:
			return cutPoint
		}
	}

	return -1
}
