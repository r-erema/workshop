package threesum

import (
	"sort"
)

// ThreeSum Time O(N^2), since we need to iterate an input for each number
// Time O(1) or O(n) depends on sorting algorithm.
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int

	for i := range len(nums) - 2 {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		res = ProcessThreeSumIteration(nums, i, res)
	}

	return res
}

func ProcessThreeSumIteration(nums []int, i int, res [][]int) [][]int {
	left, right := i+1, len(nums)-1

	for left < right {
		sum := nums[i] + nums[left] + nums[right]

		switch {
		case sum > 0:
			right--
		case sum < 0:
			left++
		default:
			res = append(res, []int{nums[i], nums[left], nums[right]})
			left++
			right--

			for left < right && nums[left] == nums[left-1] {
				left++
			}

			for left < right && nums[right] == nums[right+1] {
				right--
			}
		}
	}

	return res
}
