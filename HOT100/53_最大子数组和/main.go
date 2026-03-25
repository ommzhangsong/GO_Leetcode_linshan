package main

import (
	"fmt"
)

func maxSubArray(nums []int) (res int) {
	mnums := make([]int, len(nums)+1)
	for i, v := range nums {
		mnums[i+1] = mnums[i] + v
	}
	minpre := 0
	res = mnums[1]
	for i := 1; i < len(mnums); i++ {
		res = max(res, mnums[i]-minpre)
		minpre = min(minpre, mnums[i])
	}
	return res

}
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
