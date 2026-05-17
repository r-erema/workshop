package dailytemperatures_test

import (
	"testing"

	dailytemperatures "github.com/r-erema/workshop/algorithms/daily_temperatures"
	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{
			name:         "Random temperatures",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, dailytemperatures.DailyTemperatures(tt.temperatures))
		})
	}
}
