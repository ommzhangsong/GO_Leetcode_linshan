package main

import "fmt"

func maxArea(height []int) (res int) {
	left, n := 0, len(height)
	right := n - 1
	res = (right - left) * min(height[left], height[right])
	for left < right {
		if height[left] < height[right] {
			left++
			res = max(res, (right-left)*min(height[left], height[right]))
		} else {
			right--
			res = max(res, (right-left)*min(height[left], height[right]))
		}
	}
	return res

}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}
