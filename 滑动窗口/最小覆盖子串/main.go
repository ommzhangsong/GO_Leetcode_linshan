package main

import "fmt"

/*
 */
func minWindow(s string, t string) string {
	s1, t1 := []byte(s), []byte(t)
	left, need := 0, len(t)
	m := make(map[byte]int)
	resright, resleft := len(s1), 0
	for _, v := range t1 {
		m[v]++
	}
	for right := 0; right < len(s1); right++ {
		if _, ok := m[s1[right]]; ok {
			if m[s1[right]] > 0 {
				need--
			}
			m[s1[right]]--
		}
		for need == 0 {
			if right-left+1 < resright-resleft+1 {
				resright = right
				resleft = left
			}
			if _, ok := m[s1[left]]; ok {
				m[s1[left]]++
				if m[s1[left]] > 0 {
					need++
				}
			}
			left++

		}

	}
	if resright == len(s1) {
		return ""
	}
	return string(s1[resleft : resright+1])
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
