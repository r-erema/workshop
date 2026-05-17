package nodedepth_test

import (
	"testing"

	nodedepth "github.com/r-erema/workshop/algorithms/node_depth"
	"github.com/r-erema/workshop/utils/data_structure/bst"
	"github.com/stretchr/testify/assert"
)

func tree() *bst.BST {
	return bst.NewBST(5).
		InsertRecursively(bst.NewBST(2)).
		InsertRecursively(bst.NewBST(10)).
		InsertRecursively(bst.NewBST(8)).
		InsertRecursively(bst.NewBST(34))
}

func tree2() *bst.BST {
	return bst.NewBST(9).
		InsertRecursively(bst.NewBST(4)).
		InsertRecursively(bst.NewBST(17)).
		InsertRecursively(bst.NewBST(3)).
		InsertRecursively(bst.NewBST(6)).
		InsertRecursively(bst.NewBST(22)).
		InsertRecursively(bst.NewBST(5)).
		InsertRecursively(bst.NewBST(7)).
		InsertRecursively(bst.NewBST(20)).
		InsertRecursively(bst.NewBST(23))
}

func TestFindClosestValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		bst  *bst.BST
		want int
	}{
		{
			name: "Case 0",
			bst:  tree(),
			want: 6,
		},
		{
			name: "Case 1",
			bst:  tree2(),
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			depth := nodedepth.NodeDepthRecursively(tt.bst)
			assert.Equal(t, tt.want, depth)
			depth = nodedepth.NodeDepthIterative(tt.bst)
			assert.Equal(t, tt.want, depth)
			depth = nodedepth.NodeDepthIterative2(tt.bst)
			assert.Equal(t, tt.want, depth)
		})
	}
}
