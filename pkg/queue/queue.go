package queue

import "sync"

type IQueue[T any] interface {
	Enqueue(item T)
	Dequeue() (item T, ok bool)
	Head() (item T, ok bool)
	Size() uint
}

// Queue is a First In First Out data structure implementation.
type Queue[T any] struct {
	container *Deque[T]
}

// NewQueue produces a new Queue instance.
func NewQueue[T any](items ...T) *Queue[T] {
	dq := NewDeque[T]()

	for i := 0; i < len(items); i++ {
		dq.data.PushFront(items[i])
	}

	return &Queue[T]{
		container: dq,
	}
}

// Enqueue adds an item at the back of the Queue in *O(1)* time complexity.
func (q *Queue[T]) Enqueue(item T) {
	q.container.Prepend(item)
}

// Dequeue removes and returns the Queue's front item in *O(1)* time complexity.
func (q *Queue[T]) Dequeue() (item T, ok bool) {
	return q.container.Pop()
}

// Head returns the Queue's front queue item in *O(1)* time complexity.
func (q *Queue[T]) Head() (item T, ok bool) {
	return q.container.Last()
}

// Size returns the size of the Queue.
func (q *Queue[T]) Size() uint {
	return q.container.Size()
}

var _ IQueue[any] = (*SyncQueue[any])(nil)

type SyncQueue[T any] struct {
	mu    sync.Mutex
	queue *Queue[T]
}

func NewSyncQueue[T any](items ...T) *SyncQueue[T] {
	sq := &SyncQueue[T]{
		queue: NewQueue[T](items...),
	}

	return sq
}

func (s *SyncQueue[T]) Enqueue(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.queue.Enqueue(item)
}

func (s *SyncQueue[T]) Dequeue() (item T, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.queue.Dequeue()
}

func (s *SyncQueue[T]) Head() (item T, ok bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.queue.Head()
}

func (s *SyncQueue[T]) Size() uint {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.queue.Size()
}
