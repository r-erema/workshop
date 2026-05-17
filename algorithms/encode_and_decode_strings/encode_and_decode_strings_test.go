package encodeanddecodestrings_test

import (
	"testing"

	encodeanddecodestrings "github.com/r-erema/workshop/algorithms/encode_and_decode_strings"
	"github.com/stretchr/testify/assert"
)

func TestEncodeAndDecodeStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		source     []string
		encodedStr string
	}{
		{
			name:       "Simple words",
			source:     []string{"lint", "code", "love", "you"},
			encodedStr: "4#lint4#code4#love3#you",
		},
		{
			name:       "Long words",
			source:     []string{"estimation#", "#highlighted", "br#ing_it_to_us"},
			encodedStr: "11#estimation#12##highlighted15#br#ing_it_to_us",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			encoded := encodeanddecodestrings.Encode(tt.source)
			assert.Equal(t, tt.encodedStr, encoded)
			decoded := encodeanddecodestrings.Decode(encoded)
			assert.Equal(t, tt.source, decoded)
		})
	}
}
