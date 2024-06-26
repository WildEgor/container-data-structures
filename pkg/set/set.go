package set

// Set represent unique set structure
type Set[K comparable] struct {
	data map[K]bool
}

// New creates Set
func New[K comparable](data ...K) *Set[K] {
	set := &Set[K]{}
	set.data = make(map[K]bool, len(data))
	set.Add(data...)
	return set
}

// Add element to Set
func (s *Set[K]) Add(data ...K) {
	for _, datum := range data {
		s.data[datum] = true
	}
}

// Delete element from Set
func (s *Set[K]) Delete(data K) {
	delete(s.data, data)
}

// Clear Set
func (s *Set[K]) Clear() {
	clear(s.data)
}

// Contains check if element in Set
func (s *Set[K]) Contains(data K) bool {
	if _, ok := s.data[data]; !ok {
		return false
	}
	return true
}

// Has alias for Contains
func (s *Set[K]) Has(data K) bool {
	return s.Contains(data)
}

// Len returns the length keys
func (s *Set[K]) Len() int {
	return len(s.data)
}
