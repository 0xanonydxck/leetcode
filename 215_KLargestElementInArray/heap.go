package klargestelementinarray

type HeapItem[T any] struct {
	Value    T
	Priority int
}

type Heap[T any] struct {
	heap []HeapItem[T]
}

func (h *Heap[T]) Push(n HeapItem[T]) {
	h.heap = append(h.heap, n)
	h.BubbleUp()
}

func (h *Heap[T]) BubbleUp() {
	index := h.Size() - 1
	for index > 0 {
		el := h.heap[index]
		parentIndex := (index - 1) / 2

		parent := h.heap[parentIndex]
		if parent.Priority <= el.Priority {
			break
		}

		h.heap[index] = parent
		h.heap[parentIndex] = el
		index = parentIndex
	}
}

func (h *Heap[T]) BubbleDown() {
	index := 0
	min := index
	size := h.Size()

	for index < size {
		left := 2*index + 1
		right := left + 1
		if (left < size && h.heap[left].Priority < h.heap[min].Priority) ||
			(right < size && h.heap[right].Priority < h.heap[min].Priority) {
			if right < size && h.heap[right].Priority < h.heap[left].Priority {
				min = right
			} else {
				min = left
			}
		}

		if min == index {
			break
		}
		h.heap[min], h.heap[index] = h.heap[index], h.heap[min]
		index = min
	}
}

func (h *Heap[T]) Pop() (min HeapItem[T]) {
	min = h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	h.BubbleDown()
	return
}

func (h *Heap[T]) Peek() HeapItem[T] {
	return h.heap[0]
}

func (h *Heap[T]) Size() int {
	return len(h.heap)
}
