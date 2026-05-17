package isvalidsubsequence_test

import (
	"testing"

	isvalidsubsequence "github.com/r-erema/workshop/algorithms/is_valid_subsequence"
	"github.com/stretchr/testify/assert"
)

func Test_validateSubsequence(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		sequence    []int
		subSequence []int
		isValid     bool
	}{
		{
			name:        "Test case 0",
			sequence:    []int{5, 1, 22, 25, 6, -1, 8, 10},
			subSequence: []int{1, 6, -1, 10},
			isValid:     true,
		},
		{
			name:        "Test case 1",
			sequence:    []int{5, 1, 22, 25, 6, -1, 8, 10},
			subSequence: []int{5, 6, -1, 10},
			isValid:     true,
		},
		{
			name:        "Test case 2",
			sequence:    []int{5, 1, 22, 25, 6, -1, 8, 10},
			subSequence: []int{5, 6, -1, 22},
			isValid:     false,
		},
		{
			name:        "Test case 3",
			sequence:    []int{5, 1, 22, 25, 6, -1, 8, 10},
			subSequence: []int{4},
			isValid:     false,
		},
		{
			name:        "Test case 4",
			sequence:    []int{5, 1, 22, 25, 6, -1, 8, 10},
			subSequence: []int{},
			isValid:     true,
		},
		{
			name:        "Test case 5",
			sequence:    []int{5, 1, 22, 25, 6, -1, 8, 10},
			subSequence: []int{5, 1, 22, 25, 6, -1, 8, 10, 9},
			isValid:     false,
		},
		{
			name:        "Test case 6",
			sequence:    []int{5, 1, 22, 25, 6, -1, 8, 10},
			subSequence: []int{3, 5, 1, 22, 25, 6, -1, 8, 10},
			isValid:     false,
		},
		{
			name:        "Test case 7",
			sequence:    []int{1, 4, 8, 8, 9, 9, 6, 3, 6, 0, 0, 6, 15, 7, 15},
			subSequence: []int{4, 8, 6, 15},
			isValid:     true,
		},
		{
			name:        "Test case 8",
			sequence:    []int{1, 4, 8, 8, 9, 9, 6, 3, 6, 0, 0, 15, 7, 15},
			subSequence: []int{4, 8, 0, 6, 15},
			isValid:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := isvalidsubsequence.IsValidSubsequence(tt.subSequence, tt.sequence)
			assert.Equal(t, tt.isValid, result)
			result = isvalidsubsequence.IsValidSubsequence2(tt.subSequence, tt.sequence)
			assert.Equal(t, tt.isValid, result)
		})
	}
}
