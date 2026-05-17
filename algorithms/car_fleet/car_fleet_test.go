package carfleet_test

import (
	"testing"

	carfleet "github.com/r-erema/workshop/algorithms/car_fleet"
	"github.com/stretchr/testify/assert"
)

func TestCarFleet(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		target   int
		position []int
		speed    []int
		want     int
	}{
		{
			name:     "3 fleets of 5 cars",
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			want:     3,
		},
		{
			name:     "1 fleets of 1 car",
			target:   10,
			position: []int{3},
			speed:    []int{3},
			want:     1,
		},
		{
			name:     "1 fleets of 3 cars",
			target:   10,
			position: []int{0, 4, 2},
			speed:    []int{2, 1, 3},
			want:     1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, carfleet.CarFleet(tt.target, tt.position, tt.speed))
		})
	}
}
