package meetingrooms

import (
	"sort"
)

// CanAttendMeetings checks if a person can attend all meetings without conflicts.
// Time O(NlogN), other than the sort invocation, we do a simple linear scan of the list,
// Space O(logN), due to the sorting.
func CanAttendMeetings(meetings [][2]int) bool {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for i := 1; i < len(meetings); i++ {
		if meetings[i][0] <= meetings[i-1][1] {
			return false
		}
	}

	return true
}
