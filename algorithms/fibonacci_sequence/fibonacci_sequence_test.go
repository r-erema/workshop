package fibonaccisequence_test

import (
	"testing"

	fibonaccisequence "github.com/r-erema/workshop/algorithms/fibonacci_sequence"
	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		n, want int
	}{
		{
			name: "Case 0",
			n:    3,
			want: 2,
		},
		{
			name: "Case 1",
			n:    10,
			want: 55,
		},
		{
			name: "Case 2",
			n:    0,
			want: 0,
		},
		{
			name: "Case 3",
			n:    30,
			want: 832040,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, fibonaccisequence.FibonacciRecursive(tt.n))
			assert.Equal(t, tt.want, fibonaccisequence.FibonacciCache(tt.n))
			assert.Equal(t, tt.want, fibonaccisequence.FibonacciIterative(tt.n))
		})
	}
}
