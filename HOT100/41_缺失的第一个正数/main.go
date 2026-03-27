package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[i]-1 != i {
			swap(nums, i, nums[i]-1)
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	// 全部都对，返回 n+1
	return n + 1
}
func swap(nums []int, i int, j int) {
	tmp := nums[j]
	nums[j] = nums[i]
	nums[i] = tmp
}
func main() {
	nums := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(nums))
}
