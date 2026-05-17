package longestconsecutivesequence_test

import (
	"testing"

	longestconsecutivesequence "github.com/r-erema/workshop/algorithms/longest_consecutive_sequence"
	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "6 nums",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "10 nums",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, longestconsecutivesequence.LongestConsecutive(tt.nums))
		})
	}
}
