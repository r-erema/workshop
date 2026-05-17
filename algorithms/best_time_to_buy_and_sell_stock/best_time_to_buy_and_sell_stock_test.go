package besttimetobuyandsellstock_test

import (
	"testing"

	besttimetobuyandsellstock "github.com/r-erema/workshop/algorithms/best_time_to_buy_and_sell_stock"
	"github.com/stretchr/testify/assert"
)

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "Simple set",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, besttimetobuyandsellstock.MaxProfit(tt.prices))
		})
	}
}
