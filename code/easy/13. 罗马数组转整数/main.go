package main

import "fmt"

func romanToInt(s string) int {
	hashmap := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	byteArray := []byte(s) // 字符串转为字节数组
	count := 0
	// 循环
	for i := 0; i < len(byteArray); i++ {
		// 总和
		count += hashmap[byteArray[i]]
		// 三个条件 :
		// 1.出现'I'
		// 2.'I'不是最后一个
		// 3. 'I' 右面是'V' 或者 'X'
		if (byteArray[i] == 'I' && i < len(byteArray)-1) && (byteArray[i+1] == 'V' || byteArray[i+1] == 'X') {
			count = count - 2
		}
		if (byteArray[i] == 'X' && i < len(byteArray)-1) && (byteArray[i+1] == 'L' || byteArray[i+1] == 'C') {
			count = count - 20
		}
		if (byteArray[i] == 'C' && i < len(byteArray)-1) && (byteArray[i+1] == 'D' || byteArray[i+1] == 'M') {
			count = count - 200
		}
	}
	return count
}

func main() {
	fmt.Println(romanToInt("CMXCIX"))
}
