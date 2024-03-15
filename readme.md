# circularqueue

Two implementations are provided. 

The first one is the classic circular queue(New), and the second one simulates using slices(NewSlice).

slice implementation is faster when cap is small

Additionally, NewSafe converts the queue into a thread-safe structure.

# usage

```go
package main

import (
	"fmt"

	"github.com/xiaotushaoxia/circularqueue"
)

func main() {
	q := circularqueue.New[int](10)
	for i := 0; i < 15; i++ {
		q.Push(i)
	}
	fmt.Println(q.GetItems()) // [5 6 7 8 9 10 11 12 13 14]

	fmt.Println(q.Peek()) // 5 true

	fmt.Println(q.Pop()) // 5 true

	fmt.Println(q.GetItems()) // [6 7 8 9 10 11 12 13 14]

	for !q.Empty() {
		q.Pop()
	}

	fmt.Println(q.Pop())      // 0 false
	fmt.Println(q.GetItems()) // []
}
```