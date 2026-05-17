package kclosestpointstoorigin

import (
	"iter"
	"math"
)

type point struct {
	coordinates        [2]int
	distanceFromOrigin float64
}

// KClosest finds the k closest points to the origin.
// Time O(N*log k), since we pop from and push to the heap which size is k
// Space O(k), as we allocate additional space of size k.
func KClosest(rawCoordinates [][]int, k int) [][]int {
	points := Heapify(rawCoordinates[:k])

	coordinates := func() iter.Seq[[]int] {
		return func(yield func([]int) bool) {
			for _, v := range rawCoordinates[k:] {
				if !yield(v) {
					return
				}
			}
		}
	}

	for c := range coordinates() {
		dist := EuclideanDistance([2]int(c))

		if dist < points[0].distanceFromOrigin {
			Pop(&points)
			Push(&points, [2]int(c))
		}
	}

	res := make([][]int, k)
	for i := range points {
		res[i] = points[i].coordinates[:]
	}

	return res
}

func Heapify(rawCoordinates [][]int) []point {
	heap := make([]point, 0, len(rawCoordinates))
	for i := range rawCoordinates {
		Push(&heap, [2]int(rawCoordinates[i]))
	}

	return heap
}

func EuclideanDistance(coordinates [2]int) float64 {
	originPoint := [2]int{0, 0}

	const power = 2

	return math.Sqrt(
		math.Pow(
			float64(coordinates[0])-float64(originPoint[0]),
			power,
		) + math.Pow(
			float64(coordinates[1])-float64(originPoint[1]),
			power,
		),
	)
}

func Push(heap *[]point, coordinates [2]int) {
	*heap = append(*heap, point{coordinates, EuclideanDistance(coordinates)})

	PercolateUpLastPoint(*heap)
}

func PercolateUpLastPoint(heap []point) {
	const (
		offset  = 2
		divisor = 2
	)

	parentIdx := (len(heap) - offset) / divisor

	if parentIdx >= 0 && heap[len(heap)-1].distanceFromOrigin > heap[parentIdx].distanceFromOrigin {
		heap[len(heap)-1], heap[parentIdx] = heap[parentIdx], heap[len(heap)-1]
		PercolateUpLastPoint(heap[:parentIdx+1])
	}
}

func Pop(heap *[]point) [2]int {
	res := (*heap)[0].coordinates
	(*heap)[0], *heap = (*heap)[len(*heap)-1], (*heap)[:len(*heap)-1]

	PercolateDownFirstPoint(*heap, 0)

	return res
}

func PercolateDownFirstPoint(heap []point, i int) {
	const (
		multiplier = 2
		offset1    = 1
		offset2    = 2
	)

	leftChildIdx, rightChildIdx := i*multiplier+offset1, i*multiplier+offset2

	if len(heap) > leftChildIdx && heap[i].distanceFromOrigin < heap[leftChildIdx].distanceFromOrigin {
		heap[i], heap[leftChildIdx] = heap[leftChildIdx], heap[i]

		if len(heap) > rightChildIdx && heap[i].distanceFromOrigin < heap[rightChildIdx].distanceFromOrigin {
			heap[i], heap[rightChildIdx] = heap[rightChildIdx], heap[i]
			PercolateDownFirstPoint(heap, rightChildIdx)
		}

		PercolateDownFirstPoint(heap, leftChildIdx)

		return
	}

	if len(heap) > rightChildIdx && heap[i].distanceFromOrigin < heap[rightChildIdx].distanceFromOrigin {
		heap[i], heap[rightChildIdx] = heap[rightChildIdx], heap[i]

		PercolateDownFirstPoint(heap, rightChildIdx)
	}
}
