package mergesort

// MergeSort sorts an array using merge sort algorithm.
// Time O(n*log(n)), since we divide array to sub-arrays until sub-array become 1 length,
// and then merge each sub array in sorted order
// Space O(n), since we use sub-arrays that sum is equal an input.
func MergeSort(arr []int) []int {
	const minArrSize = 2

	if len(arr) < minArrSize {
		return arr
	}

	subArr1 := MergeSort(arr[:len(arr)/2])
	subArr2 := MergeSort(arr[len(arr)/2:])

	return Merge(subArr1, subArr2)
}

func Merge(subArr1, subArr2 []int) []int {
	result := make([]int, 0)
	subArr1pointer, subArr2pointer := 0, 0

	for subArr1pointer < len(subArr1) && subArr2pointer < len(subArr2) {
		if subArr1[subArr1pointer] < subArr2[subArr2pointer] {
			result = append(result, subArr1[subArr1pointer])
			subArr1pointer++
		} else {
			result = append(result, subArr2[subArr2pointer])
			subArr2pointer++
		}
	}

	if subArr1pointer < len(subArr1) {
		result = append(result, subArr1[subArr1pointer:]...)
	}

	if subArr2pointer < len(subArr2) {
		result = append(result, subArr2[subArr2pointer:]...)
	}

	return result
}
