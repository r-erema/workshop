package caesarcipherencrypt_test

import (
	"testing"

	caesarcipherencrypt "github.com/r-erema/workshop/algorithms/caesar_cipher_encrypt"
	"github.com/stretchr/testify/assert"
)

func TestCaesarCipherEncrypt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name, str, want string
		shift           int32
	}{
		{
			name:  "Case 0",
			str:   "xyz",
			shift: 2,
			want:  "zab",
		},
		{
			name:  "Case 1",
			str:   "lorem ipsum dolor sit amet, consectetur adipiscing elit",
			shift: 732,
			want:  "psviq mtwyq hspsv wmx eqix, gsrwigxixyv ehmtmwgmrk ipmx",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, caesarcipherencrypt.CaesarCipherEncrypt(tt.str, tt.shift))
		})
	}
}
