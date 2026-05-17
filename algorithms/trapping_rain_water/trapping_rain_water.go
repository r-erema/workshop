package trappingrainwater

// Trap Time O(N), since we iterate input one time
// Space O(1), since we don't involve any additional data structure.
func Trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := height[left], height[right]
	res := 0

	for left < right {
		if leftMax < rightMax {
			left++
			leftMax = max(leftMax, height[left])
			res += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			res += rightMax - height[right]
		}
	}

	return res
}
