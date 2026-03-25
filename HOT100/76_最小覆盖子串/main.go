package main

import "fmt"

func minWindow(s string, t string) string {
	m := make(map[byte]int)

	s1 := []byte(s)
	t2 := []byte(t)
	for _, v := range t2 {
		m[v]++
	}
	left, need := 0, len(t)
	resright := len(s1)
	resleft := 0
	for right := 0; right < len(s1); right++ {

		m[s1[right]]--
		if m[s1[right]] >= 0 {
			need--
		}
		for need == 0 {
			if right-left+1 < resright-resleft+1 {
				resright = right
				resleft = left
			}
			m[s1[left]]++
			if m[s1[left]] > 0 {
				need++
			}
			left++
		}
	}
	if resright == len(s1) {
		return ""
	}
	return s[resleft : resright+1]
}
func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
