package containsduplicate

// ContainsDuplicate checks if the array contains duplicate elements.
// Time O(N), since we iterate input one time
// Space O(N), since we involve map.
func ContainsDuplicate(nums []int) bool {
	visitedNumbers := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := visitedNumbers[num]; ok {
			return true
		}

		visitedNumbers[num] = struct{}{}
	}

	return false
}
