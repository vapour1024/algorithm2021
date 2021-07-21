package week05

func minEatingSpeed(piles []int, h int) int {
	maxSpeed := 0
	// 获取piles的最大值，作为右边界
	for _, pile := range piles {
		if pile > maxSpeed {
			maxSpeed = pile
		}
	}
	//设置左右边界
	left := 1
	right := maxSpeed
	answer := right

	for left <= right {
		counter := 0
		mid := (left + right) / 2
		for _, pile := range piles {
			counter += (pile-1)/mid + 1
		}
		if counter > h {
			left = mid + 1
		} else {
			right = mid - 1
			answer = mid
		}
	}
	return answer
}
