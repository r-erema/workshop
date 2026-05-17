package combinationsum

func CombinationSum(candidates []int, target int) [][]int {
	var (
		res [][]int
		set []int
	)

	var backtrack func(start, total int)

	backtrack = func(start, total int) {
		if total > target {
			return
		}

		if total == target {
			newSet := make([]int, len(set))
			copy(newSet, set)
			res = append(res, newSet)

			return
		}

		for i := start; i < len(candidates); i++ {
			set = append(set, candidates[i])
			backtrack(i, total+candidates[i])

			set = set[:len(set)-1]
		}
	}

	backtrack(0, 0)

	return res
}
