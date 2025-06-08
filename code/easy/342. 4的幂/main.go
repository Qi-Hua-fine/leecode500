package main

import "fmt"

func isPowerOfFour(n int) bool {
	// 时间太长
	// x := float64(n)
	// for {
	// 	if x == 1 {
	// 		return true
	// 	}
	// 	if x < 1 {
	// 		break
	// 	}
	// }
	// return false
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}

func main() {
	fmt.Println(isPowerOfFour(1))
}
