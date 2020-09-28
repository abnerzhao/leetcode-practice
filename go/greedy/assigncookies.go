package greedy

import (
	"sort"
)

/*
https://leetcode-cn.com/problems/assign-cookies/
分饼干

思路：
尽量满足更多数量的孩子，每次分配尺寸最小的
*/
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}
