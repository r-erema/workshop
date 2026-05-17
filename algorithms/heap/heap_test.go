package heap_test

import (
	"testing"

	"github.com/r-erema/workshop/algorithms/heap"
	"github.com/r-erema/workshop/utils/data_structure/tree"
	"github.com/stretchr/testify/assert"
)

func TestHeapToArrayAndArrayToHeap(t *testing.T) {
	t.Parallel()

	arr := heap.TreeToArray(heap.TestHeap())
	assert.Equal(t, []int{14, 19, 16, 21, 26, 19, 68, 65, 30}, arr)

	h := heap.ArrayToTree(arr)
	assert.Equal(t, heap.TestHeap(), h)
}

func TestPushToHeap(t *testing.T) {
	t.Parallel()

	expectedHeap := &tree.Node{
		Val: 14,
		Left: &tree.Node{
			Val: 17,
			Left: &tree.Node{
				Val: 21,
				Left: &tree.Node{
					Val: 65,
				},
				Right: &tree.Node{
					Val: 30,
				},
			},
			Right: &tree.Node{
				Val: 19,
				Left: &tree.Node{
					Val: 26,
				},
			},
		},
		Right: &tree.Node{
			Val: 16,
			Left: &tree.Node{
				Val: 19,
			},
			Right: &tree.Node{
				Val: 68,
			},
		},
	}

	h := heap.TestHeap()
	heap.Push(h, 17)
	assert.Equal(t, expectedHeap, h)
}

func TestPopHeap(t *testing.T) {
	t.Parallel()

	expectedHeap := &tree.Node{
		Val: 16,
		Left: &tree.Node{
			Val: 19,
			Left: &tree.Node{
				Val: 21,
				Left: &tree.Node{
					Val: 65,
				},
			},
			Right: &tree.Node{
				Val: 26,
			},
		},
		Right: &tree.Node{
			Val: 19,
			Left: &tree.Node{
				Val: 30,
			},
			Right: &tree.Node{
				Val: 68,
			},
		},
	}

	clonedTestHeap := *heap.TestHeap()
	popped := heap.Pop(&clonedTestHeap)
	assert.Equal(t, 14, popped)
	assert.Equal(t, expectedHeap, &clonedTestHeap)
}

func TestHeapify(t *testing.T) {
	t.Parallel()

	sourceHeap := &tree.Node{
		Val: 50,
		Left: &tree.Node{
			Val: 80,
			Left: &tree.Node{
				Val: 30,
				Left: &tree.Node{
					Val: 90,
				},
				Right: &tree.Node{
					Val: 60,
				},
			},
			Right: &tree.Node{
				Val: 10,
			},
		},
		Right: &tree.Node{
			Val: 40,
			Left: &tree.Node{
				Val: 70,
			},
			Right: &tree.Node{
				Val: 20,
			},
		},
	}

	expectedHeap := &tree.Node{
		Val: 10,
		Left: &tree.Node{
			Val: 30,
			Left: &tree.Node{
				Val: 50,
				Left: &tree.Node{
					Val: 90,
				},
				Right: &tree.Node{
					Val: 60,
				},
			},
			Right: &tree.Node{
				Val: 80,
			},
		},
		Right: &tree.Node{
			Val: 20,
			Left: &tree.Node{
				Val: 70,
			},
			Right: &tree.Node{
				Val: 40,
			},
		},
	}

	heap.Heapify(sourceHeap)
	assert.Equal(t, expectedHeap, sourceHeap)
}
