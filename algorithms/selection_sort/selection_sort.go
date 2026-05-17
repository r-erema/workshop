package selectionsort

// SelectionSort sorts an array using selection sort algorithm.
// Average, Worst: O(n^2) time | O(1) space.
func SelectionSort(array *[]int) {
	arr := *array
	startIndex, length := 0, len(arr)

	for startIndex < length {
		minIndex := startIndex
		for i := startIndex; i < len(arr); i++ {
			if arr[minIndex] > arr[i] {
				minIndex = i
			}
		}

		arr[startIndex], arr[minIndex] = arr[minIndex], arr[startIndex]
		startIndex++
	}
}
