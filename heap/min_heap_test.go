package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {

	h := &IntHeap{1, 2, 9, 4, 5, 6, 3, 8, 0, 10}

	heap.Init(h)

	heap.Push(h, -1)

	fmt.Println("min heap:", (*h)[0])

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}

}
