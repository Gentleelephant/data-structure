package queue

import (
	"testing"
)

func TestArrayQueue(t *testing.T) {
	a := new(ArrayQueue)
	a.Add(1)
	a.Add(2)
	a.Add(3)
	t.Log("size:", a.Size())
	t.Log(a.Remove())
	// size
	t.Log("size:", a.Size())
	t.Log(a.Remove())
	// size
	t.Log("size:", a.Size())
	t.Log(a.Remove())
	// size
	t.Log("size:", a.Size())
	//t.Log(a.Remove())
	//t.Log(a.Size())
	//t.Log(a.Remove())
}
