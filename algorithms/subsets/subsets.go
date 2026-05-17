package subsets

// Subsets Time O(N * 2^N) it's a brute force way, there are 2^N subsets in total.
// Time O(N), since we allocate memory for each backtracking call and overwrite it on the next backtracking call,
// N is a number of calls.
func Subsets(nums []int) [][]int {
	backtrack := func(num int, subs [][]int) [][]int {
		const multiplier = 2

		newSubs := make([][]int, 0, multiplier*len(subs))

		for _, sub := range subs {
			newSub := make([]int, len(sub))
			copy(newSub, sub)
			sub = append(sub, num)
			newSubs = append(newSubs, newSub, sub)
		}

		return newSubs
	}

	subs := [][]int{{}}
	for i := range nums {
		subs = backtrack(nums[i], subs)
	}

	return subs
}
