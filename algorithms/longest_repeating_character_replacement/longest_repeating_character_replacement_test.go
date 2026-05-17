package longestrepeatingcharacterreplacement_test

import (
	"testing"

	longestrepeatingcharacterreplacement "github.com/r-erema/workshop/algorithms/longest_repeating_character_replacement"
	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		k     int
		want  int
	}{
		{
			name:  "Equal count of 2 chars",
			input: "ABAB",
			k:     2,
			want:  4,
		},
		{
			name:  "Count of 1 char is bigger then other",
			input: "AABABBA",
			k:     1,
			want:  4,
		},
		{
			name:  "All chars the same",
			input: "AAAA",
			k:     0,
			want:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, longestrepeatingcharacterreplacement.CharacterReplacement(tt.input, tt.k))
		})
	}
}
