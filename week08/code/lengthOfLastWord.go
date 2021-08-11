package week08

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	return len(s) - strings.LastIndex(s, " ") - 1
}
