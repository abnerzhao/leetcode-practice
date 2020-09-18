package stack

// https://leetcode-cn.com/problems/valid-parentheses/
// 匹配到括号对则出栈

func isValid(s string) bool {
	pattern := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	var stackList []string
	for _, v := range s {
		i := string(v)
		if len(stackList) > 0 && stackList[len(stackList)-1] == pattern[i] {
			stackList = stackList[:len(stackList)-1]
		} else {
			stackList = append(stackList, i)
		}
	}
	return len(stackList) == 0
}
