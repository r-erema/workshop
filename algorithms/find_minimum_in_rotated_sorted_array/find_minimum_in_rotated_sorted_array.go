package findminimuminrotatedsortedarray

// FindMin finds the minimum element in a rotated sorted array.
// Time O(log*N), since the binary search
// Space O(1), since we don't involve any additional space.
func FindMin(nums []int) int {
	left, right, res := 0, len(nums)-1, nums[0]

	const divisor = 2

	for left <= right {
		mid := (left + right) / divisor

		if nums[mid] >= res {
			left = mid + 1
		} else {
			right = mid - 1
		}

		res = min(res, nums[mid])
	}

	return res
}
