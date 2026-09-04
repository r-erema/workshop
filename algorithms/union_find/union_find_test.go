package union_find_test

import (
	"testing"

	"github.com/r-erema/workshop/algorithms/union_find"
	assert "github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		edges          []int
		expectedCycles int
	}{
		{
			name:           "2 cycles",
			edges:          []int{0, 0, 2, 2, 0},
			expectedCycles: 2,
		},
		{
			name:           "2 cycles, input started not from 0",
			edges:          []int{1, 2, 0},
			expectedCycles: 1,
		},
		{
			name:           "2 cycles of single nodes",
			edges:          []int{0, 1},
			expectedCycles: 2,
		},
		{
			name:           "3 cycles",
			edges:          []int{0, 0, 1, 3, 3, 4, 6, 6, 7},
			expectedCycles: 3,
		},
		{
			name:           "3 cycles, input started not from 0",
			edges:          []int{3, 4, 2, 2, 4, 0, 6, 6},
			expectedCycles: 3,
		},
		{
			"4 cycles",
			[]int{0, 0, 0, 0, 4, 1, 4, 8, 8, 10, 0, 8, 13, 13},
			4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expectedCycles, union_find.UnionFind(tt.edges))
		})
	}
}

/*func TestUnionFind(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                string
		edges               [][2]int
		expectedGraphsCount int
	}{
		{
			name:                "1 graph",
			edges:               [][2]int{{0, 1}, {0, 2}},
			expectedGraphsCount: 1,
		},
		{
			name:                "1 graph",
			edges:               [][2]int{{0, 1}, {2, 3}, {1, 2}},
			expectedGraphsCount: 1,
		},
		{
			name: "2 graphs",
			edges: [][2]int{
				{0, 1},
				{0, 2},
				{3, 4},
				{4, 5},
				{6, 3},
			},
			expectedGraphsCount: 2,
		},
		{
			name: "4 graphs",
			edges: [][2]int{
				{0, 1},
				{1, 2},
				{3, 4},
				{4, 5},
				{3, 6},
				{6, 7},
				{4, 7},
				{8, 8},
				{9, 10},
				{9, 11},
				{10, 11},
				{10, 12},
				{12, 13},
				{10, 13},
			},
			expectedGraphsCount: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.expectedGraphsCount, UnionFind(tt.edges))
		})
	}
}*/
