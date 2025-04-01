package heap

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxIntHeap_Len(t *testing.T) {
	h := &MaxIntHeap{}

	assert.Equal(t, 0, h.Len())

	h = &MaxIntHeap{3, 2, 1}

	assert.Equal(t, 3, h.Len())
}

func TestMaxIntHeap_Less(t *testing.T) {
	h := &MaxIntHeap{3, 2, 1}

	assert.True(t, h.Less(0, 1))
	assert.True(t, h.Less(1, 2))
}

func TestMaxIntHeap_Swap(t *testing.T) {
	h := &MaxIntHeap{3, 2, 1}

	h.Swap(0, 2)

	assert.Equal(t, 1, (*h)[0])
	assert.Equal(t, 3, (*h)[2])
}

func TestMaxIntHeap_Push(t *testing.T) {
	h := &MaxIntHeap{3, 2, 1}

	h.Push(4)

	assert.Equal(t, 4, (*h)[3])
	assert.Panics(t, func() { h.Push("invalid") })
}

func TestMaxIntHeap_Pop(t *testing.T) {
	h := &MaxIntHeap{3, 2, 1}

	assert.Equal(t, 1, h.Pop())
	assert.Equal(t, 2, h.Pop())
	assert.Equal(t, 3, h.Pop())
	assert.Panics(t, func() { h.Pop() })
}

func TestMaxIntHeap_Top(t *testing.T) {
	h := &MaxIntHeap{3, 2, 1}

	assert.Equal(t, 3, h.Top())
	h.Pop()
	assert.Equal(t, 3, h.Top())
	h.Pop()
	assert.Equal(t, 3, h.Top())
	h.Pop()
	assert.Panics(t, func() { h.Top() })
}

func TestMaxIntHeap(t *testing.T) {
	heap.Init(&MaxIntHeap{})

	h := MaxIntHeap{}

	// push 3, 2, 5
	heap.Push(&h, 3)
	assert.Equal(t, 3, h.Top())
	heap.Push(&h, 2)
	assert.Equal(t, 3, h.Top())
	heap.Push(&h, 5)
	assert.Equal(t, 5, h.Top())

	// pop 5, 3, 2
	assert.Equal(t, 5, heap.Pop(&h))
	assert.Equal(t, 3, heap.Pop(&h))
	assert.Equal(t, 2, heap.Pop(&h))
}
