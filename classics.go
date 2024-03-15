package circularqueue

import "fmt"

type classics[T any] struct {
	capp1 int
	queue []T

	front int
	rear  int
}

func New[T any](capacity int) Queue[T] {
	if capacity <= 0 {
		panic(fmt.Sprintf("runtime error: makeslice: cap %d out of range", capacity))
	}
	return &classics[T]{
		capp1: capacity + 1,
		queue: make([]T, capacity+1),
	}
}

func (cq *classics[T]) Push(item T) {
	if cq.Full() {
		cq.Pop()
	}
	cq.queue[cq.rear] = item
	cq.rear = (cq.rear + 1) % cq.capp1
}

func (cq *classics[T]) Peek() (t T, ok bool) {
	if cq.Empty() {
		return
	}
	return cq.queue[cq.front], true
}

func (cq *classics[T]) Pop() (t T, ok bool) {
	t, ok = cq.Peek()
	if !ok {
		return
	}
	cq.front = (cq.front + 1) % cq.capp1
	return
}

func (cq *classics[T]) GetItems() []T {
	return cq.getItemsFast()
}

func (cq *classics[T]) Size() int {
	return (cq.rear + cq.capp1 - cq.front) % cq.capp1
}

func (cq *classics[T]) Cap() int {
	return cq.capp1 - 1
}

func (cq *classics[T]) Empty() bool {
	return cq.Size() == 0
}

func (cq *classics[T]) Full() bool {
	return cq.Size() == cq.Cap()
}

func (cq *classics[T]) getItemsFast() []T {
	// getItemsFast is three times faster than getItemsSlow
	var vs = make([]T, cq.Size())
	if len(vs) == 0 {
		return vs
	}
	if cq.front < cq.rear {
		copy(vs, cq.queue[cq.front:cq.rear])
	} else {
		tail := cq.queue[cq.front:]
		n := copy(vs, tail)
		copy(vs[n:], cq.queue[:cq.rear])
		//copy(vs[copy(vs, cq.queue[cq.front:]):], cq.queue[:cq.rear])
	}
	return vs
}

func (cq *classics[T]) getItemsSlow() []T {
	size := cq.Size()
	items := make([]T, 0, size)
	for i := 0; i < size; i++ {
		index := (cq.front + i) % cq.capp1
		items = append(items, cq.queue[index])
	}
	return items
}
