package circularqueue

import "sync"

func NewSafe[T any](q Queue[T]) Queue[T] {
	if _, ok := q.(*safe[T]); ok {
		return q
	}
	return &safe[T]{q: q}
}

type safe[T any] struct {
	q Queue[T]
	m sync.RWMutex
}

func (s *safe[T]) Pop() (T, bool) {
	s.m.Lock()
	defer s.m.Unlock()
	return s.q.Pop()
}

func (s *safe[T]) Peek() (T, bool) {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.q.Peek()
}

func (s *safe[T]) Push(t T) {
	s.m.Lock()
	defer s.m.Unlock()
	s.q.Push(t)
}

func (s *safe[T]) Cap() int {
	return s.q.Cap()
}

func (s *safe[T]) Size() int {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.q.Size()
}

func (s *safe[T]) Empty() bool {
	return s.Size() == 0
}

func (s *safe[T]) Full() bool {
	return s.Size() == s.Cap()
}

func (s *safe[T]) GetItems() []T {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.q.GetItems()
}
