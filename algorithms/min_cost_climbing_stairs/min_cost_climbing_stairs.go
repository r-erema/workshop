package mincostclimbingstairs

// MinCostClimbingStairs finds the minimum cost to reach the top of the stairs.
// Time O(N), since we iterate input one time
// Space O(1), since we don't involve any additional data structure.
func MinCostClimbingStairs(cost []int) int {
	const startOffset = 3

	for i := len(cost) - startOffset; i >= 0; i-- {
		cost[i] += min(cost[i+1], cost[i+2])
	}

	return min(cost[0], cost[1])
}
