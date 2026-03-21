package main

import (
	"maps"
	"slices"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	//存的字符串切片，key是排序好的字符串，v是排序前
	for _, v := range strs {
		tmp := []byte(v)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		sorteds := string(tmp)
		m[sorteds] = append(m[sorteds], v)
	}
	return slices.Collect(maps.Values(m))
}
