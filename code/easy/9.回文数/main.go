package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	// 放入数组中判断
	y := []int{}
	z := 0 // z是余数
	for x > 0 {
		z = x % 10
		x = x / 10
		y = append(y, z)
	}
	for i := 0; i < len(y)/2; i++ {
		if y[i] != y[len(y)-1-i] {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(isPalindrome(123))
}
