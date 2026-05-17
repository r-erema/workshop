package kthlargestelementinanarray_test

import (
	"testing"

	kthlargestelementinanarray "github.com/r-erema/workshop/algorithms/kth_largest_element_in_an_array"
	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "ordinary input",
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		{
			name: "input with duplicates",
			nums: []int{1, 2, 1, 1, 1, -1, -2},
			k:    1,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, kthlargestelementinanarray.FindKthLargest(tt.nums, tt.k))
		})
	}
}
