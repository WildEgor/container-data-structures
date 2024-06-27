package queue

import list "github.com/WildEgor/container-data-structures/internal"

type Deque[T any] struct {
	data *list.List[T]
}

func NewDeque[T any](items ...T) *Deque[T] {
	dq := &Deque[T]{
		data: list.New[T](),
	}

	for i := 0; i < len(items); i++ {
		dq.data.PushBack(items[i])
	}

	return dq
}

// Append inserts item at the back of the Deque in an *O(1)* time complexity.
func (d *Deque[T]) Append(item T) {
	d.data.PushBack(item)
}

// Prepend inserts item at the Deque's front in an *O(1)* time complexity.
func (d *Deque[T]) Prepend(item T) {
	d.data.PushFront(item)
}

// Pop removes and returns the back element of the Deque in an *O(1)* time complexity.
func (d *Deque[T]) Pop() (item T, ok bool) {
	lastElement := d.data.Back()
	if lastElement != nil {
		item = d.data.Remove(lastElement)
		ok = true
	}

	return
}

// Shift removes and returns the front element of the Deque in *O(1)* time complexity.
func (d *Deque[T]) Shift() (item T, ok bool) {
	firstElement := d.data.Front()
	if firstElement != nil {
		item = d.data.Remove(firstElement)
		ok = true
	}

	return
}

// First returns the first value stored in the Deque in *O(1)* time complexity.
func (d *Deque[T]) First() (item T, ok bool) {
	frontItem := d.data.Front()
	if frontItem != nil {
		item = frontItem.Value
		ok = true
	}

	return
}

// Last returns the last value stored in the Deque in *O(1)* time complexity.
func (d *Deque[T]) Last() (item T, ok bool) {
	if backItem := d.data.Back(); backItem != nil {
		item = backItem.Value
		ok = true
	}

	return
}

// Size returns the Deque's size.
func (d *Deque[T]) Size() uint {
	return uint(d.data.Len())
}

// Empty checks if the Deque is empty.
func (d *Deque[T]) Empty() bool {
	return d.data.Len() == 0
}
