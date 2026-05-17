package insertionsort

// InsertionSort sorts an array using insertion sort algorithm.
// Time O(n^2), since it's possible that we iterate the input twice for the each number
// Space O(1), since we don't allocate any additional memory.
func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := range i {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}
