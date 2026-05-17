package productofarrayexceptself

import (
	"slices"
)

// ProductExceptSelf returns product of all elements except self for each position.
// Time O(N), we iterate linearly through the result array
// Space O(N), we consume only result array which equals input.
func ProductExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	prefix := 1

	for i := range res {
		res[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := range slices.Backward(res) {
		res[i] *= postfix
		postfix *= nums[i]
	}

	return res
}
