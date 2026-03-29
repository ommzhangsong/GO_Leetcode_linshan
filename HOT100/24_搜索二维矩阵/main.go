package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	i := 0
	j := n - 1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j-- // 往左
		} else {
			i++ // 往下
		}
	}
	return false
}
func main() {
	matrix := [][]int{{4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 19
	fmt.Println(searchMatrix(matrix, target))
}
