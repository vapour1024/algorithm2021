package week06

func climbStairs(n int) int {
	state := make([]int, n+1)
	state[0] = 1
	state[1] = 1
	for i := 2; i <= n; i++ {
		state[i] = state[i-1] + state[i-2]
	}
	return state[n]
}
