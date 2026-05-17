package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func bst1() *BST {
	return &BST{
		value: 5,
		left:  &BST{value: 3, left: nil, right: nil},
		right: &BST{value: 7, left: nil, right: nil},
	}
}

func bst2() *BST {
	return &BST{
		value: 4,
		left: &BST{
			value: 2,
			left:  &BST{value: 1, left: nil, right: nil},
			right: &BST{value: 3, left: nil, right: nil},
		},
		right: &BST{
			value: 8,
			left: &BST{
				value: 4,
				right: &BST{value: 6, left: nil, right: nil},
				left:  nil,
			},
			right: &BST{
				value: 11,
				left:  &BST{value: 10, left: nil, right: nil},
				right: nil,
			},
		},
	}
}

func bst3() *BST {
	return &BST{
		value: 1,
		left: &BST{
			value: 2,
			left:  &BST{value: 4, left: nil, right: nil},
			right: &BST{value: 5, left: nil, right: nil},
		},
		right: &BST{value: 3, left: nil, right: nil},
	}
}

func bst4() *BST {
	return &BST{
		value: 1,
		left: &BST{
			value: 2,
			left:  &BST{value: 4, left: nil, right: nil},
			right: &BST{value: 5, left: nil, right: nil},
		},
		right: &BST{
			value: 3,
			left:  nil,
			right: &BST{value: 6, left: nil, right: nil},
		},
	}
}

func bst5() *BST {
	return &BST{
		value: 50,
		left: &BST{
			value: 40,
			left:  nil,
			right: nil,
		},
		right: &BST{
			value: 70,
			left:  &BST{value: 60, left: nil, right: nil},
			right: &BST{value: 80, left: nil, right: nil},
		},
	}
}

func bst6() *BST {
	return &BST{
		value: 50,
		left: &BST{
			value: 30,
			left:  nil,
			right: &BST{value: 40, left: nil, right: nil},
		},
		right: &BST{
			value: 70,
			left:  &BST{value: 60, left: nil, right: nil},
			right: &BST{value: 80, left: nil, right: nil},
		},
	}
}

func bst7() *BST {
	return &BST{
		value: 50,
		left: &BST{
			value: 25,
			left:  &BST{value: 6, left: nil, right: nil},
			right: &BST{value: 30, left: nil, right: nil},
		},
		right: &BST{
			value: 75,
			right: nil,
			left: &BST{
				value: 60,
				left:  &BST{value: 52, left: nil, right: nil},
				right: &BST{value: 70, left: nil, right: nil},
			},
		},
	}
}

func TestBST_Insert(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		initBST *BST
		numbers []float32
		want    *BST
	}{
		{
			name:    "insert case 0",
			initBST: NewBST(5),
			numbers: []float32{3, 7},
			want:    bst1(),
		},
		{
			name:    "insert case 1",
			initBST: NewBST(4),
			numbers: []float32{2, 1, 8, 3, 11, 4, 6, 10},
			want:    bst2(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			bst := *tt.initBST
			for _, number := range tt.numbers {
				bst.InsertRecursively(NewBST(number))
			}

			assert.Equal(t, *tt.want, bst)

			bst = *tt.initBST
			for _, number := range tt.numbers {
				bst.InsertIteratively(NewBST(number))
			}

			assert.Equal(t, *tt.want, bst)
		})
	}
}

func TestBST_Traverse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                                      string
		bst                                       BST
		wantTraversePreorder, wantTraverseInorder []float32
	}{
		{
			name:                 "traverse case 0",
			bst:                  *bst1(),
			wantTraversePreorder: []float32{5, 3, 7},
			wantTraverseInorder:  []float32{3, 5, 7},
		},
		{
			name:                 "traverse case 1",
			bst:                  *bst2(),
			wantTraversePreorder: []float32{4, 2, 1, 3, 8, 4, 6, 11, 10},
			wantTraverseInorder:  []float32{1, 2, 3, 4, 4, 6, 8, 10, 11},
		},
		{
			name:                 "traverse case 2",
			bst:                  *bst3(),
			wantTraversePreorder: []float32{1, 2, 4, 5, 3},
			wantTraverseInorder:  []float32{4, 2, 5, 1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := make([]float32, 0)
			for v := range tt.bst.TraversePreorder() {
				result = append(result, v.value)
			}

			assert.Equal(t, tt.wantTraversePreorder, result)

			result = make([]float32, 0)
			for v := range tt.bst.TraverseInorder() {
				result = append(result, v.value)
			}

			assert.Equal(t, tt.wantTraverseInorder, result)
		})
	}
}

