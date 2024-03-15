package circularqueue

type Queue[T any] interface {
	Pop() (T, bool)
	Peek() (T, bool)
	Push(T)

	Cap() int
	Size() int
	Empty() bool
	Full() bool
	GetItems() []T
}
