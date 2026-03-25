package main

import "fmt"

// 2026.3.24 ,暴力破解, 失败。
//2. 看答案,完成，待复习

func maxSlidingWindow(nums []int, k int) (ans []int) {
	q := []int{}
	for i, v := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= v {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if i-k+1 > q[0] {
			q = q[1:]
		}
		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}
	return
}
func main() {
	nums := []int{1, 3, 1, 2, 0, 5}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}

/*
3.25 复习一次，这里是用到了单调队列，每次加入一个索引前，如果这个索引的值比队列中的部分值都大，我们就可以把队列中的淘汰，判断头结点是否过期，就可以了
*/
