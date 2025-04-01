package heap

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxFloatHeap_Len(t *testing.T) {
	h := &MaxFloatHeap{}

	assert.Equal(t, 0, h.Len())

	h = &MaxFloatHeap{3.0, 2.0, 1.0}

	assert.Equal(t, 3, h.Len())
}

func TestMaxFloatHeap_Less(t *testing.T) {
	h := &MaxFloatHeap{3.0, 2.0, 1.0}

	assert.True(t, h.Less(0, 1))
	assert.True(t, h.Less(1, 2))
}

func TestMaxFloatHeap_Swap(t *testing.T) {
	h := &MaxFloatHeap{3.0, 2.0, 1.0}

	h.Swap(0, 2)

	assert.Equal(t, 1.0, (*h)[0])
	assert.Equal(t, 3.0, (*h)[2])
}

func TestMaxFloatHeap_Push(t *testing.T) {
	h := &MaxFloatHeap{3.0, 2.0, 1.0}

	h.Push(4.0)

	assert.Equal(t, 4.0, (*h)[3])
	assert.Panics(t, func() { h.Push("invalid") })
}

func TestMaxFloatHeap_Pop(t *testing.T) {
	h := &MaxFloatHeap{3.0, 2.0, 1.0}

	assert.Equal(t, 1.0, h.Pop())
	assert.Equal(t, 2.0, h.Pop())
	assert.Equal(t, 3.0, h.Pop())
	assert.Panics(t, func() { h.Pop() })
}

func TestMaxFloatHeap_Top(t *testing.T) {
	h := &MaxFloatHeap{3.0, 2.0, 1.0}

	assert.Equal(t, 3.0, h.Top())
	h.Pop()
	assert.Equal(t, 3.0, h.Top())
	h.Pop()
	assert.Equal(t, 3.0, h.Top())
}

func TestMaxFloatHeap(t *testing.T) {
	heap.Init(&MaxFloatHeap{})

	h := &MaxFloatHeap{3.0, 2.0, 1.0}

	heap.Push(h, 4.0)
	assert.Equal(t, 4.0, h.Top())

	heap.Push(h, 5.0)
	assert.Equal(t, 5.0, h.Top())

	heap.Push(h, 6.0)
	assert.Equal(t, 6.0, h.Top())

	assert.Equal(t, 6.0, heap.Pop(h))
	assert.Equal(t, 5.0, heap.Pop(h))
	assert.Equal(t, 4.0, heap.Pop(h))
	assert.Equal(t, 3.0, heap.Pop(h))
	assert.Equal(t, 2.0, heap.Pop(h))
	assert.Equal(t, 1.0, heap.Pop(h))
}
