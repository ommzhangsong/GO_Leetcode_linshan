package main

func findAnagrams(s string, p string) (res []int) {
	cnt := [26]int{}
	for _, c := range p {
		cnt[c-'a']++
	}
	left := 0
	cnts := [26]int{}
	for right, c := range s {
		cnts[c-'a']++
		if right-left+1 < len(p) {
			left++
		}
		if cnts == cnt {
			res = append(res, left)
		}
		cnt[s[left]-'a']--

	}
	return
}
