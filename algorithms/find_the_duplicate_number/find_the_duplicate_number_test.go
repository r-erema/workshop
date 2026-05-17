package findtheduplicatenumber_test

import (
	"testing"

	findtheduplicatenumber "github.com/r-erema/workshop/algorithms/find_the_duplicate_number"
	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "2 duplicates",
			nums: []int{1, 3, 4, 5, 2, 2},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, findtheduplicatenumber.FindDuplicate(tt.nums))
		})
	}
}
