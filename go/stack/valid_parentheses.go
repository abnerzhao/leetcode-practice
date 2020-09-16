package stack

// https://leetcode-cn.com/problems/valid-parentheses/
// 匹配到括号对则出栈

func isValid(s string) bool {
	s := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]"
	}
}
