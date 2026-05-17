package trappingrainwater_test

import (
	"testing"

	trappingrainwater "github.com/r-erema/workshop/algorithms/trapping_rain_water"
	"github.com/stretchr/testify/assert"
)

func TestTrappingRainWater(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "Normal trap",
			heights: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:    6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, trappingrainwater.Trap(tt.heights))
		})
	}
}
