package longestconsecutivesequence

func LongestConsecutive(nums []int) int {
	numsMap := make(map[int]struct{})
	for i := range nums {
		numsMap[nums[i]] = struct{}{}
	}

	var result int

	for num := range numsMap {
		if _, ok := numsMap[num-1]; !ok {
			i := 1

			for {
				if _, ok := numsMap[num+i]; !ok {
					break
				}

				i++
			}

			result = max(i, result)
		}
	}

	return result
}
