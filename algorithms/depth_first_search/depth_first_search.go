package depthfirstsearch_test_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func DFS() {

}

func TestDFS(t *testing.T) {
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

			var dfsRecursive func(nodeIndex int)

			dfsRecursiveVertices := make([]int, len(tt.vertices))
			copy(dfsRecursiveVertices, tt.vertices)

			dfsRecursive = func(nodeIndex int) {
				dfsRecursiveVertices[nodeIndex] = 1
				for _, adjIndex := range tt.edges[nodeIndex] {
					if dfsRecursiveVertices[adjIndex] != 1 {
						dfsRecursive(adjIndex)
					}
				}
			}

			dfsStackVertices := make([]int, len(tt.vertices))
			copy(dfsStackVertices, tt.vertices)

			dfsStack := func(nodeIndex int) {
				stack := []int{nodeIndex}

				for len(stack) > 0 {
					nodeIndex, stack = stack[len(stack)-1], stack[:len(stack)-1]
					dfsStackVertices[nodeIndex] = 1

					for _, edge := range tt.edges[nodeIndex] {
						if dfsStackVertices[edge] == 0 {
							stack = append(stack, edge)
						}
					}
				}
			}

			dfsRecursive(0)
			assert.Equal(t, tt.want, dfsRecursiveVertices)

			dfsStack(0)
			assert.Equal(t, tt.want, dfsStackVertices)
		})
	}
}
