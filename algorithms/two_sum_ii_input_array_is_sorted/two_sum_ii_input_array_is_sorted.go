package twosumiiinputarrayissorted

// TwoSum Time O(n), since in the worst case we don't exceed iterations count more then input
// Space O(1), since we don't involve an additional data structure.
func TwoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left <= right {
		sum := numbers[left] + numbers[right]

		switch {
		case sum > target:
			right--
		case sum < target:
			left++
		default:
			return []int{left + 1, right + 1}
		}
	}

	return nil
}
