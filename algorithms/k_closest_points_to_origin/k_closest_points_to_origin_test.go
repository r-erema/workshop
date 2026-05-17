package kclosestpointstoorigin_test

import (
	"testing"

	kclosestpointstoorigin "github.com/r-erema/workshop/algorithms/k_closest_points_to_origin"
	"github.com/stretchr/testify/assert"
)

func TestKClosest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		points [][]int
		k      int
		want   [][]int
	}{
		{
			name:   "3 points",
			points: [][]int{{0, 2}, {2, 0}, {2, 2}},
			k:      2,
			want:   [][]int{{0, 2}, {2, 0}},
		},
		{
			name:   "3 points",
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:      2,
			want:   [][]int{{3, 3}, {-2, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.ElementsMatch(t, tt.want, kclosestpointstoorigin.KClosest(tt.points, tt.k))
		})
	}
}
