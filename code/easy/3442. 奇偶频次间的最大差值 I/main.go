package main

import (
	"fmt"
)

func maxDifference(s string) int {
	byteList := []byte(s)
	byteMap := map[byte]int{}
	// map 中保存了 字符:出现的次数
	for _, by := range byteList {
		if _, ok := byteMap[by]; !ok {
			byteMap[by] = 1
		} else {
			byteMap[by]++
		}
	}
	maxTemp := 0   // 存储最大奇数
	minTemp := 100 // 存储最小偶数
	for i, _ := range byteMap {
		if byteMap[i] > maxTemp && byteMap[i]%2 == 1 {
			maxTemp = byteMap[i]
		}
		if byteMap[i] < minTemp && byteMap[i]%2 == 0 {
			minTemp = byteMap[i]
		}
	}
	return maxTemp - minTemp
}

func main() {
	fmt.Println(maxDifference("abcabcabaaaa"))
}
