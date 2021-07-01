package code

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	var res []string
	mp := make(map[string]int)
	for i := 0; i < len(cpdomains); i++ {
		//domains := make([]string, 3)
		res := strings.Split(cpdomains[i], " ")
		count, _ := strconv.Atoi(res[0])
		potnum := strings.Count(res[1], ".")
		mp[res[1]] += count
		for j := 0; j < potnum; j++ {
			index := strings.Index(res[1], ".")
			if index != -1 {
				res[1] = res[1][index+1:]
				mp[res[1]] += count
			}
		}
	}
	//fmt.Println(mp)
	for k, v := range mp {
		res = append(res, strconv.Itoa(v)+" "+k)
	}
	//fmt.Println(res)
	return res
}
