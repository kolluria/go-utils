package heap

// A MinIntHeap is a min-heap of ints.
type MinIntHeap []int

func NewMinIntHeap() MinIntHeap {
	return MinIntHeap{}
}

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.

func (h *MinIntHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		panic("invalid type")
	}

	*h = append(*h, val)
}

func (h *MinIntHeap) Pop() any {
	if len(*h) == 0 {
		panic("heap is empty")
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinIntHeap) Top() int {
	if len(h) < 1 {
		panic("heap is empty")
	}

	return h[0]
}
