package main

import "fmt"

/*
给你一个二进制数组 nums ，你需要从中删掉一个元素。
请你在删掉元素的结果数组中，返回最长的且只包含 1 的非空子数组的长度。
如果不存在这样的子数组，请返回 0 。
提示 1：
输入：nums = [1,1,0,1]
输出：3
解释：删掉位置 2 的数后，[1,1,1] 包含 3 个 1 。
*/
/*
分析：其实就和翻转一样，我先k=1,然后求出最后的子数组长度-1就行了
*/
func longestSubarray(nums []int) (res int) {
	left, n := 0, len(nums)
	zero := 0
	for right := 0; right < n; right++ {
		if nums[right] == 0 {
			zero++
		}
		for zero > 1 {
			if nums[left] == 0 {
				zero--
			}
			left++
		}
		if res < right-left {
			res = right - left
		}
	}
	return res
}
func main() {
	var nums = []int{1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))
}
