package productsum_test

import (
	"testing"

	productsum "github.com/r-erema/workshop/algorithms/product_sum"
	"github.com/stretchr/testify/assert"
)

func TestProductSum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		array []any
		want  int
	}{
		{
			name: "Case 0",
			array: []any{
				5,
				2,
				[]any{7, -1},
				3,
				[]any{
					6,
					[]any{-13, 8},
					4,
				},
			},
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, productsum.ProductSum(tt.array))
		})
	}
}
