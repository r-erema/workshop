package subsetsii

import (
	"slices"
	"sort"
)

func SubsetsWithDup(nums []int) [][]int {
	res := [][]int{{}}

	var (
		curr []int
		dfs  func(int)
	)

	sort.Ints(nums)

	dfs = func(i int) {
		for ; i < len(nums); i++ {
			curr = append(curr, nums[i])

			res = append(res, slices.Clone(curr))

			dfs(i + 1)

			curr = curr[:len(curr)-1]

			for i+1 < len(nums) && nums[i] == nums[i+1] {
				i++
			}
		}
	}

	dfs(0)

	return res
}
