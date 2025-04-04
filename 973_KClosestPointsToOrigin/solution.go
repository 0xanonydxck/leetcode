package kclosestpointstoorigin

func buildHeap(points [][]int) *Heap[[]int] {
	heap := new(Heap[[]int])

	for _, point := range points {
		distance := point[0]*point[0] + point[1]*point[1]
		heap.Push(HeapItem[[]int]{Value: point, Priority: distance})
	}

	return heap
}

func kClosest(points [][]int, k int) [][]int {
	heap := buildHeap(points)
	result := make([][]int, k)
	for i := range k {
		result[i] = heap.Pop().Value
	}

	return result
}
