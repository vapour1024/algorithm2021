package week04

func findMin(nums []int) int {
	var l = 0
	var r = len(nums) - 1
	for l < r {
		var mid = (l + r) / 2
		if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r--
		}
	}
	return nums[l]
}
