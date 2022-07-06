package heap

// Push 和 Pop 方法需要使用指针，因为它们会修改 slice 的长度，而不仅仅只内容。

type MinHeap struct {
	arr []int
}

func (h MinHeap) Len() int {
	return len(h.arr)
}

func (h MinHeap) Less(i, j int) bool {
	// h[i] < h[j] 为小顶堆，大顶堆为 h[i] > h[j]
	return h.arr[i] < h.arr[j]
}

func (h MinHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *MinHeap) Push(x interface{}) {
	h.arr = append(h.arr, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := h.arr
	n := len(old)
	x := old[n-1]
	h.arr = old[0 : n-1]
	return x
}
