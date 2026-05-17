package quicksort

// QuickSort sorts an array using quick sort algorithm.
// Time O(n*log(n)), on average, the pivot divides the array into two parts, but not necessarily equal
// Space O(log(n)), on average, due to recursion stack depth in balanced partitions.
func QuickSort(arr []int) {
	if len(arr) == 0 {
		return
	}

	var i, j int

	for pivot := len(arr) - 1; j <= pivot; j++ {
		if arr[j] <= arr[pivot] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	QuickSort(arr[i:])
	QuickSort(arr[:i-1])
}
