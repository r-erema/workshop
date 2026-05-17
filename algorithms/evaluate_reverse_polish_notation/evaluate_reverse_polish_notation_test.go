package evaluatereversepolishnotation_test

import (
	"testing"

	evaluatereversepolishnotation "github.com/r-erema/workshop/algorithms/evaluate_reverse_polish_notation"
	"github.com/stretchr/testify/assert"
)

func TestEvaluateReversePolishNotation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "Short tokens sequence",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			name:   "Long tokens sequence",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, evaluatereversepolishnotation.EvaluateReversePolishNotation(tt.tokens))
		})
	}
}
