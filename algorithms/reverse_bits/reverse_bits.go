package reversebits

// ReverseBits reverses the bits of a number.
// Time O(N), since we iterate input one time
// Space O(N), since we involve additional allocation to collect reversed bits.
func ReverseBits(num uint32) uint32 {
	var res uint32

	const (
		bits         = 32
		lastBitIndex = bits - 1
	)

	for i := range bits {
		bit := (num >> i) & 1
		res |= bit << (lastBitIndex - i)
	}

	return res
}
