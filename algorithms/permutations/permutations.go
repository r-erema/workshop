package permutations

import (
	"slices"
)

// Permute generates all permutations of an array.
// Time O(n * n!), the total number of permutations generated is n!,
// since we're generating all possible permutations of nums,
// additionally, creating the newSet is O(n) per permutation (due to the copy operation)
//
// Space O(n * n!), the output contains n! permutations, each of length n.
// Thus, the output itself occupies O(n * n!) space.
func Permute(nums []int) [][]int {
	res := [][]int{{}}

	var set []int

	for len(res[0]) < len(nums) {
		set, res = res[0], res[1:]

		for _, n := range nums {
			if !Contains(n, set) {
				newSet := make([]int, len(set)+1)
				copy(newSet, set)
				newSet[len(newSet)-1] = n

				res = append(res, newSet)
			}
		}
	}

	return res
}

func Contains(needle int, nums []int) bool {
	return slices.Contains(nums, needle)
}
