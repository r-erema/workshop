package countgoodnodesinbinarytree_test

import (
	"testing"

	countgoodnodesinbinarytree "github.com/r-erema/workshop/algorithms/count_good_nodes_in_binary_tree"
	"github.com/r-erema/workshop/utils/data_structure/tree"
	"github.com/stretchr/testify/assert"
)

func TestGoodNodes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		root *tree.Node
		want int
	}{
		{
			name: "4 good nodes",
			root: &tree.Node{
				Val: 3,
				Left: &tree.Node{
					Val:  1,
					Left: &tree.Node{Val: 3},
				},
				Right: &tree.Node{
					Val:   4,
					Left:  &tree.Node{Val: 1},
					Right: &tree.Node{Val: 5},
				},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, countgoodnodesinbinarytree.GoodNodes(tt.root))
		})
	}
}
