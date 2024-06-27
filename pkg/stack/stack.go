package stack

// Stack represent stack data structure
type Stack[K comparable] struct {
	elements []K
}

type option[K comparable] struct {
	capacity int
	data     []K
}

type Options[K comparable] func(config *option[K])

// WithCapacity allows giving a capacity hint for the Stack, like make([]K, len, capacity).
func WithCapacity[K comparable, V any](capacity int) Options[K] {
	return func(c *option[K]) {
		c.capacity = capacity
	}
}

// WithData allows passing in initial data for the Stack
func WithData[K comparable](data ...K) Options[K] {
	return func(c *option[K]) {
		c.data = data
		if c.capacity < len(data) {
			c.capacity = len(data)
		}
	}
}

// New creates Stack
func New[K comparable](options ...any) *Stack[K] {

	stack := &Stack[K]{}

	var config option[K]
	for _, untypedOption := range options {
		switch option := untypedOption.(type) {
		default:
			errInvalidOption()
		case int:
			if len(options) != 1 {
				errInvalidOption()
			}
			config.capacity = option
		case Options[K]:
			option(&config)
		}
	}

	stack.elements = make([]K, 0, config.capacity)

	stack.Push(config.data...)

	return stack
}

// Push add element to Stack
func (s *Stack[K]) Push(data ...K) {
	for _, el := range data {
		s.elements = append(s.elements, el)
	}
}

// Pop removes the last element in the stack and returns it
func (s *Stack[K]) Pop() (value K) {
	if s.Empty() {
		return
	}
	e := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return e
}

// Top returns the last element in the stack without removing it
func (s *Stack[K]) Top() (value K) {
	if s.Empty() {
		return
	}
	return s.elements[len(s.elements)-1]
}

// Len get Stack size
func (s *Stack[K]) Len() int {
	return len(s.elements)
}

// Empty check if Stack is empty
func (s *Stack[K]) Empty() bool {
	return len(s.elements) == 0
}

// Clear remove elements from Stack
func (s *Stack[K]) Clear() {
	s.elements = make([]K, 0)
}

// Values return elements from Stack
func (s *Stack[K]) Values() []K {
	return s.elements
}
