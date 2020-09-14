// leetcode: https://leetcode-cn.com/problems/implement-strstr/

package main

import "fmt"

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

func main() {
	fmt.Println(strStrV1("a", "a"))
}
