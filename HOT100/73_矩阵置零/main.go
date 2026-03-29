package main

import "fmt"

func setZeroes(matrix [][]int) {
	firstrow := false
	firstcol := false
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			firstrow = true
			break
		}
	}

	// 判断首列
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstcol = true
			break
		}
	}
	if len(matrix) == 1 {
		return
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if firstrow == true {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if firstcol == true {
		for j := 0; j < len(matrix); j++ {
			matrix[j][0] = 0
		}
	}
}
func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
