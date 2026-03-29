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

// 分析：因为题目要的是index，我们可以存map的值为index，我们只需遍历之前的存在性，
/*
不要总是想着map的计数，计数的话，少了索引信息，我们ok语法可以判断存在性，因为如果map中不存在的话，是返回map的v的零值。
那我们index的零值不太好判断，所以可用存在性判断。
ok才是判断存在性的真理！！！！
因为可能value也可能有零值。
*/
