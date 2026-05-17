package mergeintervals

import (
	"sort"
)

// MergeIntervals merges overlapping intervals.
// Time O(NlogN)
// Other than the sort invocation,
// we do a simple linear scan of the list,
// so the runtime is dominated by the (NlogN) complexity of sorting
//
// Space O(logN)
// If we can sort intervals in place,
// we do not need more than constant additional space,
// although the sorting itself takes O(logN) space.
func MergeIntervals(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	output := [][]int{intervals[0]}

	for _, interval := range intervals[1:] {
		start, end := interval[0], interval[1]

		lastEnd := output[len(output)-1][1]

		if start <= lastEnd {
			output[len(output)-1][1] = max(lastEnd, end)
		} else {
			output = append(output, []int{start, end})
		}
	}

	return output
}
