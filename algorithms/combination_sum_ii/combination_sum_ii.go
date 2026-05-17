package combinationsumii

import (
	"slices"
	"sort"
)

func CombinationSum2(candidates []int, target int) [][]int {
	var (
		result  [][]int
		currSet []int
		total   int
		dfs     func(startIndex int)
	)

	sort.Ints(candidates)

	dfs = func(startIndex int) {
		for i := startIndex; i < len(candidates); i++ {
			currSet = append(currSet, candidates[i])
			total += candidates[i]

			if total < target {
				dfs(i + 1)
			}

			if total == target {
				result = append(result, slices.Clone(currSet))
			}

			total -= currSet[len(currSet)-1]
			currSet = currSet[:len(currSet)-1]

			for i+1 < len(candidates) && candidates[i+1] == candidates[i] {
				i++
			}
		}
	}

	dfs(0)

	return result
}
