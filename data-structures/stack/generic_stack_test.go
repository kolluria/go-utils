package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New(5)
	assert.Equal(t, 5, s.capacity)
	assert.Equal(t, 0, len(s.stack))
}

func TestStack_Push(t *testing.T) {
	s, exp := New(5), make([]any, 0, 5)

	for i := 0; i < 5; i++ {
		s.Push(i)
		exp = append(exp, i)
		assert.Equal(t, exp, s.stack)
	}

	assert.Panics(t, func() { s.Push(5) })
}

func TestStack_Pop(t *testing.T) {
	s := New(5)

	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	for i := 4; i >= 0; i-- {
		assert.Equal(t, i, s.Pop())
	}

	assert.Panics(t, func() { s.Pop() })
}

func TestStack_Top(t *testing.T) {
	s := New(5)

	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	assert.Equal(t, 4, s.Top())

	for i := 4; i >= 0; i-- {
		s.Pop()
	}

	assert.Panics(t, func() { s.Top() })
}
