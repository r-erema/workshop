package breadthfirstsearch_test

import (
	"testing"

	breadthfirstsearch "github.com/r-erema/workshop/algorithms/breadth_first_search"
	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		edges          [][]int
		vertices, want []int
	}{
		{
			name: "8 nodes graph",
			edges: [][]int{
				0: {1, 2, 3},
				1: {0, 8},
				2: {0, 4},
				3: {0, 6},
				4: {2, 5, 7},
				5: {4},
				6: {3},
				7: {4},
				8: {1},
			},
			vertices: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
			want:     []int{1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "2 unconnected graphs",
			edges: [][]int{
				0: {1, 2},
				1: {0, 2},
				2: {0, 1},
				3: {4, 5},
				4: {3, 5},
				5: {3, 4},
			},
			vertices: []int{0, 0, 0, 0, 0, 0},
			want:     []int{1, 1, 1, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			breadthfirstsearch.Bfs(0, tt.edges, tt.vertices)
			assert.Equal(t, tt.want, tt.vertices)
		})
	}
}
