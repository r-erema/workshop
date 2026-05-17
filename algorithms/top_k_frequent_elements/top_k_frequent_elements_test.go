package topkfrequentelements_test

import (
	"testing"

	topkfrequentelements "github.com/r-erema/workshop/algorithms/top_k_frequent_elements"
	"github.com/stretchr/testify/assert"
)

func TestTopKFrequentElements(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "Normal array",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "One element array",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, topkfrequentelements.TopKFrequent(tt.nums, tt.k))
		})
	}
}