func TestBST_find(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		needle    float32
		bst, want *BST
	}{
		{
			name:   "find case 0",
			bst:    bst2(),
			needle: 6,
			want:   &BST{value: 6, left: nil, right: nil},
		},
		{
			name:   "find case 1",
			bst:    bst2(),
			needle: 11,
			want: &BST{
				value: 11,
				left:  &BST{value: 10, left: nil, right: nil},
				right: nil,
			},
		},
		{
			name:   "find case 2",
			bst:    bst2(),
			needle: 8,
			want: &BST{
				value: 8,
				left: &BST{
					value: 4,
					right: &BST{value: 6, left: nil, right: nil},
					left:  nil,
				},
				right: &BST{
					value: 11,
					left:  &BST{value: 10, left: nil, right: nil},
					right: nil,
				},
			},
		},
		{
			name:   "find case 3",
			bst:    bst2(),
			needle: -1,
			want:   nil,
		},
		{
			name:   "find case 4",
			bst:    bst2(),
			needle: 2,
			want: &BST{
				value: 2,
				left:  &BST{value: 1, left: nil, right: nil},
				right: &BST{value: 3, left: nil, right: nil},
			},
		},
		{
			name:   "find case 5",
			bst:    bst2(),
			needle: 4,
			want:   bst2(),
		},
		{
			name:   "find case 6",
			bst:    bst1(),
			needle: 5,
			want:   bst1(),
		},
		{
			name:   "find case 7",
			bst:    bst1(),
			needle: 3,
			want:   &BST{value: 3, left: nil, right: nil},
		},
		{
			name:   "find case 8",
			bst:    bst1(),
			needle: 7,
			want:   &BST{value: 7, left: nil, right: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			node := tt.bst.find(tt.needle)
			assert.Equal(t, tt.want, node)
		})
	}
}

func TestBST_inorderSuccessor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		bst, want *BST
		needle    float32
	}{
		{
			name:   "inorder successor case 0",
			bst:    bst4(),
			needle: 4,
			want: &BST{
				value: 2,
				left:  &BST{value: 4, left: nil, right: nil},
				right: &BST{value: 5, left: nil, right: nil},
			},
		},
		{
			name:   "inorder successor case 1",
			bst:    bst4(),
			needle: 5,
			want: &BST{
				value: 1,
				left: &BST{
					value: 2,
					left:  &BST{value: 4, left: nil, right: nil},
					right: &BST{value: 5, left: nil, right: nil},
				},
				right: &BST{
					value: 3,
					left:  nil,
					right: &BST{value: 6, left: nil, right: nil},
				},
			},
		},
		{
			name:   "inorder successor case 2",
			bst:    bst4(),
			needle: 6,
			want:   nil,
		},
		{
			name:   "inorder successor case 3",
			bst:    bst4(),
			needle: -1,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, node := tt.bst.inorderSuccessor(tt.needle)
			assert.Equal(t, tt.want, node)
		})
	}
}

func TestBST_Delete(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		bst, want *BST
		toDelete  float32
	}{
		{
			name:     "delete case 0",
			bst:      bst5(),
			toDelete: 50,
			want: &BST{
				value: 60,
				left: &BST{
					value: 40,
					left:  nil,
					right: nil,
				},
				right: &BST{
					value: 70,
					left:  nil,
					right: &BST{value: 80, left: nil, right: nil},
				},
			},
		},
		{
			name:     "delete case 1",
			bst:      bst6(),
			toDelete: 40,
			want: &BST{
				value: 50,
				left: &BST{
					value: 30,
					left:  nil,
					right: nil,
				},
				right: &BST{
					value: 70,
					left:  &BST{value: 60, left: nil, right: nil},
					right: &BST{value: 80, left: nil, right: nil},
				},
			},
		},
		{
			name:     "delete case 2",
			bst:      bst6(),
			toDelete: 30,
			want: &BST{
				value: 50,
				left: &BST{
					value: 40,
					left:  nil,
					right: nil,
				},
				right: &BST{
					value: 70,
					left:  &BST{value: 60, left: nil, right: nil},
					right: &BST{value: 80, left: nil, right: nil},
				},
			},
		},
		{
			name:     "delete case 3",
			bst:      bst7(),
			toDelete: 50,
			want: &BST{
				value: 52,
				left: &BST{
					value: 25,
					left:  &BST{value: 6, left: nil, right: nil},
					right: &BST{value: 30, left: nil, right: nil},
				},
				right: &BST{
					value: 75,
					right: nil,
					left: &BST{
						value: 60,
						left:  nil,
						right: &BST{value: 70, left: nil, right: nil},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.bst.Delete(tt.toDelete)
			assert.Equal(t, tt.want, tt.bst)
		})
	}
}

func TestBST_BranchVectors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		bst  *BST
		want [][]float32
	}{
		{
			name: "branch vectors case 0",
			bst:  bst1(),
			want: [][]float32{
				{5, 3},
				{5, 7},
			},
		},
		{
			name: "branch vectors case 1",
			bst:  bst2(),
			want: [][]float32{
				{4, 2, 1},
				{4, 2, 3},
				{4, 8, 4, 6},
				{4, 8, 11, 10},
			},
		},
		{
			name: "branch vectors case 2",
			bst:  bst3(),
			want: [][]float32{
				{1, 2, 4},
				{1, 2, 5},
				{1, 3},
			},
		},
		{
			name: "branch vectors case 3",
			bst:  bst5(),
			want: [][]float32{
				{50, 40},
				{50, 70, 60},
				{50, 70, 80},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			vectors := tt.bst.BranchVectors()
			assert.Equal(t, tt.want, vectors)
		})
	}
}
