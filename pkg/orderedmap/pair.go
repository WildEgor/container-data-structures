package orderedmap

import list "github.com/WildEgor/container-data-structures/internal"

type Pair[K comparable, V any] struct {
	Key     K
	Value   V
	element *list.Element[*Pair[K, V]]
}

// Next returns a pointer to the next pair.
func (p *Pair[K, V]) Next() *Pair[K, V] {
	if p == nil {
		return nil
	}

	return listElementToPair(p.element.Next())
}

// Prev returns a pointer to the previous pair.
func (p *Pair[K, V]) Prev() *Pair[K, V] {
	if p == nil {
		return nil
	}

	return listElementToPair(p.element.Prev())
}
