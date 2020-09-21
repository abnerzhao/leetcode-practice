package array

/*
https://leetcode-cn.com/problems/implement-strstr/

字符查找，返回字符串出现的位置第一个位置

思路：
1.根据子字符串的长度进行 步长循环查找 数组切片实现 时间复杂度 O(N)
2.双指针 回溯 单个字符比较 第一个字符相同才比较 时间复杂 O(N)
*/

func strStrV1(haystack string, needle string) int {
	needleLenght := len(needle)
	haystackLenght := len(haystack)
	if needleLenght == 0 {
		return 0
	}
	if needleLenght > haystackLenght {
		return -1
	}
	for i := 0; i <= haystackLenght-needleLenght; i++ {
		step := needleLenght + i
		if needle == haystack[i:step] {
			return i
		}
	}
	return -1
}

func strStrV2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}
	return -1
}
