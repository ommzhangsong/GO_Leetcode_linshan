package main

import "fmt"

func rotate(nums []int, k int) {

	k = k % len(nums)
	if k == 0 {
		return
	}
	var res []int
	res = append(res, nums[len(nums)-k:]...)
	res = append(res, nums[:len(nums)-k]...)
	copy(nums, res)
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
