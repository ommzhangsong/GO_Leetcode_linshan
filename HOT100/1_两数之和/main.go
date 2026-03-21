package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for index, value := range nums {
		if index2, ok := m[target-value]; ok {
			return []int{index, index2}
		}
		m[value] = index
	}
	return nil
}
func main() {
	nums := []int{3, 3, 2}
	target := 6
	fmt.Println(twoSum(nums, target))
}
