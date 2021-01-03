// 求「众数」，数组中出现次数超过一半的数字
// https://leetcode-cn.com/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/

// majorityElement 借助字典
func majorityElement(nums []int) int {
    if len(nums) <= 1 {
        return nums[0]
    }
    lenght := len(nums)/2
    tmp := make(map[int]int)
    for i := range nums {
        val := nums[i]
        if _, ok:= tmp[val]; ok {
            if tmp[val] >= lenght {
                return val
            }
            tmp[val]++
        } else {
            tmp[val] = 1
        }
    }
    return 0
}

// majorityElementV2 排序取中间值
func majorityElementV2(nums []int) int {
    sort.Ints(nums)
    return nums[len(nums)/2]
}

// majorityElementV3 摩尔投票法
// TODO
func majorityElementV3(nums []int)int {

}
