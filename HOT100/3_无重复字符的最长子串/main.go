package main

import "fmt"

func lengthOfLongestSubstring(s string) (res int) {
	m := make(map[byte]int)
	nums := []byte(s)
	left := 0
	if s == "" {
		return 0
	}
	res = 1
	for right := 0; right < len(nums); right++ {
		m[nums[right]]++
		for m[nums[right]] > 1 {
			m[nums[left]]--
			left++
		}
		res = max(res, right-left+1)

	}
	return res
}
func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
