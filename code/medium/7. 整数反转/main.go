package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	// 自己的解法：
	stack := []int{}
	fuhao := 1
	if x < 0 {
		x = -x
		fuhao = -1
	}
	for {
		if x == 0 {
			break
		}
		stack = append(stack, x%10)
		x /= 10
	}
	fmt.Println(stack)
	y := 0
	jinwei := 1
	for i := len(stack) - 1; i >= 0; i-- {
		y += stack[i] * jinwei
		jinwei *= 10
	}
	// 判断x之前是不是负数
	if float64(y) <= math.Pow(2, 31)-1 && float64(y) >= math.Pow(2, -31) {
		if fuhao == -1 {
			y = -y
		}
		return y
	}
	return 0

}

func main() {
	fmt.Println(reverse(321))
}
