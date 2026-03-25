package main

import "sort"

func merge(intervals [][]int) (res [][]int) {
	if len(intervals) == 0 {
		return res
	}
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		if last[1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else {
			last[1] = max(last[1], intervals[i][1])
		}

	}
	return res
}
