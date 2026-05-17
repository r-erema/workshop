package topkfrequentelements

// TopKFrequent Time O(n) since we have single loops
// Time O(N+M) since we have a map N containing counts of each number and slice M grouped numbers by count.
func TopKFrequent(nums []int, frequency int) []int {
	numsCountMap := make(map[int]int, 0)

	for _, num := range nums {
		numsCountMap[num]++
	}

	countsArr := make([][]int, len(nums)+1)

	for num, count := range numsCountMap {
		countsArr[count] = append(countsArr[count], num)
	}

	var result []int
	for i := len(countsArr) - 1; i > 0; i-- {
		result = append(result, countsArr[i]...)

		if len(result) >= frequency {
			return result[:frequency]
		}
	}

	return result
}
