package main

import "fmt"

// 找到任意的峰值
func findPeakElement(nums []int) int {
	left := 0
	right := len(nums) - 1
	var mid int

	for left < right {
		mid = left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(nums))
}
