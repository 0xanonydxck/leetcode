package mergeksortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildHeap(lists []*ListNode) *Heap[int] {
	heap := Heap[int]{}
	for _, list := range lists {
		current := list
		for current != nil {
			heap.Push(HeapItem[int]{current.Val, current.Val})
			current = current.Next
		}
	}

	return &heap
}

func mergeKLists(lists []*ListNode) *ListNode {
	heap := buildHeap(lists)
	if heap.Size() == 0 {
		return nil
	}

	root := &ListNode{Val: heap.Pop().Value}
	current := root
	for heap.Size() > 0 {
		value := heap.Pop().Value
		node := &ListNode{Val: value}
		current.Next = node
		current = current.Next
	}

	return root
}
