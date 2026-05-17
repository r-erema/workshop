package kthlargestelementinastream

import (
	"container/heap"
)

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x any) {
	if xInt, ok := x.(int); ok {
		*h = append(*h, xInt)
	}
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

type kthLargest struct {
	k       int
	minHeap *IntHeap
}

func Constructor(k int, nums []int) kthLargest {
	minHeap := IntHeap(nums)
	heap.Init(&minHeap)

	for len(minHeap) > k {
		heap.Pop(&minHeap)
	}

	return kthLargest{
		k:       k,
		minHeap: &minHeap,
	}
}

func (l *kthLargest) Add(val int) int {
	heap.Push(l.minHeap, val)

	if len(*l.minHeap) > l.k {
		heap.Pop(l.minHeap)
	}

	return (*l.minHeap)[0]
}
