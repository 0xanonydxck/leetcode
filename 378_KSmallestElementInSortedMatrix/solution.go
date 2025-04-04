package ksmallestelementinsortedmatrix

func buildHeap(matrix [][]int) *Heap[int] {
	heap := Heap[int]{}
	for _, row := range matrix {
		for _, col := range row {
			heap.Push(HeapItem[int]{col, col})
		}
	}
	return &heap
}

func kthSmallest(matrix [][]int, k int) int {
	heap := buildHeap(matrix)
	result := 0
	for i := 0; i < k; i++ {
		result = heap.Pop().Value
	}
	return result
}
