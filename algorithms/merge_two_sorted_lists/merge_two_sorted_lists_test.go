package mergetwosortedlists_test

import (
	"testing"

	mergetwosortedlists "github.com/r-erema/workshop/algorithms/merge_two_sorted_lists"
	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-two-sorted-lists/
func TestMerge2SortedLists(t *testing.T) {
	t.Parallel()

	tests := []mergetwosortedlists.TestCase{
		mergetwosortedlists.NormalLists(),
		mergetwosortedlists.ListIsNil(),
		mergetwosortedlists.ListHasNegativeNumber(),
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.Want, mergetwosortedlists.MergeTwoSortedLists(tt.List1, tt.List2))
		})
	}
}
