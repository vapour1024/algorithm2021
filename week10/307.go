package week10

var a []Node

type Node struct {
	l, r int
	sum  int
	mark int
}
type NumArray struct {
	n int
	a []Node
}

func build(a []Node, curr int, l int, r int, nums []int) {
	a[curr].l = l
	a[curr].r = r
	a[curr].mark = 0
	if l == r {
		a[curr].sum = nums[l]
		return
	}
	mid := (l + r) / 2
	build(a, curr*2, l, mid, nums)
	build(a, curr*2+1, mid+1, r, nums)
	a[curr].sum = a[curr*2].sum + a[curr*2+1].sum
}
func Constructor(nums []int) NumArray {
	n := len(nums)
	a = make([]Node, 4*n)
	build(a, 1, 0, n-1, nums)
	return NumArray{
		n: n,
		a: a,
	}
}
func (this *NumArray) change(curr int, index int, val int) {
	// 递归边界：叶子[index, index]
	if a[curr].l == a[curr].r {
		a[curr].sum = val
		return
	}
	spread(curr)

	mid := (a[curr].l + a[curr].r) / 2
	if index <= mid {
		this.change(curr*2, index, val)
	} else {
		this.change(curr*2+1, index, val)
	}
	a[curr].sum = a[curr*2].sum + a[curr*2+1].sum
}
func (this *NumArray) query(curr, l, r int) int {
	// 查询的是  [l     ,     r]
	// curr结点是[a[curr].l, a[curr].r]
	// l  a[curr].l  a[curr].r  r
	if l <= a[curr].l && r >= a[curr].r {
		return a[curr].sum
	}
	spread(curr)
	mid := (a[curr].l + a[curr].r) / 2
	ans := 0
	if l <= mid {
		ans += this.query(curr*2, l, r)
	}
	if r > mid {
		ans += this.query(curr*2+1, l, r)
	}
	return ans
}
func spread(curr int) {
	if a[curr].mark != 0 { // 有bug标记
		a[curr*2].sum += a[curr].mark * (a[curr*2].r - a[curr*2].l + 1)
		a[curr*2].mark += a[curr].mark
		a[curr*2+1].sum += a[curr].mark * (a[curr*2+1].r - a[curr*2+1].l + 1)
		a[curr*2+1].mark += a[curr].mark
		a[curr].mark = 0
	}
}
func (this *NumArray) Update(index int, val int) {
	this.change(1, index, val)
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.query(1, left, right)
}
