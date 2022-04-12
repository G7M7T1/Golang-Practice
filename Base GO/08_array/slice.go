package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(x[:])
	fmt.Println(x[2:])
	fmt.Println(x[2:3])
}
