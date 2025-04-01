package stack

type Stack struct {
	capacity int
	stack    []any
}

func New(capacity int) Stack {
	return Stack{
		capacity: capacity,
		stack:    make([]any, 0, capacity),
	}
}

func (s *Stack) Push(val any) {
	if len(s.stack) == s.capacity {
		panic("cannot insert beyond capacity")
	}

	s.stack = append(s.stack, val)
}

func (s *Stack) Pop() (pop any) {
	if len(s.stack) <= 0 {
		panic("cannot pop from an empty stack")
	}

	pop = s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return
}

func (s *Stack) Top() any {
	if len(s.stack) <= 0 {
		panic("no top in an empty stack")
	}

	return s.stack[len(s.stack)-1]
}
