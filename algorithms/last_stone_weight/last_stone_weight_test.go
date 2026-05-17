package laststoneweight_test

import (
	"testing"

	laststoneweight "github.com/r-erema/workshop/algorithms/last_stone_weight"
	"github.com/stretchr/testify/assert"
)

func TestLastStoneWeight(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{
			name:   "6 stones",
			stones: []int{2, 7, 4, 1, 8, 1},
			want:   1,
		},
		{
			name:   "2 stones ascending ordering",
			stones: []int{1, 3},
			want:   2,
		},
		{
			name:   "2 stones descending ordering",
			stones: []int{3, 1},
			want:   2,
		},
		{
			name:   "4 stones",
			stones: []int{9, 3, 2, 10},
			want:   0,
		},
		{
			name:   "8 stones",
			stones: []int{10, 5, 4, 10, 3, 1, 7, 8},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, laststoneweight.LastStoneWeight(tt.stones))
		})
	}
}
