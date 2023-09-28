package stack

import "errors"


type stack struct {
	top int
	maxIndices int
	values []interface{}
}

func Init(size int) (stack, error) {
	if size < 1 {
		return stack{}, errors.New("the size should be a minimum of 1")
	}

	var s stack
	s.top = -1
	s.maxIndices = size - 1
	return s, nil
}

func (s *stack) Push(v interface{}) error {
	tmp := s.top

	if tmp == s.maxIndices {
		return errors.New("the stack is full")
	}

	s.top = tmp + 1
	s.values = append(s.values, v)
	return nil
}

func (s *stack) Pop() interface{} {
	if s.top == -1 {
		return nil
	}

	val := s.values[s.top]
	s.values = s.values[:len(s.values) - 1]
	s.top--
	return val
}

func (s *stack) Peek() interface{} {
	if s.top == -1 {
		return nil
	}

	return s.values[s.top]
}

func (s *stack) IsEmpty() bool {
	return s.top == -1
}

func (s *stack) IsFull() bool {
	return s.top == s.maxIndices
}
