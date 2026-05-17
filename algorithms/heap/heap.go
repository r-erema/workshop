package heap

import (
	"github.com/r-erema/workshop/utils/data_structure/tree"
)

func TestHeap() *tree.Node {
	const (
		rootVal       = 14
		leftVal       = 19
		leftLeftVal   = 21
		leftLeftLVal  = 65
		leftLeftRVal  = 30
		leftRightVal  = 26
		rightVal      = 16
		rightLeftVal  = 19
		rightRightVal = 68
	)

	return &tree.Node{
		Val: rootVal,
		Left: &tree.Node{
			Val: leftVal,
			Left: &tree.Node{
				Val: leftLeftVal,
				Left: &tree.Node{
					Val: leftLeftLVal,
				},
				Right: &tree.Node{
					Val: leftLeftRVal,
				},
			},
			Right: &tree.Node{
				Val: leftRightVal,
			},
		},
		Right: &tree.Node{
			Val: rightVal,
			Left: &tree.Node{
				Val: rightLeftVal,
			},
			Right: &tree.Node{
				Val: rightRightVal,
			},
		},
	}
}

// Push adds an element to a heap.
// Time O(log(N)), the number of operations required depends only on the number of levels
// the new element must rise to satisfy the heap property
// Space O(N), since we need an array containing elements from input heap.
func Push(heap *tree.Node, val int) {
	arrHeap := TreeToArray(heap)

	arrHeap = append(arrHeap, val)
	i := len(arrHeap) - 1

	PercolateUp(i, arrHeap)

	*heap = *ArrayToTree(arrHeap)
}

// Pop removes and returns the top element from a heap.
// Time O(log(N)), the new root has to be swapped with its child on each level
// until it reaches the bottom level of the heap
// Space O(N), since we need an array containing elements from input heap.
func Pop(heapTree *tree.Node) int {
	heap := TreeToArray(heapTree)
	heapPtr := &heap

	res := heap[0]
	heap[0], *heapPtr = heap[len(heap)-1], heap[:len(heap)-1]

	PercolateDown(0, heap)

	*heapTree = *ArrayToTree(heap)

	return res
}

// Heapify transforms a binary tree into a heap.
// Time O(N), since it iterates through the elements from the bottom level upwards,
// potentially shifting down elements multiple times to maintain the heap property
// Space O(N), since we need an array containing elements from input tree.
func Heapify(tree *tree.Node) {
	arrTree := TreeToArray(tree)

	const (
		divisor = 2
		offset  = 2
	)

	for i := (len(arrTree) - offset) / divisor; i >= 0; i-- {
		PercolateDown(i, arrTree)
	}

	*tree = *ArrayToTree(arrTree)
}

func PercolateUp(i int, heap []int) {
	const divisor = 2

	for heap[i] < heap[(i-1)/divisor] {
		heap[(i-1)/divisor], heap[i] = heap[i], heap[(i-1)/divisor]
		i = (i - 1) / divisor
	}
}

func PercolateDown(i int, heap []int) {
	const (
		multiplier = 2
		offset1    = 1
		offset2    = 2
	)

	if len(heap) == 2 && heap[0] > heap[1] {
		heap[0], heap[1] = heap[1], heap[0]

		return
	}

	for i*multiplier+offset2 < len(heap) &&
		(heap[i] > heap[i*multiplier+offset1] || heap[i] > heap[i*multiplier+offset2]) {
		leftChildGreaterRightChild := heap[i*multiplier+offset1] > heap[i*multiplier+offset2]

		if leftChildGreaterRightChild {
			heap[i], heap[i*multiplier+offset2] = heap[i*multiplier+offset2], heap[i]
			i = i*multiplier + offset2

			continue
		}

		heap[i], heap[i*multiplier+offset1] = heap[i*multiplier+offset1], heap[i]
		i = i*multiplier + offset1
	}
}

func TreeToArray(heap *tree.Node) []int {
	var arr []int

	queue := []*tree.Node{heap}

	for len(queue) > 0 {
		heap, queue = queue[0], queue[1:]
		arr = append(arr, heap.Val)

		if heap.Left != nil {
			queue = append(queue, heap.Left)
		}

		if heap.Right != nil {
			queue = append(queue, heap.Right)
		}
	}

	return arr
}

func ArrayToTree(arr []int) *tree.Node {
	var helper func(node *tree.Node, i int)

	helper = func(node *tree.Node, i int) {
		if i*2+1 < len(arr) {
			node.Left = &tree.Node{Val: arr[i*2+1]}
			helper(node.Left, i*2+1)
		}

		const (
			multiplier       = 2
			rightChildOffset = 2
		)

		if i*multiplier+rightChildOffset < len(arr) {
			node.Right = &tree.Node{Val: arr[i*multiplier+rightChildOffset]}
			helper(node.Right, i*multiplier+rightChildOffset)
		}
	}

	heap := &tree.Node{Val: arr[0]}
	helper(heap, 0)

	return heap
}
