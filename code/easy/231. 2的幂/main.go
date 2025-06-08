package main

import "fmt"

func isPowerOfTwo(n int) bool {
	return n >= 1 && n&(n-1) == 0
}

func main() {
	fmt.Println(isPowerOfTwo(3))
}
