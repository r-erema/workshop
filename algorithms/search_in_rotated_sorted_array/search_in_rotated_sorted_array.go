package searchinrotatedsortedarray

// Search searches for a target in a rotated sorted array.
// Time O(log(N)), since it's binary search
// Space O(1), sine we don't allocate any additional memory.
func Search(nums []int, target int) int {
	const divisor = 2

	left, right := 0, len(nums)-1

	isLeftPortionSorted := func(mid int) bool {
		return nums[left] <= nums[mid]
	}

	handleLeftPortion := func(mid int) {
		targetInLeftPortion := nums[left] <= target && target < nums[mid]
		if targetInLeftPortion {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	handleRightPortion := func(mid int) {
		targetInRightPortion := nums[mid] < target && target <= nums[right]
		if targetInRightPortion {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	for left <= right {
		mid := (left + right) / divisor

		if target == nums[mid] {
			return mid
		}

		if isLeftPortionSorted(mid) {
			handleLeftPortion(mid)
		} else {
			handleRightPortion(mid)
		}
	}

	return -1
}
