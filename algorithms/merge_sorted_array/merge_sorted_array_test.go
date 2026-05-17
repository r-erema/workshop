package mergesortedarray_test

import (
	"testing"

	mergesortedarray "github.com/r-erema/workshop/algorithms/merge_sorted_array"
	"github.com/stretchr/testify/assert"
)

func TestMergeSortedArray(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		nums1, nums2 []int
		m, n         int
		want         []int
	}{
		{
			name:  "2 normal arrays",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:  "second array is empty",
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			name:  "first array is empty",
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mergesortedarray.Merge(tt.nums1, tt.m, tt.nums2, tt.n)
			assert.Equal(t, tt.want, tt.nums1)
		})
	}
}
