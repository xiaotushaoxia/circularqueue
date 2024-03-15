package circularqueue

import (
	"fmt"
	"testing"
)

func TestCC(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			msg := fmt.Sprintf("%v", p)
			want := "runtime error: makeslice: cap -10 out of range"
			if msg != want {
				t.Fatalf("panic msg not equal: \ngot : %s\nwant: %s", msg, want)
			}
		}
	}()
	q := New[int](-10)
	q.Cap()
}

func BenchmarkClassicsLargeCap(b *testing.B) {
	q := New[int](50000) //  classics is faster when cap and size is large
	for i := 0; i < b.N; i++ {
		bTestQueue(q)
	}
}

func BenchmarkSliceLargeCap(b *testing.B) {
	q := NewSlice[int](50000)
	for i := 0; i < b.N; i++ {
		bTestQueue(q)
	}
}

func BenchmarkClassicsSmallCap(b *testing.B) {
	q := New[int](1000)
	for i := 0; i < b.N; i++ {
		bTestQueue(q)
	}
}

func BenchmarkSliceSmallCap(b *testing.B) {
	q := NewSlice[int](1000)
	for i := 0; i < b.N; i++ {
		bTestQueue(q)
	}
}

func bTestQueue[T any](q Queue[T]) {
	var tt T
	for i := 0; i < q.Cap(); i++ {
		q.Push(tt)
		q.Pop()
	}

	for i := 0; i < q.Cap()*2; i++ {
		q.Push(tt)
		q.Pop()
	}
	for i := 0; i < q.Cap()*2; i++ {
		q.Push(tt)
	}
	for i := 0; i < q.Cap()*2; i++ {
		q.Pop()
	}
	for i := 0; i < q.Cap()*2; i++ {
		q.Push(tt)
	}
	for i := 0; i < q.Cap()*4; i++ {
		q.Pop()
	}
}

func TestQueueClassics(t *testing.T) {
	q := New[int](10)
	testQueue(q, t)
}

func TestQueueSlice(t *testing.T) {
	q := NewSlice[int](10)
	testQueue(q, t)
}

func TestQueueSafe(t *testing.T) {
	q := NewSlice[int](10)
	testQueue(NewSafe(q), t)
}

func testQueue(q Queue[int], t *testing.T) {
	_, b := q.Peek()
	if b {
		t.Fatalf("empty peek return true, want false")
	}
	q.Push(1)
	v, ok := q.Peek()
	if !ok || v != 1 {
		t.Fatalf("Peek return %d,%t, want %d,%t", v, ok, 1, true)
	}
	s := q.Size()
	if s != 1 {
		t.Fatalf("Size return %d, want %d", s, 1)
	}
	v, ok = q.Pop()
	if !ok || v != 1 {
		t.Fatalf("Pop return %d,%t, want %d,%t", v, ok, 1, true)
	}
	s = q.Size()
	if s != 0 {
		t.Fatalf("Size return %d, want %d", s, 0)
	}

	var a []int
	for i := 0; ; i++ {
		if q.Full() {
			break
		}
		q.Push(i)
		a = append(a, i)
		if msg := sliceEq(a, q.GetItems()); msg != "" {
			t.Fatalf("slice not equel when push: %s \n%v \n%v", msg, a, q.GetItems())
		}
	}

	// test  push on full
	for i := 888; i < 999; i++ {
		q.Push(i)
		a = append(a, i)
		a = a[1:]
		if msg := sliceEq(a, q.GetItems()); msg != "" {
			t.Fatalf("slice not equel when push on full: %s \n%v \n%v", msg, a, q.GetItems())
		}
	}

	for {
		if q.Empty() {
			break
		}
		q.Pop()
		a = a[1:]
		if msg := sliceEq(a, q.GetItems()); msg != "" {
			t.Fatalf("slice not equel when pop: %s \n%v \n%v", msg, a, q.GetItems())
		}
	}

	// test pop on empty
	for i := 0; i < 10; i++ {
		q.Pop()
		if msg := sliceEq(a, q.GetItems()); msg != "" {
			t.Fatalf("slice not equel when pop on empty: %s \n%v \n%v", msg, a, q.GetItems())
		}

	}
}

func sliceEq[T comparable](s1, s2 []T) string {
	if len(s1) != len(s2) {
		return "length not equal"
	}
	for i := 0; i < len(s2); i++ {
		if s2[i] != s1[i] {
			return fmt.Sprintf("item %d not equal: %v, %v", i, s1[i], s2[i])
		}
	}
	return ""
}
