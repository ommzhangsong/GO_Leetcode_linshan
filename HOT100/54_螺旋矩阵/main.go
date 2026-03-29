package main

import "fmt"

func spiralOrder(matrix [][]int) (res []int) {
	m := len(matrix)
	n := len(matrix[0])
	top := 0
	bottom := m - 1
	left := 0
	right := n - 1
	for top <= bottom && left <= right {
		//从左到右
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		// 上到下
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// 右到左
		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}
	return res
}
func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
}
