package kokoeatingbananas

import (
	"math"
	"slices"
)

// MinEatingSpeed finds the minimum eating speed for Koko to eat all bananas in given hours.
// Time O(NlogM), other than the sort invocation, we do a simple linear scan of the list,
// Space O(1), where N is the length of the input array piles and M is the maximum number of bananas in a pile.
func MinEatingSpeed(piles []int, hours int) int {
	minSpeed := slices.Max(piles)
	left, right := 1, minSpeed

	eatingTime := func(speed int) int {
		var time int
		for _, pile := range piles {
			time += int(math.Ceil(float64(pile) / float64(speed)))
		}

		return time
	}

	const divisor = 2

	for left <= right {
		potentialMinSpeed := (left + right) / divisor

		if eatingTime(potentialMinSpeed) <= hours {
			minSpeed = potentialMinSpeed
			right = potentialMinSpeed - 1
		} else {
			left = potentialMinSpeed + 1
		}
	}

	return minSpeed
}
