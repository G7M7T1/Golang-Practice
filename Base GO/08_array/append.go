package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	x = append(x, 6, 7, 8)
	fmt.Println(x)

	y := []int{9, 10, 11}
	x = append(x, y...)
	fmt.Println(x)

	z := make([]int, 5, 10)
	fmt.Println(z)

}
