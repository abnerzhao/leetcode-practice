// 括号匹配 https://leetcode-cn.com/problems/valid-parentheses/
// 栈实现

func isValid(s string) bool {
    parentheses := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	if s == "" {
		return true
	}
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			// 入栈
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack) - 1] == parentheses[s[i]] {
			// 出栈
			stack = stack[:len(stack) -1]
		} else {
			return false
		}
		
	}
	return len(stack) == 0
}
