// https://leetcode-cn.com/problems/two-sum/ 两数之和

package main

import "fmt"

func twoNum(nums []int, target int) []int {
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

func main() {
	nums := []int{2, 7, 11, 15}
	target := 22
	fmt.Println(twoNum(nums, target))
}
