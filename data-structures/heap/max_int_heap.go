package heap

// A MaxIntHeap is a max-heap of ints.
type MaxIntHeap []int

func NewMaxIntHeap() MaxIntHeap {
	return MaxIntHeap{}
}

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.

func (h *MaxIntHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		panic("invalid type")
	}

	*h = append(*h, val)
}

func (h *MaxIntHeap) Pop() any {
	if len(*h) == 0 {
		panic("heap is empty")
	}

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MaxIntHeap) Top() int {
	if len(h) == 0 {
		panic("heap is empty")
	}

	return h[0]
}
