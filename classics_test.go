package circularqueue

import (
	"fmt"
	"testing"
)

func Test_classice_getItems(t *testing.T) {
	queue := New[int](20).(*classics[int])
	for i := 0; i < 25; i++ {
		for j := 0; j < 14; j++ {
			queue.Push(i * j)
			if fmt.Sprintln(queue.getItemsSlow()) != fmt.Sprintln(queue.getItemsFast()) {
				t.Fatalf(fmt.Sprintln(queue.getItemsSlow()) + "not equal to " + fmt.Sprintln(queue.getItemsFast()) + " " + fmt.Sprintln(i, j))
			}
		}
	}
	for i := 0; i < queue.Size(); i++ {
		queue.Pop()
		if fmt.Sprintln(queue.getItemsSlow()) != fmt.Sprintln(queue.getItemsFast()) {
			t.Fatalf(fmt.Sprintln(queue.getItemsSlow()) + "not equal to " + fmt.Sprintln(queue.getItemsFast()))
		}
	}
}

// Benchmark_getItemsFast1
// Benchmark_getItemsFast1-16       2142888               606.1 ns/op
// Benchmark_getItemsSlow1
// Benchmark_getItemsSlow1-16        633244              1935 ns/op
// Benchmark_getItemsFast2
// Benchmark_getItemsFast2-16       4051210               336.8 ns/op
// Benchmark_getItemsSlow2
// Benchmark_getItemsSlow2-16       1000000              1041 ns/op
func Benchmark_getItemsFast1(b *testing.B) {
	queue := New[int](100).(*classics[int])
	for i := 0; i < 150; i++ {
		queue.Push(i)
	}
	for i := 0; i < b.N; i++ {
		queue.getItemsFast()
	}
}

func Benchmark_getItemsSlow1(b *testing.B) {
	queue := New[int](100).(*classics[int])
	for i := 0; i < 150; i++ {
		queue.Push(i)
	}
	for i := 0; i < b.N; i++ {
		queue.getItemsSlow()
	}
}

func Benchmark_getItemsFast2(b *testing.B) {
	queue := New[int](100).(*classics[int])
	for i := 0; i < 50; i++ {
		queue.Push(i)
	}
	for i := 0; i < b.N; i++ {
		queue.getItemsFast()
	}
}

func Benchmark_getItemsSlow2(b *testing.B) {
	queue := New[int](100).(*classics[int])
	for i := 0; i < 50; i++ {
		queue.Push(i)
	}
	for i := 0; i < b.N; i++ {
		queue.getItemsSlow()
	}
}
