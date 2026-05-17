package climbingstairs

// ClimbStairs calculates the number of ways to climb stairs.
// Time O(n), since we walk trough the n 1 time
// Space O(1), we don't use any extra space.
func ClimbStairs(n int) int {
	one, two := 1, 1
	for range n - 1 {
		one, two = one+two, one
	}

	return one
}
