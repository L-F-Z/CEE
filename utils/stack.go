package utils

// Stack type
type Stack struct {
	top    *node
	length int64
}

// node type for stack
type node struct {
	value interface{}
	prev  *node
}

// NewStack create a new stack
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// Len return the number of items in the stack
func (s *Stack) Len() int64 {
	return s.length
}

// Peek view the top item on the stack
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Pop pop the top item of the stack and return it
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push push a value onto the top of the stack
func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}
