package main

import "fmt"

func isValid(s string) bool {
	stack := []byte{}
	byteList := []byte(s) // 长度必须是偶数
	if len(byteList)%2 != 0 {
		return false
	}
	all := map[byte]byte{'{': '}', '[': ']', '(': ')'}
	for _, c := range byteList {
		// 当你尝试访问一个 map 中不存在的 key 时，返回的值是该 value 类型的零值
		if _, ok := all[c]; ok {
			stack = append(stack, c)
		} else {
			// 比较最后一个左 对应的右 是不是key
			if len(stack) == 0 || all[stack[len(stack)-1]] != c {
				return false
			}
			stack = stack[:len(stack)-1]
			fmt.Println(stack)
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
}
