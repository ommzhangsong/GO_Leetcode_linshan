package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) (res [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1

		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum < 0 {
				left++
			}
			if sum > 0 {
				right--
			}
			if sum == 0 {
				res = append(res, []int{nums[left], nums[right], nums[i]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for nums[right] == nums[right-1] && left < right {
					right--
				}
				left++
				right--
			}
		}

	}
	return res
}
func main() {
	var a = []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(a))

}
