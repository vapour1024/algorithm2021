package week09

import "container/heap"

var a []int

//大根堆
type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return a[(*h)[i]] > a[(*h)[j]]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func maxSlidingWindow(nums []int, k int) []int {
	//懒惰删除
	a = nums
	ans := []int{}
	q := new(myHeap)
	for i := 0; i < k-1; i++ {
		heap.Push(q, i)
	}
	for i := k - 1; i < len(nums); i++ {
		heap.Push(q, i)
		for (*q)[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[(*q)[0]])
	}
	return ans
}
