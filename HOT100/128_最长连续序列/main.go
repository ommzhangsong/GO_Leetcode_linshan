package main

import (
	"fmt"
	"maps"
	"slices"
)

func longestConsecutive(nums []int) (res int) {
	if len(nums) == 0 {
		return 0
	}
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	tmp := 1
	a := slices.Collect(maps.Keys(m))
	slices.Sort(a)
	fmt.Println(a)
	for i := 0; i < len(a)-1; i++ {
		if a[i]+1 == a[i+1] {
			tmp++
			if res < tmp {
				res = tmp
			}
		} else {

			tmp = 1
		}
	}
	return res
}
func longestConsecutive2(nums []int) (res int) {
	if len(nums) == 0 {
		return 0
	}
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	res = 1
	longest := 1
	for n := range m {
		if !m[n-1] {
			curnum := n
			for m[curnum+1] {
				longest++
				curnum++
				if res < longest {
					res = longest
				}
			}
		}
		longest = 1
	}
	return res
}

func main() {
	nums := []int{1, 0, 1, 2}
	fmt.Println(longestConsecutive2(nums))
}
