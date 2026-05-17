package kthlargestelementinanarray

// FindKthLargest finds the kth largest element in an array.
// Time O(n) in average case, random pivots split the array roughly in half each time,
// leading to a geometric series of work: n + n/2 + n/4 + … = O(n).
// O(n^2) in worst case since the pivot is always the smallest or largest element,
// causing each partition to reduce the problem size by only one,
// leading to n + (n−1) + (n−2) + … + 1 = O(n²) total operations
//
// Space O(1), since we use only an initial array.
func FindKthLargest(arr []int, k int) int {
	var pointer int

	for j, pivot := 0, len(arr)-1; j <= pivot; j++ {
		if arr[j] <= arr[pivot] {
			arr[pointer], arr[j] = arr[j], arr[pointer]
			pointer++
		}
	}

	if pointer == len(arr) {
		pointer--
	}

	for len(arr) > 1 && arr[0] == arr[1] {
		arr = arr[1:]
		pointer--
	}

	if len(arr) == 1 {
		return arr[0]
	}

	if pointer <= len(arr)-k {
		return FindKthLargest(arr[pointer:], k)
	}

	if k <= 1 {
		return arr[len(arr)-1]
	}

	k -= len(arr) - len(arr[:pointer])

	return FindKthLargest(arr[:pointer], k)
}
