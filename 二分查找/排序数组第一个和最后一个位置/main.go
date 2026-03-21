package main

import "fmt"

func lowerbound(nums []int, t int) (res int) {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if mid < t {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
func searchRange(nums []int, target int) []int {
	start := lowerbound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := lowerbound(nums, target+1) - 1
	return []int{start, end}
}
func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	t := 8
	fmt.Println(lowerbound(nums, t))
}
