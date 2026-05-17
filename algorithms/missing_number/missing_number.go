package missingnumber

// MissingNumber finds the missing number in an array.
// Time O(N), since we iterate input one time
// Space O(1), since we don't involve any additional data structure.
func MissingNumber(nums []int) int {
	res := len(nums)

	for i := range nums {
		res += i - nums[i]
	}

	return res
}

// MissingNumber2 finds the missing number using XOR.
// Time O(N), since we iterate input one time
// Space O(1), since we don't involve any additional data structure.
func MissingNumber2(nums []int) int {
	var res int

	for i := range nums {
		res ^= (i + 1) ^ nums[i]
	}

	return res
}
