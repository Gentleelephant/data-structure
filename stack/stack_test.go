package stack

import (
	"github.com/Gentleelephant/data-structure/sort"
	"testing"
)

func TestStack(t *testing.T) {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
}

func TestName(t *testing.T) {
	quickSort := sort.QuickSort([]int{1, 2, 3, 4, 6, 5, 7, 8, 9, 10})
	t.Log(quickSort)
}
