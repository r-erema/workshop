package productofarrayexceptself_test

import (
	"testing"

	productofarrayexceptself "github.com/r-erema/workshop/algorithms/product_of_array_except_self"
	"github.com/stretchr/testify/assert"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Simple array",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "Array with negative numbers",
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, productofarrayexceptself.ProductExceptSelf(tt.nums))
		})
	}
}
