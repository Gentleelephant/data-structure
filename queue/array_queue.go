package queue

import "sync"

type ArrayQueue struct {
	arr  []interface{}
	size int
	lock sync.Mutex
}

func (q *ArrayQueue) Size() int {
	return len(q.arr)
}

// Add 入队列
func (q *ArrayQueue) Add(v interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.arr = append(q.arr, v)
	q.size++
}

func (q *ArrayQueue) Remove() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	// 如果队列为空,报错
	if q.size == 0 {
		panic("queue is empty")
	}
	// 取出队列头部元素
	v := q.arr[0]
	// 将队列头部元素移除
	q.arr = q.arr[1:]
	q.size--
	return v
}
