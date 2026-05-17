package threesum_test

import (
	"testing"

	threesum "github.com/r-erema/workshop/algorithms/three_sum"
	"github.com/stretchr/testify/assert"
)

func Test3Sum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Two sets",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, threesum.ThreeSum(tt.nums))
		})
	}
}
