package kthsmallestelementinabst_test

import (
	"testing"

	kthsmallestelementinabst "github.com/r-erema/workshop/algorithms/kth_smallest_element_in_a_bst"
	"github.com/r-erema/workshop/utils/data_structure/tree"
	"github.com/stretchr/testify/assert"
)

func TestKthSmallest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		root    *tree.Node
		k, want int
	}{
		{
			name: "Tree 1",
			root: &tree.Node{
				Val: 3,
				Left: &tree.Node{
					Val: 1,
					Right: &tree.Node{
						Val: 2,
					},
				},
				Right: &tree.Node{
					Val: 4,
				},
			},
			k:    1,
			want: 1,
		},
		{
			name: "Tree 2",
			root: &tree.Node{
				Val: 5,
				Left: &tree.Node{
					Val: 3,
					Left: &tree.Node{
						Val:  2,
						Left: &tree.Node{Val: 1},
					},
					Right: &tree.Node{
						Val: 4,
					},
				},
				Right: &tree.Node{
					Val: 6,
				},
			},

			k:    3,
			want: 3,
		},
		{
			name: "Tree 3",
			root: &tree.Node{
				Val: 3,
				Left: &tree.Node{
					Val: 1,
					Right: &tree.Node{
						Val: 2,
					},
				},
				Right: &tree.Node{
					Val: 4,
				},
			},

			k:    2,
			want: 2,
		},
		{
			name: "Tree 4",
			root: &tree.Node{
				Val: 5,
				Left: &tree.Node{
					Val: 3,
					Left: &tree.Node{
						Val: 1,
						Right: &tree.Node{
							Val: 2,
						},
					},
					Right: &tree.Node{
						Val: 4,
					},
				},
				Right: &tree.Node{
					Val: 6,
				},
			},

			k:    2,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, kthsmallestelementinabst.KthSmallest(tt.root, tt.k))
		})
	}
}
