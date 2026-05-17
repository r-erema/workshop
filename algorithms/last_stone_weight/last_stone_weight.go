package laststoneweight

// LastStoneWeight finds the last stone weight after smashing stones.
// Time O(N*logN), since we need to pop each element from the heap
// Space O(n), since the heap size equals the input.
func LastStoneWeight(stones []int) int {
	heap := Heapify(stones)

	for len(heap) > 1 {
		stone1, stone2 := Pop(&heap), Pop(&heap)
		newStone := stone1 - stone2

		if newStone > 0 {
			Push(newStone, &heap)
		}
	}

	if len(heap) == 1 {
		return heap[0]
	}

	return 0
}

func Heapify(stones []int) []int {
	var heap []int

	for i := range stones {
		Push(stones[i], &heap)
	}

	return heap
}

func Push(node int, heap *[]int) {
	*heap = append(*heap, node)
	PercolateUp(*heap)
}

func Pop(heap *[]int) int {
	popped := (*heap)[0]
	(*heap)[0] = (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]
	PercolateDown(*heap)

	return popped
}

func PercolateUp(heap []int) {
	const (
		divisor = 2
		offset  = 1
	)

	child1 := len(heap) - 1
	parent := (child1 - 1) / divisor
	child2 := parent*divisor + offset

	for heap[parent] < heap[child1] {
		if heap[child2] > heap[child1] {
			heap[child2], heap[child1] = heap[child1], heap[child2]
		} else {
			heap[parent], heap[child1] = heap[child1], heap[parent]
		}

		child1 = parent
		parent = (child1 - 1) / divisor
		child2 = parent*divisor + offset
	}
}

func PercolateDown(heap []int) {
	const (
		divisor = 2
		offset  = 1
	)

	parent, child1, child2 := 0, 1, 2

	for (child1 < len(heap) && heap[parent] < heap[child1]) || (child2 < len(heap) && heap[parent] < heap[child2]) {
		if child2 < len(heap) && heap[child1] < heap[child2] {
			heap[parent], heap[child2], parent = heap[child2], heap[parent], child2
		} else {
			heap[parent], heap[child1], parent = heap[child1], heap[parent], child1
		}

		child1, child2 = parent*divisor+offset, (parent+offset)*divisor
	}
}
