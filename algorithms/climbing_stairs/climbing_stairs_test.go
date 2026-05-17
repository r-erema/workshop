package climbingstairs_test

import (
	"testing"

	climbingstairs "github.com/r-erema/workshop/algorithms/climbing_stairs"
	"github.com/stretchr/testify/assert"
)

func TestClimbingStairs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		n, want int
	}{
		{
			name: "2 steps",
			n:    2,
			want: 2,
		},
		{
			name: "3 steps",
			n:    3,
			want: 3,
		},
		{
			name: "4 steps",
			n:    4,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, climbingstairs.ClimbStairs(tt.n))
		})
	}
}
