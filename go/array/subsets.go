package array

/*
https://leetcode-cn.com/problems/subsets/
子集

思路：
TODO
*/

func getLenOfSet(n int, nums []int) [][]int {
	var sets [][]int
	switch n {
	case 0:
		return sets
	case 1:
		for v := range nums {
			sets = append(sets, []int{v})
		}
	default:
		for i := 0; i < len(nums)-n+1; i++ {
			for j := 0; j < n; j++ {
				ss := []int{nums[i], nums[i+j]}
				sets = append(sets, ss)
			}
		}
	}
	return sets
}
