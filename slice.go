package circularqueue

import "fmt"

func NewSlice[T any](capacity int) Queue[T] {
	if capacity <= 0 {
		panic(fmt.Sprintf("runtime error: makeslice: cap %d out of range", capacity))
	}
	return &slice[T]{cap: capacity}
}

type slice[T any] struct {
	s   []T
	cap int
}

func (q *slice[T]) Pop() (t T, ok bool) {
	t, ok = q.Peek()
	if !ok {
		return
	}
	q.s = q.s[1:] // O(1)
	return
}

func (q *slice[T]) Peek() (t T, ok bool) {
	if q.Empty() {
		return
	}
	return q.s[0], true
}

func (q *slice[T]) Push(t T) {
	if q.Full() {
		q.Pop()
	}
	q.s = append(q.s, t)
}

func (q *slice[T]) Size() int {
	return len(q.s)
}

func (q *slice[T]) Cap() int {
	return q.cap
}

func (q *slice[T]) Empty() bool {
	return q.Size() == 0
}

func (q *slice[T]) Full() bool {
	return q.Size() == q.Cap()
}

func (q *slice[T]) GetItems() []T {
	var vs = make([]T, len(q.s))
	copy(vs, q.s)
	return vs
}
