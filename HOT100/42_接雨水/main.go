package main

import "fmt"

func trap(height []int) (ans int) {
	left := 0
	right := len(height) - 1
	premax, sufmax := 0, 0
	for left <= right {
		premax = max(premax, height[left])
		sufmax = max(sufmax, height[right])
		if premax < sufmax {
			ans += premax - height[left]
			left++
		} else {
			ans += sufmax - height[right]
			right--
		}
	}
	return ans

}
func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
