package week08

func firstUniqChar(s string) int {

	sl := make([]int, 123)
	for _, i := range s {
		sl[i] += 1
	}
	for k, v := range s {
		if sl[v] == 1 {
			return k
		}
	}
	return -1
}
