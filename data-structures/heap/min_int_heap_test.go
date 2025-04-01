package heap

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinIntHeap_Len(t *testing.T) {
	h := &MinIntHeap{}

	assert.Equal(t, 0, h.Len())

	h = &MinIntHeap{1, 2, 3}

	assert.Equal(t, 3, h.Len())
}

func TestMinIntHeap_Less(t *testing.T) {
	h := &MinIntHeap{1, 2, 3}

	assert.True(t, h.Less(0, 1))
	assert.True(t, h.Less(1, 2))
}

func TestMinIntHeap_Swap(t *testing.T) {
	h := &MinIntHeap{1, 2, 3}

	h.Swap(0, 2)

	assert.Equal(t, 3, (*h)[0])
	assert.Equal(t, 1, (*h)[2])
}

func TestMinIntHeap_Push(t *testing.T) {
	h := &MinIntHeap{1, 2, 3}

	h.Push(4)

	assert.Equal(t, 4, (*h)[3])
	assert.Panics(t, func() { h.Push("invalid") })
}

func TestMinIntHeap_Pop(t *testing.T) {
	h := &MinIntHeap{1, 2, 3}

	assert.Equal(t, 3, h.Pop())
	assert.Equal(t, 2, h.Pop())
	assert.Equal(t, 1, h.Pop())
	assert.Panics(t, func() { h.Pop() })
}

func TestMinIntHeap_Top(t *testing.T) {
	h := &MinIntHeap{1, 2, 3}

	assert.Equal(t, 1, h.Top())
	h.Pop()
	assert.Equal(t, 1, h.Top())
	h.Pop()
	assert.Equal(t, 1, h.Top())
	h.Pop()
	assert.Panics(t, func() { h.Top() })
}

// TestMinIntHeap is a test function that simulates the usage of MinIntHeap.
func TestMinIntHeap(t *testing.T) {
	heap.Init(&MinIntHeap{})

	h := MinIntHeap{}

	// Push 3, 2, 5
	heap.Push(&h, 3)
	assert.Equal(t, 3, h.Top())
	heap.Push(&h, 2)
	assert.Equal(t, 2, h.Top())
	heap.Push(&h, 5)
	assert.Equal(t, 2, h.Top())

	// Pop 2, 3, 5
	assert.Equal(t, 2, heap.Pop(&h))
	assert.Equal(t, 3, heap.Pop(&h))
	assert.Equal(t, 5, heap.Pop(&h))
}
