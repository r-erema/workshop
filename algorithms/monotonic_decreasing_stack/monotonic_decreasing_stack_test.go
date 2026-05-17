package monotonicdecreasingstack_test_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonotonicDecreasingStack(t *testing.T) {
	t.Parallel()

	monoStack := make([]int, 0)

	for _, number := range []int{3, 1, 6, 2, 5, 4} {
		for len(monoStack) > 0 && monoStack[len(monoStack)-1] < number {
			monoStack = monoStack[:len(monoStack)-1]
		}

		monoStack = append(monoStack, number)
	}

	assert.Equal(t, []int{6, 5, 4}, monoStack)
}
