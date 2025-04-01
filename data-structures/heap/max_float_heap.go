package heap

type MaxFloatHeap []float64

func NewMaxFloatHeap() MaxFloatHeap {
	return MaxFloatHeap{}
}

func (h MaxFloatHeap) Len() int { return len(h) }

func (h MaxFloatHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h MaxFloatHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.

func (h *MaxFloatHeap) Push(x any) {
	val, ok := x.(float64)
	if !ok {
		panic("invalid type")
	}

	*h = append(*h, val)
}

func (h *MaxFloatHeap) Pop() any {
	if len(*h) == 0 {
		panic("heap is empty")
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MaxFloatHeap) Top() float64 {
	if len(h) == 0 {
		panic("heap is empty")
	}

	return h[0]
}
