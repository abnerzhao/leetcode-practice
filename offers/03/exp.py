# -*- coding:utf-8 -*-
"""
找出数组中的重复数字
https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
"""


class Solution(object):
    def findRepeatNumberV1(self, nums):
        """
        借助Map空间换时间，空间复杂度O(n),时间复杂度O(n)
        :type nums: List[int]
        :rtype: int
        """
        repeat_dict = {}
        for i in nums:
            if i not in repeat_dict:
                repeat_dict[i] = 1
            else:
                return i

    def findRepeatNumberV2(self, nums):
        """
        先遍历，然后查询排除该元素的数组中是否还有该值
        """
        for i in xrange(len(nums)):
            n = nums[i]
            if n in (nums[:i] + nums[i + 1:]):
                return n

    def findRepeatNumberV3(self, nums):
        """
        先排序，两两比较
        """
        pass
