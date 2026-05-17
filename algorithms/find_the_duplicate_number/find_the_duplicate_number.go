package findtheduplicatenumber

// FindDuplicate finds the duplicate number in the array.
// Time O(N), since we iterate input one time
// Space O(1), since we don't involve any additional data structure.
func FindDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}

	for slow = 0; slow != fast; {
		slow, fast = nums[slow], nums[fast]
	}

	return slow
}
