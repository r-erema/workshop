package clone_graph_test

import (
	"testing"

	"github.com/r-erema/workshop/algorithms/clone_graph"
	"github.com/r-erema/workshop/utils/data_structure/graph"
	"github.com/stretchr/testify/assert"
)

func TestCloneGraph(t *testing.T) {
	t.Parallel()

	node1 := &graph.Node{Val: 1}
	node2 := &graph.Node{Val: 2}
	node3 := &graph.Node{Val: 3}
	node4 := &graph.Node{Val: 4}

	node1.Neighbors = []*graph.Node{node2, node4}
	node2.Neighbors = []*graph.Node{node1, node3}
	node3.Neighbors = []*graph.Node{node2, node4}
	node4.Neighbors = []*graph.Node{node1, node3}

	tests := []struct {
		name  string
		graph *graph.Node
		want  *graph.Node
	}{
		{
			name:  "connected graph",
			graph: node1,
			want:  node1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cg := clone_graph.CloneGraph(tt.graph)

			assert.Equal(t, tt.want, cg)
		})
	}
}
