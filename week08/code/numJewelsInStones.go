package week08

func numJewelsInStones(jewels string, stones string) int {
	jmap := make(map[int32]bool, len(jewels))
	res := 0
	for _, i := range jewels {
		jmap[i] = true
	}
	for _, s := range stones {
		if jmap[s] == true {
			res++
		}
	}
	return res
}
