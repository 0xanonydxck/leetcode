package klargestelementinarray

func buildHeap(nums []int) *Heap[int] {
	heap := Heap[int]{}
	for _, num := range nums {
		heap.Push(HeapItem[int]{num, num})
	}
	return &heap
}

func findKthLargest(nums []int, k int) int {
	heap := buildHeap(nums)
	l := heap.Size() - k + 1
	result := 0
	for i := 0; i < l; i++ {
		result = heap.Pop().Value
	}

	return result
}
