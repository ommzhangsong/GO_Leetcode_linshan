package main

import "fmt"

func moveZeroes(nums []int) {
	if len(nums) == 1 {
		return
	}
	slow := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[slow] = nums[i]
			slow++
		}
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}
func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
