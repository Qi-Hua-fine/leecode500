package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	x := ""
	// 排除0和1个的情况
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return string(strs[0])
	}
outloop:
	// j是循环第一个字符串的字符
	for j := 0; j < len(strs[0]); j++ {
		count := 0
		// i是循环字符串数组中的字符都进行比较
		for i := 1; i < len(strs); i++ {
			// 这里 j可能会超出范围 ["ab", "a"]
			if j < len(strs[i]) {
				if strs[0][j] == strs[i][j] {
					count++
				} else {
					break outloop
				}
				if count == len(strs)-1 {
					x += string(strs[0][j])
				}
			}
		}
	}
	return x
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
}
