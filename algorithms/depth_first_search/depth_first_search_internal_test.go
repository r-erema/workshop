package depth_first_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		graph            [][]int
		startIndex       int
		wantVisitedOrder map[int]int // map[vertex]order of visiting
	}{
		{
			name: "3 nodes graph",
			graph: [][]int{
				0: {1, 2},
				1: {0, 2},
				2: {0, 1},
			},
			startIndex:       0,
			wantVisitedOrder: map[int]int{0: 1, 1: 2, 2: 3},
		},
		{
			name: "5 nodes graph",
			graph: [][]int{
				0: {4},
				1: {3},
				2: {4},
				3: {1},
				4: {3, 0, 2},
			},
			startIndex:       4,
			wantVisitedOrder: map[int]int{4: 1, 3: 2, 1: 3, 0: 4, 2: 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.wantVisitedOrder, DFSRecursive(tt.graph, tt.startIndex))
			assert.Equal(t, tt.wantVisitedOrder, DFSStack(tt.graph, tt.startIndex))
		})
	}
}
