package main

type NumArray []int

func Constructor(nums []int) (n NumArray) {
	sum := 0
	n = append(n, sum)
	for _, v := range nums {
		sum += v
		n = append(n, sum)
	}
	return n
}

func (this *NumArray) SumRange(left int, right int) int {
	return (*this)[right+1] - (*this)[left]
}
