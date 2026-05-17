package merge2nondecreasingarrays

// MergeArrays merges two non-decreasing arrays into one sorted array.
// Time O(m+n), where m and n are length of the arrays, and we need to iterate each of element
// Space O(m+n), since we create resulted array which equals m+n of lengths of input arrays.
func MergeArrays(arr1, arr2 []int) []int {
	pointer1, pointer2 := 0, 0
	result := make([]int, 0)

	for pointer1 < len(arr1) && pointer2 < len(arr2) {
		if arr1[pointer1] > arr2[pointer2] {
			result = append(result, arr2[pointer2])
			pointer2++
		} else {
			result = append(result, arr1[pointer1])
			pointer1++
		}
	}

	if len(arr1) > pointer1 {
		result = append(result, arr1[pointer1:]...)
	}

	if len(arr2) > pointer2 {
		result = append(result, arr2[pointer2:]...)
	}

	return result
}
