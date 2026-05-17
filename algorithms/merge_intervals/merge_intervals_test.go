package mergeintervals_test

import (
	"testing"

	mergeintervals "github.com/r-erema/workshop/algorithms/merge_intervals"
	"github.com/stretchr/testify/assert"
)

func TestMergeIntervals(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "One interval",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "Two intervals",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "Three intervals",
			intervals: [][]int{{9, 14}, {1, 2}, {4, 6}, {2, 2}, {8, 9}, {2, 2}, {8, 10}, {11, 15}},
			want:      [][]int{{1, 2}, {4, 6}, {8, 15}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, mergeintervals.MergeIntervals(tt.intervals))
		})
	}
}
