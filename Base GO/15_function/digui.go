package main

import "fmt"

// 递归函数(自调用)
func sumss(num int) int {
	if num == 1 {
		return num
	}
	return sumss(num-1) + num
}
func main() {
	// 求和1-10
	fmt.Println(sumss(10))
}
