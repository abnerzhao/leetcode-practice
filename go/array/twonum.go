package array

/*
https://leetcode-cn.com/problems/two-sum

思路：
1.两层for循环 时间复杂度O(N^2)
2.求和问题转换为求差 空间换时间 借用 map 存储元素和索引关系 时间复杂度为 O(1)
*/

func towNumV1(nums []int, target int) []int {
	var data []int
	for k, v := range nums {
		for i, j := range nums[k:] {
			if v+j == target {
				data = []int{k, i}
			}
		}
	}
	return data
}

func twoNumV2(nums []int, target int) []int {
	mapDict := make(map[int]int)
	var data []int
	for i, k := range nums {
		_, status := mapDict[target-k]
		if !status {
			mapDict[k] = i
		} else {
			data = []int{i, mapDict[target-k]}
			break
		}
	}
	return data
}
