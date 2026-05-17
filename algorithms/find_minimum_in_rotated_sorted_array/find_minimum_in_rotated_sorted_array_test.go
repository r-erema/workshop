package findminimuminrotatedsortedarray_test

import (
	"testing"

	findminimuminrotatedsortedarray "github.com/r-erema/workshop/algorithms/find_minimum_in_rotated_sorted_array"
	"github.com/stretchr/testify/assert"
)

func TestFindMin(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "middle element in the array is searched",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "middle element in the longer array is searched",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "first element in the array is searched",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, findminimuminrotatedsortedarray.FindMin(tt.nums))
		})
	}
}
