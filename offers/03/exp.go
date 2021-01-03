package main

// 找出数组中的重复数字
// https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/


import (
	"fmt"
	"sort"
)

// findRepeatNumber 借助Map
func findRepeatNumber(nums []int) int {
    repeatDict := make(map[int]bool)
    for i := range nums{
        if repeatDict[i] {
            return nums[i]
        } else {
            repeatDict[i] = true
        }
    }
    return 0
}

// findReatNumberV2 先排序后两两比较
func findReatNumberV2(nums []int)int{
	sort.Ints(nums)
	for i := range nums{
        if nums[i] == nums[i+1]  {
            return nums[i]
        }
    }
    return 0
}

// func findReatNumberV3(nums []int) int{
// 	for i := range nums{
// 		newSlice := append(nums[:i], nums[i+1:]...)
// 		fmt.Println(newSlice)
// 		fmt.Println(nums)
// 		idx := sort.SearchInts(newSlice, nums[i])
// 		fmt.Println(idx)
// 		if idx > len(newSlice) || newSlice[idx] != nums[i]{
// 			return nums[i]
// 		}    
// 	}
// 	return 0
// }
