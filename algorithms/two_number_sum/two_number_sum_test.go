package twonumbersum_test

import (
	"testing"

	twonumbersum "github.com/r-erema/workshop/algorithms/two_number_sum"
	"github.com/stretchr/testify/assert"
)

func Test_sum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		numbers        []int
		targetSum      int
		expectedResult [2]int
	}{
		{
			name:           "Test case 0",
			numbers:        []int{3, 5, -4, 8, 11, -1, 6},
			targetSum:      10,
			expectedResult: [2]int{11, -1},
		},
		{
			name: "Test case 1",
			numbers: []int{
				-305,
				3,
				5,
				-4,
				8,
				312,
				11,
				-1,
				6,
				7,
				11,
				-4,
				5,
				8,
				0,
				0,
				0,
				9,
				-2,
				2,
			},
			targetSum:      7,
			expectedResult: [2]int{-305, 312},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := twonumbersum.SumLinear(tt.numbers, tt.targetSum)
			assert.ElementsMatch(t, tt.expectedResult, result)
			result = twonumbersum.SumHashTable(tt.numbers, tt.targetSum)
			assert.ElementsMatch(t, tt.expectedResult, result)
			result = twonumbersum.SumShiftingPointer(tt.numbers, tt.targetSum)
			assert.ElementsMatch(t, tt.expectedResult, result)
		})
	}
}
