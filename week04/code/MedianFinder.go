package week04
import "container/heap"

type MinHeap struct {
	nums []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (minH *MinHeap) Len() int {
	return len(minH.nums)
}

func (minH *MinHeap) Swap(i, j int) {
	minH.nums[i], minH.nums[j] = minH.nums[j], minH.nums[i]
}

func (minH *MinHeap) Less(i, j int) bool {
	return minH.nums[i] < minH.nums[j]
}

func (minH *MinHeap) Push(x interface{}) {
	minH.nums = append(minH.nums, x.(int))
}

func (minH *MinHeap) Pop() interface{} {
	x := minH.nums[minH.Len()-1]
	minH.nums = minH.nums[:minH.Len()-1]
	return x
}

func (minH *MinHeap) Top() int {
	x := minH.nums[0] // 注意：在这里(golang里)堆顶（无论是大顶堆，还是小顶堆）是第1个元素
	return x
}

// MaxHeap
type MaxHeap struct {
	nums []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (maxH *MaxHeap) Len() int {
	return len(maxH.nums)
}

func (maxH *MaxHeap) Swap(i, j int) {
	maxH.nums[i], maxH.nums[j] = maxH.nums[j], maxH.nums[i]
}

func (maxH *MaxHeap) Less(i, j int) bool {
	return maxH.nums[i] > maxH.nums[j]
}

func (maxH *MaxHeap) Push(x interface{}) {
	maxH.nums = append(maxH.nums, x.(int))
}

func (maxH *MaxHeap) Pop() interface{} {
	x := maxH.nums[maxH.Len()-1]
	maxH.nums = maxH.nums[:maxH.Len()-1]
	return x
}

func (maxH *MaxHeap) Top() int {
	return maxH.nums[0] // 注意：在这里(golang里)堆顶（无论是大顶堆，还是小顶堆）是第1个元素
}

// MedianFinder
// 用大、小顶堆存储数据。
// 其中前半部分是大顶堆
// 后半部分是小顶堆
// 而且前半部分（大顶堆）最多比后半部分（小顶堆）多存一个元素
// 因此：
//    当元素总共有奇数个时，则返回大顶堆的堆顶元素。
//    当元素纵隔有偶数个时，则返回两个堆顶元素之和再除以2.0
type MedianFinder struct {
	maxH *MaxHeap // 前半部分，用大顶堆存储，都是较小的数（都小于后半部分）
	minH *MinHeap // 后半部分，用小顶堆存储，都是较大的数（都大于前半部分）
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	res := MedianFinder{
		maxH: NewMaxHeap(),
		minH: NewMinHeap(),
	}
	heap.Init(res.maxH)
	heap.Init(res.minH)
	return res
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxH, num) // 先加入到大顶堆

	heap.Push(this.minH, heap.Pop(this.maxH)) // 调整两个堆平衡，此时从大顶堆Pop出最大元素，加入到小顶堆。

	for this.maxH.Len() < this.minH.Len() {
		heap.Push(this.maxH, heap.Pop(this.minH))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxH.Len() > this.minH.Len() {
		return float64(this.maxH.Top())
	}

	return float64(this.maxH.Top()+this.minH.Top()) * 0.5
}