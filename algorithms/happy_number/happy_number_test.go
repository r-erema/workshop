package happynumber_test

import (
	"testing"

	happynumber "github.com/r-erema/workshop/algorithms/happy_number"
	"github.com/stretchr/testify/assert"
)

func TestHappyNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "number 19",
			input: 19,
			want:  true,
		},
		{
			name:  "number 2",
			input: 2,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, happynumber.IsHappy(tt.input))
		})
	}
}
