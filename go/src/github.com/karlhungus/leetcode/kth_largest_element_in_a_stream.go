package main

type KthLargest struct {
	h *IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	for _, n := range nums {
		heap.Push(h, n)
	}
	for h.Len() > k {
		heap.Pop(h)
	}

	return KthLargest{
		k: k,
		h: h,
	}

}

func (this *KthLargest) Add(val int) int {

	heap.Push(this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(this.h)
	}
	// low-rent peek
	return (*this.h)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

//Lifted from golang https://golang.org/pkg/container/heap/

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
