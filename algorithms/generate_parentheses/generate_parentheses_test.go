package generateparentheses_test

import (
	"testing"

	generateparentheses "github.com/r-erema/workshop/algorithms/generate_parentheses"
	"github.com/stretchr/testify/assert"
)

func TestGenerateParentheses(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "n=3",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "n=1",
			n:    1,
			want: []string{"()"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, generateparentheses.GenerateParenthesis(tt.n))
		})
	}
}
