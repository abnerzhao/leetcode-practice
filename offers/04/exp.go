// 二维数组快速查找
// https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/

// findNumberIn2DArray 暴力枚举
func findNumberIn2DArray(matrix [][]int, target int) bool {
    for i := 0; i < len(matrix); i++ {
        for j := 0 ; j < len(matrix[i]); j++ {
            if matrix[i][j] == target {
                return true
            }
        }
    }
    return false
}

// findNumberIn2DArrayV2 二维变一维，然后查找
func findNumberIn2DArrayV2(matrix [][]int, target int) bool {
    var arry []int
    for i := range matrix{
        arry = append(arry, matrix[i]...)
    }
    for i := range arry{
        if arry[i] == target {
            return true
        }
    }
    return false
}

// findNumberIn2DArrayV3 使用 SearchInts
func findNumberIn2DArrayV3(matrix [][]int, target int) bool {
	for _, nums := range matrix {
        i := sort.SearchInts(nums, target)
        if i < len(nums) && nums[i] == target {
            return true
        }
    }
    return false
}


