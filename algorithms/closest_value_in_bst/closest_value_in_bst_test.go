package closestvalueinbst_test

import (
	"testing"

	closestvalueinbst "github.com/r-erema/workshop/algorithms/closest_value_in_bst"
	"github.com/stretchr/testify/assert"
)

func TestFindClosestValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		bst    *closestvalueinbst.TreeNode
		target float64
		want   float64
	}{
		{
			name: "case 0",
			bst: &closestvalueinbst.TreeNode{
				Val:  5,
				Left: &closestvalueinbst.TreeNode{Val: 2},
				Right: &closestvalueinbst.TreeNode{
					Val:   10,
					Left:  &closestvalueinbst.TreeNode{Val: 8},
					Right: &closestvalueinbst.TreeNode{Val: 34},
				},
			},
			target: 20,
			want:   10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			closest := closestvalueinbst.ClosestValueInBST(tt.bst, tt.target)
			assert.InEpsilon(t, tt.want, closest, 0)
		})
	}
}
