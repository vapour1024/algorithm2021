package code

func plusOne(digits []int) []int {
	a := len(digits)
	for i := a - 1; i >= 0; i-- {
		digits[i]++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}
	}
	digits = make([]int, a+1)
	digits[0] = 1
	return digits
}
