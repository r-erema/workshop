package bubblesort

// BubbleSort sorts an array using bubble sort algorithm.
// Average, Worst: O(n^2) time | O(1) space.
func BubbleSort(array *[]int) {
	arr := *array
	boundary := len(arr)

	for boundary > 1 {
		for i := 0; i+1 < boundary; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}

		boundary--
	}
}
